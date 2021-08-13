package main

import (
	"fmt"

	"github.com/krgoodnews/BankApp/mydict"
)

// func main() {
// 	account := accounts.NewAccount("국훈")
// 	account.Deposit(10)
// 	// err := account.Withdraw(5)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	fmt.Println(account.String())
// }

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "FIRST")

	err := dictionary.Update(baseWord, "SECOND")
	if err != nil {
		println(err)
	}
	dictionary.Delete(baseWord)
	word, err2 := dictionary.Search(baseWord)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(word)
	}

}
