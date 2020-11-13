package main
/*
 * Receipt
 *
 * API version: 1.0.0
 * Contact     : dikhi.martin@tog.co.id
 */
 
import (
	"./routes"
	lib       "./lib"
)
var logs 		= lib.RecordLog("SYSTEMS -")

func main() {
	e := routes.Index()
	logs.Println("Starting Application "+ lib.GetEnv("APP_NAME"))
	// http
	e.Logger.Fatal(e.Start(":"+ lib.GetEnv("APP_PORT")))
	// https
	// e.LogFatal(e.StartAutoTLS(":443"))
}

