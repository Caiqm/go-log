package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Caiqm/go-log/util"
)

type LogClient struct {
	file *os.File
}

func NewLogClient(filePath, saveName string) *LogClient {
	F, err := openLogFile(getLogFileName(saveName, "log"), filePath)
	if err != nil {
		log.Fatalln(err)
	}
	return &LogClient{
		file: F,
	}
}

// 日志文件名称
func getLogFileName(saveName, logFileExt string) string {
	return fmt.Sprintf("%s-%s.%s", saveName, time.Now().Format("20060102"), logFileExt)
}

// 打开日志文件
func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	src := filepath.Join(dir, filePath)
	perm := util.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	err = util.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	f, err := util.Open(filepath.Join(src, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile: %v", err)
	}
	return f, nil
}
