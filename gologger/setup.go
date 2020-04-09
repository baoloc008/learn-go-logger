package main

import (
	"log"
	"os"
)

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./gologger/api.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}
