package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	layout12 := "03:04:05 PM"
	layout24 := "15:04:05"

	fmt.Print("Masukan Waktu dalam format 24 jam (HH:MM:SS AM/PM):")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Fail to read input", err)
		return
	}

	input = strings.Trim(input, "\n")
	input = input[0 : len(input)-1]

	t, err := time.Parse(layout12, input)
	if err != nil {
		fmt.Println("format tidak falid")
		return
	}

	fmt.Printf("Waktu dalam format 24 jam: %s\n", t.Format(layout24))
}
