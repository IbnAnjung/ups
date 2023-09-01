package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		isFizz := i%5 == 0
		isBuzz := i%3 == 0
		// fmt.Print(i, " => ")
		if isFizz && isBuzz {
			fmt.Println("FizzBuzz")
		} else if isBuzz {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}
}
