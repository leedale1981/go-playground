package main

import "fmt"

func main() {
	m := map[int]string{1: "Test"}
	m[2] = "Test2"
	fmt.Println(m[2])
}
