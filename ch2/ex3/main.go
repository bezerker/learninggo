package main

import "fmt"

func main() {
	var b byte = 10
	var smallI int32 = 20
	var bigI uint64 = 30

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println("b:", b)
	fmt.Println("smallI:", smallI)
	fmt.Println("bigI:", bigI)
}
