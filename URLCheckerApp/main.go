package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := map[string]string{}
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		// fmt.Println(url)
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
	fmt.Println("DONE!!")

}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}

// ---------------------- Goroutine 예제
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// go sexyCount("국훈")
// 	// go sexyCount("Tim")
// 	c := make(chan string)

// 	people := [6]string{"국훈", "팀국", "클레이튼", "제이크", "앨리슨", "제임스"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	// time.Sleep(time.Second * 10)

// 	for i := 0; i < len(people); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 2)
// 	c <- person + "is sexy"
// }

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, " is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }
