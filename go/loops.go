package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	for sum < 1000 { // C's while is spelled for in Go
		sum += sum
	}
	fmt.Println(sum)
}
