package main

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}