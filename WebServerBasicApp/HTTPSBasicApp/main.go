package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello, HTTPS!")
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil)

	if err != nil {
		log.Fatal(err)
	}
}
