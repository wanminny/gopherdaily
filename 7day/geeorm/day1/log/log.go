package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	mu sync.Mutex

	infoLog = log.New(os.Stdout, "\033[34m[info:]\033[0m", log.Ldate|log.Ltime)
	errLog  = log.New(os.Stderr, "\033[31m[error:]\033[0m", log.Ldate|log.Ltime)

	loggers = []*log.Logger{infoLog, errLog}
)

var (
	Errorln = errLog.Println
	Errorf  = errLog.Printf
	Infoln  = infoLog.Println
	Infof   = infoLog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disable
)

func SetLoggerLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
	if ErrorLevel < level {
		errLog.SetOutput(ioutil.Discard)
	}
}
