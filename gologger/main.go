package main

import (
	"log"
	"net/http"
)

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
	//log.Fatal("Fatal log") // ends program by calling os.Exit(1)
	//log.Panic("Panic log") //  throws a panic after writing the log message
}

func main() {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}
