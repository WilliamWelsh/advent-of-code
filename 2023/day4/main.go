package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func getMatchCount(lineIndex int, line string) int {
	matchCount := 0
	if line == "" {
		return 0
	}
	data := strings.SplitN((strings.SplitN(line, ": ", 2)[1]), " | ", 2)

	winningNumbers := strings.Fields(data[0])
	myNumbers := strings.Fields(data[1])

	for _, number := range myNumbers {
		for _, winningNumber := range winningNumbers {
			if number == winningNumber {
				matchCount++
			}
		}
	}

	return matchCount
}

func getScratchcards(input string) int {
	scratchCards := 0
	lines := strings.Split(input, "\n")

	for lineIndex, line := range lines {
		fmt.Println("\nLine:", line)
		matchCount := getMatchCount(lineIndex, line)

		if matchCount == 0 {
			scratchCards += matchCount
			fmt.Println("Won", matchCount, "cards")

			dupeLines := []string{}

			for {
				for i := lineIndex + 1; i < lineIndex+matchCount+1; i++ {
					dupeLines = append(dupeLines, lines[i])
				}

				for dupeIndex, dupeLine := range dupeLines {
					fmt.Println("Dupe Line:", dupeLine)
					dupeLines[dupeIndex] = ""
					matchCount := getMatchCount(lineIndex, dupeLine)
					scratchCards += matchCount
					fmt.Println("Dupe Match Count:", matchCount)

					for i := lineIndex + 1; i < lineIndex+matchCount+1; i++ {
						dupeLines = append(dupeLines, lines[i])
					}

				}

				nonEmptyLines := []string{}
				for _, line := range dupeLines {
					if line != "" {
						nonEmptyLines = append(nonEmptyLines, line)
					}
				}
				dupeLines = nonEmptyLines

				if len(dupeLines) == 0 {
					break
				}
			}

		}
	}

	fmt.Println("Scratchcards", scratchCards)
	return scratchCards
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 4 Sum:", getScratchcards(input))
}
