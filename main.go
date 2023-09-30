package main

import "fmt"

type Person struct {
	name string
	age  int
	ch   byte
}

func main() {
	fmt.Println(&Person{ch: 25})
}
