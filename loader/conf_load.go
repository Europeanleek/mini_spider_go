package loader

import (
	"fmt"
	"regexp"

	"gopkg.in/gcfg.v1"
)

type Spider struct {
	// 种子文件路径
	UrlListFile string
	// 抓取结果存储目录
	OutputDirectory string
	// 最大抓取深度
	MaxDepth int
	// 抓取间隔. 单位: 秒
	CrawlInterval int
	// 抓取超时. 单位: 秒
	CrawlTimeout int
	// 需要存储的目标网页URL Pattern
	TargetUrl string
	// 抓取routine数
	ThreadCount int
}

type Config struct {
	Spider
}

// 加载配置文件
func ConfigLoad(confPath string) (Config, error) {
	var cfg Config
	var err error
	err = gcfg.ReadFileInto(&cfg, confPath)
	if err != nil {
		return cfg, err
	}

	if cfg.UrlListFile == "" {
		return cfg, fmt.Errorf("UrlListFile is nil")
	}
	if cfg.OutputDirectory == "" {
		return cfg, fmt.Errorf("OutputDirectory is nil")
	}
	if cfg.MaxDepth <= 0 {
		return cfg, fmt.Errorf("MaxDepth error")
	}
	if cfg.CrawlInterval < 0 {
		return cfg, fmt.Errorf("CrawInterval error")
	}
	if cfg.CrawlTimeout < 1 {
		return cfg, fmt.Errorf("CrawTimeout error")
	}
	_, err = regexp.Compile(cfg.TargetUrl)

	if err != nil {
		return cfg, fmt.Errorf("regexp error")
	}
	if cfg.ThreadCount < 1 {
		return cfg, fmt.Errorf("ThreadCount<1 error")
	}
	return cfg, nil

}
