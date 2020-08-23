package main

import "fmt"

var x = 4

func main() {
	fmt.Printf("Test\n")
	f()

}

func f() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hi ")
	}
}
