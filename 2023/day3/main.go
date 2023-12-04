package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ProcessNeighbor(
	char rune,
	currentNumber *string,
	sum *int,
	starHistory *map[string]int,
	symbolCoord string,
) {
	isSymbol := (unicode.IsPunct(char) || unicode.IsSymbol(char)) && char != '.'

	if isSymbol {
		number, _ := strconv.Atoi(*currentNumber)
		if char == '*' {
			if value, ok := (*starHistory)[symbolCoord]; ok {
				*sum += (value * number)
			}
			(*starHistory)[symbolCoord] = number
		}
		*currentNumber = ""
	}
}

func getGearSum(input string) int {
	sum := 0

	starHistory := map[string]int{}

	lines := strings.Split(input, "\n")
	for lineIndex, line := range lines {

		currentNumber := ""
		for start := 0; start < len(line); start++ {
			char := rune(line[start])

			isDigit := unicode.IsDigit(char)

			if isDigit {
				currentNumber += string(char)
			}

			if currentNumber != "" && (!isDigit || start == len(line)-1) {

				index := start - len(currentNumber) - 1
				if index < 0 {
					index = 0
				}

				if isDigit && start == len(line)-1 {
					index += 1
				}

				ProcessNeighbor(
					rune(line[index]),
					&currentNumber,
					&sum,
					&starHistory,
					fmt.Sprintf("%d:%d", lineIndex, index),
				)

				length := len(currentNumber)

				rightSymbolIndex := index + length + 1

				if index == 0 && start-len(currentNumber) == 0 {
					rightSymbolIndex -= 1
				}

				if rightSymbolIndex >= len(line) {
					rightSymbolIndex = len(line) - 1
				}

				ProcessNeighbor(rune(line[rightSymbolIndex]), &currentNumber, &sum, &starHistory,
					fmt.Sprintf("%d:%d", lineIndex, rightSymbolIndex),
				)

				if index == 0 {
					if line[index] == '.' {
						length += 2
					} else {
						length += 1
					}
				} else {
					length += 2
				}

				if lineIndex > 0 {
					prevLine := lines[lineIndex-1]
					for pos := index; pos < index+length; pos++ {
						if pos < len(prevLine) {
							ProcessNeighbor(
								rune(prevLine[pos]),
								&currentNumber,
								&sum,
								&starHistory,
								fmt.Sprintf("%d:%d", lineIndex-1, pos),
							)
						}
					}
				}

				if lineIndex < len(lines)-1 {
					nextLine := lines[lineIndex+1]
					if length > len(nextLine) {
						length -= 1
					}
					for pos := index; pos < index+length; pos++ {
						if pos < len(nextLine) {
							ProcessNeighbor(
								rune(nextLine[pos]),
								&currentNumber,
								&sum,
								&starHistory,
								fmt.Sprintf("%d:%d", lineIndex+1, pos),
							)
						}
					}
				}

				currentNumber = ""
			}
		}
	}
	return sum
}

//go:embed input.txt
var file string

func main() {
	fmt.Println("Day 3 Sum:", getGearSum(file))
}
