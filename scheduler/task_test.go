package scheduler

import (
	"os/exec"
	"testing"
)

func TestSaveData(t *testing.T) {
	data := []byte("Test Save Data")
	url := "www.baidu.com"
	mkdir := exec.Command("mkdir", "-p", "../output_test")
	error := mkdir.Start()
	if error != nil {
		t.Errorf("mkdir error + %s\n", error.Error())
	}
	mkdir.Wait()
	error = SaveData(data, url, "../output_test")
	if error != nil {
		t.Errorf("save data error + %s\n", error.Error())
	}
	rmdir := exec.Command("rm", "-rf", "../output_test")
	error = rmdir.Start()
	if error != nil {
		t.Errorf("rmdir error + %s\n", error.Error())
	}
	rmdir.Wait()

}
