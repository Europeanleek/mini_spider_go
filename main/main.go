package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/baidu/go-lib/log"
	"github.com/baidu/go-lib/log/log4go"
	"icode.baidu.com/liyinjie/minispider/loader"
	"icode.baidu.com/liyinjie/minispider/scheduler"
)

var (
	confPath = flag.String("c", "../conf", "root path of configuration")
	logPath  = flag.String("l", "../log", "dir path of log")
	help     = flag.Bool("h", false, "to show help")
)

const ConfFileName = "spider.conf"

func initLog(logSwitch string, logPath *string, stdOut bool) error {
	log4go.SetLogBufferLength(10000)
	log4go.SetLogWithBlocking(false)

	err := log.Init("mini_spider", logSwitch, *logPath, stdOut, "midnight", 5)
	if err != nil {
		return fmt.Errorf("err in log.Init(): %s", err.Error())
	}

	return nil
}
func main() {
	//解析命令行参数
	flag.Parse()
	//显示帮助
	if *help {
		flag.PrintDefaults()
		return
	}
	var err error
	err = initLog("INFO", logPath, true)
	if err != nil {
		fmt.Printf("initLog(): %s\n", err.Error())
		log.Logger.Close()
		time.Sleep(100 * time.Millisecond)
		os.Exit(-1)
	}
	//创建配置文件
	var config loader.Config
	config, err = loader.ConfigLoad(filepath.Join(*confPath, ConfFileName))

	if err != nil {
		log.Logger.Error("loader.ConfigLoad(): %s", err.Error())
		log.Logger.Close()
		time.Sleep(100 * time.Millisecond)
		os.Exit(-1)
	}
	//加载种子文件
	seeds, err := loader.SeedLoad(config.UrlListFile)

	if err != nil {
		log.Logger.Error("loader.SeedLoad(): %s", err.Error())
		log.Logger.Close()
		time.Sleep(100 * time.Millisecond)
		os.Exit(-1)
	}
	mini_spider_sch := scheduler.GetScheduler()
	//初始化调度器，加载配置文件
	mini_spider_sch.Init(config, seeds)
	mini_spider_sch.Start()
	log.Logger.Close()
	time.Sleep(100 * time.Millisecond)
	os.Exit(-1)

}
