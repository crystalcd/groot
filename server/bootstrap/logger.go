package bootstrap

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func NewLogger(env *Env) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		// 添加调用者信息
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	return logger
}
