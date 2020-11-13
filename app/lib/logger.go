package lib

import (
	"log"
	"os"
	"sync"
	"time"
)

var logs = RecordLog("SYSTEMS -")
var logg *logger
var once sync.Once
var user string

type logger struct {
	filename string
	*log.Logger
}

func GetUser() string {
	return user
}

func SetUser(username string) string {
	user = username
	return user
}

func RecordLog(userlog string) *logger {
	SetUser(userlog)
	once.Do(func() {
		logg = createLogger("logs/receipt.log", GetUser())
	})
	return logg
}

func createLogger(fname string, user string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return &logger{
		filename: fname,
		Logger  :   log.New(file, time.Now().Format(time.RFC3339)+" "+GetUser()+" ", log.Lshortfile),
	}
}

func Debug(cmd string) bool {
	RecordLog(GetUser()).Println(cmd)
	return true
}
