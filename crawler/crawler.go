package crawler

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	HeaderKeyUserAgent = "User-Agent"
)

const (
	Mozilla = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36"
)

func Crawl(url string, timeout int) ([]byte, string, error) {
	var body []byte
	var contentType string
	var err error
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", fmt.Errorf("%s: http.NewRequest(): %s", url, err.Error())
	}
	//添加请求头
	req.Header.Add(HeaderKeyUserAgent, Mozilla)
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	errChan := make(chan error, 1)
	go func() {
		var response *http.Response
		response, err = client.Do(req)
		if err != nil {
			err = fmt.Errorf("%s: client.Do(): %s", url, err.Error())
			errChan <- err
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			err = fmt.Errorf("%s: status code[%d] not 200", url, response.StatusCode)
			errChan <- err
			return
		}

		contentType = response.Header.Get("Content-Type")

		body, err = io.ReadAll(response.Body)
		if err != nil {
			err = fmt.Errorf("ioutil.ReadAll(): %s", err.Error())
			errChan <- err
			return
		}
		errChan <- nil
	}()

	select {
	case err = <-errChan:
		if err != nil {
			return nil, "", err
		}
	case <-timer.C:
		return nil, "", fmt.Errorf("crawl timeout")
	}

	return body, contentType, err

}
