package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

// control flow:
// (1) sequential
// (2) loop/iterative
// (3) conditional
