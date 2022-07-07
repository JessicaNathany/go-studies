package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RightTime(response http.ResponseWriter, request *http.Request) {
	getTime := time.Now().Format("2022-07-07 03:04:05")
	fmt.Fprintf(response, "<h2> Date time Now: %s</h2>", getTime)
}

func main() {
	http.HandleFunc("/rightTime", RightTime)
	log.Println("Executing....")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// test in browser http://localhost:3000/rightTime
