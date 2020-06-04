package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomasen/realip"
)

func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	clientIP := realip.FromRequest(r)
	log.Println("GET / from", clientIP)
}

func main() {
	http.HandleFunc("/api/v0/ip", ServeIndexPage)
	fmt.Println("Listen And Serve...", "6000")
	err := http.ListenAndServe(":6000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
