package main

import (
	"bufio"
	"fmt"
	"os"
)

func twoSetsDoubleLetters(s string) bool {
	doubleLetters := 0

	return doubleLetters >= 1
}

func oneLetterBetweenMatching(s string) bool {
	oneBetween := 0

	return oneBetween >= 1

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	strList, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	niceStrings := 0

	for i := range strList {

		currStr := strList[i]

		hasTwoSetsDoubleLetters := twoSetsDoubleLetters(currStr)
		hasOneLetterBetweenMatching := oneLetterBetweenMatching(currStr)

		if hasTwoSetsDoubleLetters && hasOneLetterBetweenMatching {
			niceStrings++
		}

	}

	fmt.Println(niceStrings)

}
