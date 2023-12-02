package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {

	var totalCount int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number := processLine(line)
		totalCount += number
	}

	fmt.Println(totalCount)
}

func processLine(line string) int {
	var numberStr string

	for _, k := range line {
		if !unicode.IsLetter(k) {
			numberStr = string(k)
			break
		}
	}

	for i := range line {
		v := line[len(line)-i-1]
		if !unicode.IsLetter([]rune(string(v))[0]) {
			numberStr += string(v)
			break
		}
	}

	if len(numberStr) == 1 {
		numberStr += numberStr
	}
	number, _ := strconv.Atoi(numberStr)

	return number
}
