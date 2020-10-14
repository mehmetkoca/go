package main

import (
	"fmt"
)

// Hello returns hello message
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
