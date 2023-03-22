package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func (l *LogClient) Debug(v ...interface{}) {
	l.setPrefix(DEBUG)
	logger.Println(v...)
}

func (l *LogClient) Info(v ...interface{}) {
	l.setPrefix(INFO)
	logger.Println(v...)
}

func (l *LogClient) Warn(v ...interface{}) {
	l.setPrefix(WARNING)
	logger.Println(v...)
}

func (l *LogClient) Error(v ...interface{}) {
	l.setPrefix(ERROR)
	logger.Println(v...)
}

func (l *LogClient) Fatal(v ...interface{}) {
	l.setPrefix(FATAL)
	logger.Fatalln(v...)
}

func (l *LogClient) setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger = log.New(l.file, DefaultPrefix, log.LstdFlags)
	logger.SetPrefix(logPrefix)
}
