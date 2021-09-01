package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func MakeWebHandler() http.Handler {
	// ServeMux 인스턴스 이용하기
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler)
	return mux
}

func main() {
	// http.HandleFunc("/bar", barHandler)
	// http.ListenAndServe(":3000", nil)

	http.ListenAndServe(":3000", MakeWebHandler())

	/*
		// "/" 경로에 있는 파일 읽어오기
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		http.ListenAndServe(":3000", nil)
	*/
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")

	if name == "" {
		name = "NoName"
	}

	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! / id: %d", name, id)
}
