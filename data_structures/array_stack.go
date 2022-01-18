package data_structures

import "fmt"

/*
	Implementation of the Stack data structure using Golang
	Methods:
	+ Push
	+ Pop
	+ Peek
	+ Clear
*/

type ArrayStack[T any] struct {
	items []T
}

func PrintHello() {
	fmt.Println("Hello Data Structures")
}
