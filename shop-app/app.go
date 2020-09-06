package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "shop app!!\n")
}

func main() {
	log.Print("shop app server ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50051", nil)
}
