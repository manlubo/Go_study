package main

import "fmt"

func main() {
	
	fmt.Println(10)
	s := Student{"John", 20,}
	var stringer Stringer

	stringer = s
	
	fmt.Printf("%v\n", stringer.String())
}
