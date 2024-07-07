package loader

import (
	"fmt"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	TestCases := []struct {
		confPath string
		err      error
	}{
		{
			confPath: "../conf/spider.conf",
			err:      nil,
		},
		{
			confPath: "./spider_1_test.conf",
			err:      fmt.Errorf("Error"),
		},
		{
			confPath: "./spider_2_test.conf",
			err:      fmt.Errorf("Error"),
		},
		{
			confPath: "./spider_3_test.conf",
			err:      fmt.Errorf("Error"),
		},
		{
			confPath: "./spider_4_test.conf",
			err:      fmt.Errorf("Error"),
		}, {
			confPath: "./spider_5_test.conf",
			err:      fmt.Errorf("Error"),
		}, {
			confPath: "./spider_6_test.conf",
			err:      fmt.Errorf("Error"),
		},
	}

	for index, value := range TestCases {
		_, err := ConfigLoad(value.confPath)
		if value.err == nil {
			if err != nil {
				t.Errorf("TestConfigLoad failed, index:%d, %s\n", index, err.Error())
			}
		} else {
			if err == nil {
				t.Errorf("TestConfigLoad failed, index:%d, %s\n", index, err.Error())
			}
		}

	}

}
