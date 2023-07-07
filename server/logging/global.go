package logging

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var RuntimeLog *logrus.Logger = logrus.New()

var HOME = "/Users/crystal/groot"

func init() {
	RuntimeLog.SetReportCaller(true)
	RuntimeLog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		// 添加调用者信息
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

}
