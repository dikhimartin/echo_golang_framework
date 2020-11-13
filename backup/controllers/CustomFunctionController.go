package controllers

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"strconv"
	"time"
	"encoding/hex"
	"encoding/json"
	lib      "../../lib"
)


// ## Define Config Variable Global
var logs 		  			= lib.RecordLog("SYSTEMS -")
var redis_connect 			= lib.RedisConnection()

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
