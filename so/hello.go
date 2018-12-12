package main

import "C"
import "fmt"

// Hello : Return "hello world"
//export Hello
func Hello() {
	fmt.Println("Hello world.")
}

func main() {}
