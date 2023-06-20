package log

import (
	"fmt"
	"go.uber.org/zap"
	"path"
	"testing"
)

func TestSetLogger(t *testing.T) {
	opt := Options{
		FilePath:   path.Join("../../", "log"),
		FileName:   "api-server.log",
		MaxSize:    512,
		MaxBackups: 10,
		MaxAge:     365,
		LogLevel:   DebugLevel,
		Console:    true,
	}
	if err := InitLog(&opt); err != nil {
		t.Fatal(err)
	}

	Info("this is a warn log")
	Debug("this is a debug log")
	Warn("this is a warn log")
	Error("this is a error log", zap.Error(fmt.Errorf("this is a error")))

	Infof("this is a warn log, the content is [%s]", "info content")
	Debugf("this is a debug log, the content is [%s]", "debug content")
	Warnf("this is a warn log, the content is [%s]", "warn content")
	Errorf("this is a error log, the content is [%s]", "error content")
}
