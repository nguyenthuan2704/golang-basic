package main

import "fmt"

func main() {
	fmt.Println("Hello, Gopher")
	const name, age = "Stephen", 29
	fmt.Printf("%v is %v years old. \t and the type is %T and %T", name, age, name, age)
}
