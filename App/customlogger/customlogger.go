package customlogger

import (
	"log"
	"os"
	"sync"
	"time"
)

type logger struct {
	filename string
	*log.Logger
}

var logg *logger
var once sync.Once
var user string

func SetUser(username string) string {
	user = username
	return user
}

func GetUser() string {
	return user
}

func GetInstance(userlog string) *logger {
	SetUser(userlog)
	once.Do(func() {
		logg = createLogger("logs/martin.log", GetUser())
	})
	return logg
}

func createLogger(fname string, user string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	return &logger{
		filename: fname,
		Logger:   log.New(file, time.Now().In(location).Format(time.RFC3339)+" "+GetUser()+" ", log.Lshortfile),
	}
}

func Debug(cmd string) bool {
	GetInstance(GetUser()).Println(cmd)
	return true
}
