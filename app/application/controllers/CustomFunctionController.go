package controllers

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
	"encoding/hex"
	"encoding/json"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
	lib      "../../lib"
)

// ## Define Config Variable Global
var (
	logs 	  			= lib.RecordLog("SYSTEMS -")
	paginator 			= &pagination.Paginator{}
	redis_connect 	    = lib.RedisConnection()
)


// ## Define Type Global
type response_json map[string]interface{}

// function
func NewSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func ConvertToMD5(value string) string{
	var str string = value
	hasher := md5.New()
	hasher.Write([]byte(str))
	converId := hex.EncodeToString(hasher.Sum(nil))

	return converId
}

func ConvertToSHA1(value string) string{
    sha := sha1.New()
    sha.Write([]byte(value))
    encrypted       := sha.Sum(nil)
    encryptedString := fmt.Sprintf("%x", encrypted)
	return encryptedString
}

func ConvertStringToInt(value string) int{
	value_int, _  	:= strconv.Atoi(value)
	return value_int
}

func ConvertStringToFloat(value string) float64{
	value_float, _ 	:= strconv.ParseFloat(value, 8)
	return value_float
}

func ConvertJsonToString(payload interface{}) string{
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logs.Println(err)
	}
	return string(jsonData)
}

func current_time(format string) string{
	current_time := time.Now().Format(format)
	return current_time
}

func FormatDate(rec, format string) string{
	date,_       := time.Parse("2006-01-02 15:04:05", rec)
	conv_date    := date.Format(format)
	return conv_date
}

