package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   // Student 객체를 []byte로 전환
	w.Header().Add("content-type", "application/json") // JSON 포맷임을 표시
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	// fmt.Println("Hello, student!")
	http.ListenAndServe(":3000", MakeWebHandler())
}
