package scheduler

import (
	"regexp"
	"sync"
	"time"

	"github.com/baidu/go-lib/log"
	"github.com/baidu/go-lib/queue"
	"icode.baidu.com/liyinjie/minispider/loader"
	"icode.baidu.com/liyinjie/minispider/parse"
)

type Scheduler struct {
	//任务队列
	TaskQue queue.Queue
	//url去重<url, bool>
	UrlTab sync.Map
	// 任务channel
	TaskChan chan struct{}
	//最大爬取深度
	MaxDepth int
	//爬取的间隔
	CrawlInterval int
	//go routine数量
	ThreadCount int
	//通用配置器
	TaskConfig *TaskComConfig
	// 站点爬取间隔timer表
	TimerTab sync.Map
}

func GetScheduler() *Scheduler {
	return new(Scheduler)
}

func (s *Scheduler) Init(config loader.Config, seeds []string) {
	log.Logger.Info("Scheduler init")
	s.TaskQue.Init()
	s.MaxDepth = config.MaxDepth
	s.CrawlInterval = config.CrawlInterval
	s.ThreadCount = config.ThreadCount
	s.TaskChan = make(chan struct{}, config.ThreadCount)
	targetUrlPattern, _ := regexp.Compile(config.TargetUrl)
	//生成一个通用任务配置器
	taskCommonCfg := &TaskComConfig{
		CrawlTimeout:  config.CrawlTimeout,
		OutPutDir:     config.OutputDirectory,
		TarUrlPattern: targetUrlPattern,
	}
	s.TaskConfig = taskCommonCfg
	//对于每一个种子文件，生成一个Task放入队列
	for _, seed := range seeds {
		task := &Task{
			Url:    seed,
			Depth:  0,
			ComFig: taskCommonCfg,
		}
		s.TaskQue.Append(task)
	}
}

func (s *Scheduler) Start() {
	log.Logger.Info("Scheduler Start")
	defer close(s.TaskChan)
	for {
		if s.TaskQue.Len() == 0 && len(s.TaskChan) == 0 {
			break
		}
		task := s.TaskQue.Remove()
		s.Processtask(task.(*Task))
	}

	log.Logger.Info("Scheduler stop")
}

func (s *Scheduler) Processtask(task *Task) {
	//抓取深度不满足调度器最大深度
	if task.Depth >= s.MaxDepth {
		return
	}
	_, ok := s.UrlTab.LoadOrStore(task.Url, true)

	if ok {
		// 该url的内容正在抓取或者已经抓取过了 直接返回
		return
	}
	//向管道中写入一个数据，证明有任务执行，防止调度器退出
	s.TaskChan <- struct{}{}
	//开一个线程处理，主线程回到循环读取Task
	go func() {
		defer func() {
			log.Logger.Info("task %s done", task.Url)
			<-s.TaskChan
		}()
		hostName, err := parse.ParseHostName(task.Url)
		if err != nil {
			log.Logger.Error("%s: parser.ParseHostName(): %s", task.Url, err.Error())
			return
		}
		timer, ok := s.TimerTab.LoadOrStore(hostName, time.NewTimer(time.Duration(s.CrawlInterval)*time.Second))
		//如果存在url的定时器，等待定时器结束后再进行抓取，并重置一个定时器加入
		if ok {
			select {
			case <-timer.(*time.Timer).C:
			}
			timer.(*time.Timer).Reset(time.Duration(s.CrawlInterval) * time.Second)
		}

		newUrlList, err := task.Run()
		if err != nil {
			log.Logger.Error("%s", err.Error())
			return
		}
		//
		for _, url := range newUrlList {
			newTask := &Task{
				Url:    url,
				Depth:  task.Depth + 1,
				ComFig: s.TaskConfig,
			}
			s.TaskQue.Append(newTask)
		}
	}()

}
