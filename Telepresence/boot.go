package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Alive.")
}

func main() {
	http.HandleFunc("/hc", healthCheck)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
