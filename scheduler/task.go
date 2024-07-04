package scheduler

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"icode.baidu.com/liyinjie/minispider/crawler"
	"icode.baidu.com/liyinjie/minispider/parse"
)

// 每个Task的通用任务配置
type TaskComConfig struct {
	//超时时间
	CrawlTimeout int
	//保存目录
	OutPutDir string
	//正则表达式
	TarUrlPattern *regexp.Regexp
}

// 任务类
type Task struct {
	//抓取的Url
	Url string
	//抓取的深度
	Depth int
	//通用配置
	ComFig *TaskComConfig
}

func (task *Task) Run() ([]string, error) {
	data, contentType, err := crawler.Crawl(task.Url, task.ComFig.CrawlTimeout)
	//抓取失败
	if err != nil {
		return nil, fmt.Errorf("%s: crawler.Crawl(): %s", task.Url, err.Error())
	}
	if !strings.Contains(contentType, "text") {
		return nil, fmt.Errorf("%s: Content-Type: %s", task.Url, contentType)
	}
	data, err = parse.Convert2Utf8(data, contentType)
	if err != nil {
		return nil, fmt.Errorf("%s: parser.Convert2Utf8(): %s", task.Url, err.Error())
	}

	if task.ComFig.TarUrlPattern.MatchString(task.Url) {
		err = task.SaveData(data)
		if err != nil {
			return nil, fmt.Errorf("%s: task.SaveData(): %s", task.Url, err.Error())
		}
	}

	urlList, err := parse.GetUrlList(data, task.Url)
	if err != nil {
		return nil, fmt.Errorf("%s: parser.GetUrlList(): %s", task.Url, err.Error())
	}

	return urlList, nil
}

func (task *Task) SaveData(data []byte) error {
	err := SaveData(data, task.Url, task.ComFig.OutPutDir)
	if err != nil {
		return fmt.Errorf("saver.SaveData(): %s", err.Error())
	}

	return nil
}

func SaveData(data []byte, urlStr string, outputDirectory string) error {
	fileName := filepath.Join(outputDirectory, url.QueryEscape(urlStr))
	fmt.Println(fileName)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("%s: os.OpenFile(): %s", fileName, err.Error())
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("f.Write(): %s", err.Error())
	}

	return nil
}
