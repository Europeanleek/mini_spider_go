package crawler

import (
	"fmt"
	"testing"
)

func TestCraw(t *testing.T) {
	TestCase := []struct {
		url     string
		timeout int
		err     error
	}{
		{
			url:     "http://www.baidu.com",
			timeout: 3,
			err:     nil,
		},
		{
			url:     "http://www.ahsfoidhasfdoiuhsafdoi.com",
			timeout: 3,
			err:     fmt.Errorf("error"),
		},
		{
			url:     "asfasfd",
			timeout: 3,
			err:     fmt.Errorf("error"),
		},
	}
	for _, value := range TestCase {
		_, _, err := Crawl(value.url, value.timeout)
		if value.err == nil {
			if err != nil {
				t.Errorf("error, %s\n", err.Error())
			}
		} else {
			if err == nil {
				t.Errorf("error,%s\n", err.Error())
			}
		}

	}
}
