package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "leeroooooy app!!\n")
}

func main() {
	log.Print("Awesome go app ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50051", nil)
}