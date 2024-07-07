package scheduler

import (
	"os/exec"
	"testing"

	"icode.baidu.com/liyinjie/minispider/loader"
)

func TestGetScheduler(t *testing.T) {
	scheduler := GetScheduler()
	if scheduler == nil {
		t.Errorf("GetScheduler() should not return nil")
	}
}

func TestInit(t *testing.T) {
	scheduler := GetScheduler()
	if scheduler == nil {
		t.Errorf("GetScheduler() should not return nil")
	}
	config := loader.Config{loader.Spider{
		UrlListFile:     "../data/url.data",
		OutputDirectory: "../test_data",
		MaxDepth:        2,
		CrawlInterval:   1,
		CrawlTimeout:    1,
		TargetUrl:       ".*.(htm|html)$",
		ThreadCount:     1,
	}}
	mkdir := exec.Command("mkdir", "-p", "../test_data")
	error := mkdir.Start()
	if error != nil {
		t.Errorf("mkdir -p ../test_data should be success, err is %s\n", error.Error())
	}
	mkdir.Wait()
	chmod := exec.Command("chmod", "777", "../test_data")
	error = chmod.Start()
	if error != nil {
		t.Errorf("chmod 777 ../test_data should be success, err is %s\n", error.Error())
	}
	chmod.Wait()
	seeds := []string{}
	scheduler.Init(config, seeds)
	scheduler.Start()

	deldir := exec.Command("rm", "-rf", "../test_data")
	error = deldir.Start()
	if error != nil {
		t.Errorf("rm -rf ../test_data should be success, err is %s\n", error.Error())
	}
	deldir.Wait()
}
