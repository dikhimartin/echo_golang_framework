package logger

import (
	"log"
	"os"
)

func Init(l *log.Logger) {
	logFile, openErr := os.OpenFile("./test.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if openErr != nil {
		log.Println("Could not open log file!!")
	}

	defer logFile.Close()

	l = log.New(logFile, "", log.Lshortfile|log.Ldate|log.Ltime)
}
