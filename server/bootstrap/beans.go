package bootstrap

import (
	"fmt"
	"path"
	"runtime"

	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var (
	Logger     *logrus.Logger = logrus.New()
	Pool       *ants.Pool
	DomainPool *ants.Pool
	PortPool   *ants.Pool
	HttpPool   *ants.Pool
)

func InjectBeans(env *Env) {
	Logger = NewLogger(env)
	Pool = NewPool(env)
	DomainPool = NewDomainPool(env)
	PortPool = NewPool(env)
	HttpPool = NewPool(env)
}

func NewLogger(env *Env) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
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

func NewDomainPool(env *Env) *ants.Pool {
	pool, err := ants.NewPool(env.AsyncPoolCount)
	if err != nil {
		Logger.Fatal(err)
	}
	return pool
}

func NewPortPool(env *Env) *ants.Pool {
	pool, err := ants.NewPool(env.AsyncPoolCount)
	if err != nil {
		Logger.Fatal(err)
	}
	return pool
}

func NewHttpPool(env *Env) *ants.Pool {
	pool, err := ants.NewPool(env.AsyncPoolCount)
	if err != nil {
		Logger.Fatal(err)
	}
	return pool
}
