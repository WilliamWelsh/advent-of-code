package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	powerSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, ": ")

		rounds := strings.Split(data[1], "; ")

		possibleGame := true
		lowestAmounts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range rounds {
			records := strings.Split(round, ", ")

			for _, record := range records {
				parts := strings.Split(record, " ")

				number, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Println("Error converting number:", err)
					return
				}

				color := parts[1]

				// Check if this value is over the max set in the config
				if value, ok := config[color]; ok {
					if number > value {
						possibleGame = false
					}
				}

				// Check if this value is lower than the lowest set in the config
				if value, ok := lowestAmounts[color]; ok {
					if value == 0 || number > value {
						lowestAmounts[color] = number
					}
				}
			}

		}

		if possibleGame {
			id, err := strconv.Atoi(strings.Replace(data[0], "Game ", "", -1))
			if err != nil {
				fmt.Println("Error converting number:", err)
				return
			}
			sum += id
		}

		powerSum += (lowestAmounts["red"] * lowestAmounts["green"] * lowestAmounts["blue"])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Part 1 - Possible Games ID Sum:", sum)
	fmt.Println("Part 2 - Lowest Amounts Power Sum", powerSum)
}
