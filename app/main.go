package main
/*
 * Receipt
 *
 * API version: 2.0.0
 * Contact     : dikhi.martin@gmail.com
 */
 
import (
	"./routes"
	lib       "./lib"
)
var logs 		= lib.RecordLog("SYSTEMS -")

func main() {
	e := routes.Index()
	logs.Println("Starting Application "+ lib.GetEnv("APP_NAME"))
	e.Logger.Fatal(e.Start(":"+ lib.GetEnv("APP_PORT")))
}

