# mini-spider
## 目录结构
- conf：爬虫配置文件
- crwaler：
- data：url种子数据
- loader：配置文件加载器以及种子加载器
- parse：协议解析
- scheduler：任务调度器以及Task任务

## 编译
make .

## 运行
cd bin
./server -c [配置文件路径] -l [日志文件路径]