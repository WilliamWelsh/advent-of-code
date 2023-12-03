// https://adventofcode.com/2023/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("Day 1")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sum := 0

	numMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		digits := ""

		// Find the first word number starting from the front
	FirstLoop:
		for start := 0; start < len(line); start++ {
			char := rune(line[start])

			if unicode.IsDigit(char) {
				digits += string(char)
				break
			}

			for end := start; end < len(line); end++ {
				substr := line[start:end]

				if value, ok := numMap[substr]; ok {
					digits += value
					break FirstLoop
				}
			}
		}

		// Find the first word number starting from the end
	SecondLoop:
		for start := len(line) - 1; start >= 0; start-- {
			char := rune(line[start])

			if unicode.IsDigit(char) {
				digits += string(char)
				break
			}

			for end := len(line); end >= start; end-- {
				substr := line[start:end]

				if value, ok := numMap[substr]; ok {
					digits += value
					break SecondLoop
				}
			}
		}

		if digits != "" {
			number, err := strconv.Atoi(digits)
			if err != nil {
				fmt.Println("Error converting string to number:", err)
				return
			}
			sum += number
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Sum:", sum)
}
