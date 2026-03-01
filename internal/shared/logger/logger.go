package logger

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	once        sync.Once
)

func Init() {
	once.Do(func() {
		InfoLogger = log.New(
			gin.DefaultWriter,
			"[INFO] ",
			log.Ldate|log.Ltime,
		)

		ErrorLogger = log.New(
			gin.DefaultErrorWriter,
			"[ERROR] ",
			log.Ldate|log.Ltime|log.Lshortfile,
		)
	})
}

func Info(v ...interface{}) {
	InfoLogger.Println(v...)
}

func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}

func Error(v ...interface{}) {
	ErrorLogger.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}
