package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const file = "./soal_1.txt"

func main() {
	file, err := os.Open(file)
	if err != nil {
		log.Println("fail to read file ", err.Error())
		return
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	n := 0
	sum := 0

	for fileScanner.Scan() {
		text := fileScanner.Text()

		i, err := strconv.Atoi(string(text))
		if err != nil {
			log.Println("line must containe a number")
		} else {
			n++
			sum += i
		}
	}

	fmt.Printf("Total angka pada file: %d\n", n)
	fmt.Printf("Jumlah semua angka: %d\n", sum)

}
