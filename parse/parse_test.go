package parse

import (
	"fmt"
	"testing"

	"icode.baidu.com/liyinjie/minispider/crawler"
)

func TestConver2Utf8(t *testing.T) {
	body, contentType, err := crawler.Crawl("http://www.zhihu.com", 3)
	if err != nil {
		t.Errorf("Crawl error: %s\n", err.Error())
	}
	_, err = Convert2Utf8(body, contentType)
	if err != nil {
		t.Errorf("Convert2Utf8 error: %s\n", err.Error())
	}
}

func TestParseHostName(t *testing.T) {
	TestCases := []struct {
		url string
		err error
	}{
		{
			url: "https://www.baidu.com",
			err: nil,
		},
		{
			url: "http://www.baidu.com:8080",
			err: nil,
		},
		{
			url: "sadfaf",
			err: fmt.Errorf("error"),
		},
	}

	for _, testCase := range TestCases {
		_, err := ParseHostName(testCase.url)
		if testCase.err == nil {
			if err != nil {
				t.Errorf("ParseHostName error: %s\n", err.Error())
			}
		} else {
			if err == nil {
				t.Errorf("ParseHostName error: %s\n", testCase.err)
			}
		}
	}
}

func TestGetUrlList(t *testing.T) {
	testUrl := "https://www.baidu.com"

	data, _, err := crawler.Crawl(testUrl, 3)
	if err != nil {
		t.Errorf("%s: crawler.Crawl(): %s", testUrl, err.Error())
		return
	}

	urlList, err := GetUrlList(data, testUrl)
	if err != nil {
		t.Errorf("GetUrlList(): %s", err.Error())
		return
	}

	if len(urlList) == 0 {
		t.Errorf("no sublink in %s", testUrl)
		return
	}
}
