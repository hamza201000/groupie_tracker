package main

import (
	"fmt"
	"log"
	"net/http"

	GroupieTracker "GroupieTracker/helper"
)

func main() {
	http.HandleFunc("/", GroupieTracker.Handler)
	http.HandleFunc("/static/", GroupieTracker.HandleStatic)
	log.Println("Server running on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
