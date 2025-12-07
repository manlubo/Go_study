package main

import "fmt"

func main() {
	var mainA *account2 = &account2{
		balance:   100,
		firstName: "John",
		lastName:  "Doe",
	}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}
