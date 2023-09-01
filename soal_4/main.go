package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// numLength := len(numbers)
	mapIndex := map[int64]bool{}

	for {
		if len(mapIndex) == 10 {
			break
		}

		random := rand.Intn(10)
		if _, ok := mapIndex[int64(random)]; ok {
			continue
		} else {
			mapIndex[int64(random)] = true
			fmt.Print(numbers[random], " ")
		}

	}
}
