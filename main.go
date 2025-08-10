package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/debug", DebugHandler)

	port := "8080"
	fmt.Println("Starting server on port:", port)
	err := http.ListenAndServe(":"+port, nil) // No TLS
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
