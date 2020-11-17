package controllers

import (
	"fmt"
	"os"
	"unicode"
	"crypto/md5"
	"crypto/sha1"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
	"encoding/hex"
	"encoding/json"
	"github.com/labstack/echo"
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

func HashPassword(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
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


func RemoveFile(c echo.Context, path_file string) int{
	err := os.Remove(path_file)
	if err != nil {
		logs.Println(err)
		return 0
	}
	return 1
}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func FormFile(c echo.Context, value string) string{
	form, err := c.MultipartForm()
	if err != nil {
		logs.Println(err)
		return "nil"
	}
	files := form.File[value]
	if files == nil{
		return "nil"
	}

	file, _ 	   := c.FormFile(value)
	file_image, _  := file.Open()
	defer file_image.Close()

	timestamp 	   := time.Now().Unix()
	unix_timestamp := strconv.FormatInt(timestamp, 10)
	name_file 	   := file.Filename
	FileNamePost   := removeSpace(unix_timestamp + "_" + name_file)

	return FileNamePost
}

func MakeDirectory(folderPath string) string{
	// check_directory
    _, erot := os.Stat(folderPath)
    if os.IsExist(erot) {
    	logs.Println(erot)
        return folderPath
    }else{
	 	err := os.MkdirAll(folderPath, 0777)
	 	if err != nil{
	 		logs.Println(err)
	 		return "0"
	 	}
    }
	return folderPath
}
