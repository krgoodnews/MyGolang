package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
func main() {

	fmt.Println("-------------")
	fmt.Println("숫자 게임에 오신걸 환영합니다.")
	fmt.Println("0~99 사이의 임의의 숫자를 맞추시면 승리합니다.")
	fmt.Println("-------------")

	// 랜덤 숫자를 추출함
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(100)

	tryCount := 0

	// 입력을 받음
	for {

		fmt.Printf("숫자값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else { // 성공적으로 입력함
			tryCount++

			if n > answer {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < answer {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("정답입니다. 시도한 횟수:", tryCount)
				break
			}

		}
	}
}
