package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprint(w, "Hello World\n")
	})
	port := 3000
	log.Printf("HTTP Server listening on port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
