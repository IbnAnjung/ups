package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Masukan kata atau kalimat: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Fail to read input", err)
		return
	}

	input = strings.Trim(strings.ToLower(input), "\n")
	input = input[0 : len(input)-1]

	lengthInput := len(input)
	for k, char := range input {
		if string(char) != input[lengthInput-k-1:lengthInput-k] {
			log.Println("bukan Palindrom")
			os.Exit(0)
		}
	}

	fmt.Print("Merupakan Palindrom")
	// remove the delimeter from the string
	// input = strings.TrimSuffix(input, "\n")
	// fmt.Println(input)
}
