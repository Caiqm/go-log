# go-log

简易日志保存

### 使用

1. 拉取代码

> go get github.com/Caiqm/go-log

2. 方法调用

``` golang
package main

import "github.com/Caiqm/go-log/log"

func main() {
    // 初始化客户端，第一个参数是保存目录，第二个参数是日志文件后缀
	log := log.NewLogClient("runtime/log", "log")
	log.Debug(`{"a":123}`)
	log.Info(`{"a":123}`)
	log.Warn(`{"a":123}`)
	log.Error(`{"a":123}`)
	log.Fatal(`{"a":123}`)
}
```
