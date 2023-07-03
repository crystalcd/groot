package global

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	file, function, line := getFileFunctionLine(2) // 获取调用日志的文件名、方法名和行号
	threadID, goroutineID := getThreadGoroutineID()

	logLevel := strings.ToUpper(entry.Level.String()) // 日志级别
	message := entry.Message                          // 日志消息

	logLine := fmt.Sprintf("[%s] %s [Thread: %s, Goroutine: %d] [%s:%s:%d]\n", timestamp, logLevel, threadID, goroutineID, file, function, line)
	fullLog := logLine + message + "\n"

	return []byte(fullLog), nil
}

func getFileFunctionLine(skip int) (string, string, int) {
	pc, file, line, _ := runtime.Caller(skip)

	// 获取文件名
	fileName := file[strings.LastIndex(file, "/")+1:]

	// 获取方法名
	fullFnName := runtime.FuncForPC(pc).Name()
	fnName := fullFnName[strings.LastIndex(fullFnName, ".")+1:]

	return fileName, fnName, line
}

func getThreadGoroutineID() (string, int) {
	threadID := fmt.Sprintf("%x", runtime.Stack([]byte{}, false))
	goroutineID := getGoroutineID()

	return threadID, goroutineID
}

func getGoroutineID() int {
	buf := make([]byte, 64)
	runtime.Stack(buf, false)
	id := 0
	fmt.Sscanf(string(buf), "goroutine %d", &id)
	return id
}
