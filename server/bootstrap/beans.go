package bootstrap

import (
	"fmt"
	"path"
	"runtime"

	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
	Pool   *ants.Pool
)

func InjectBeans(env *Env) {
	Logger = NewLogger(env)
	Pool = NewPool(env)
}

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

func NewPool(env *Env) *ants.Pool {
	pool, err := ants.NewPool(env.AsyncPoolCount)
	if err != nil {
		Logger.Fatal(err)
	}
	return pool
}