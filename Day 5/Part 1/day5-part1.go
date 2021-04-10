package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func threeVowels(currStr string) bool {
	vowelCount := 0

	for j := 0; j < len(currStr); j++ {

		if currStr[j] == 'a' || currStr[j] == 'e' || currStr[j] == 'i' || currStr[j] == 'o' || currStr[j] == 'u' {
			vowelCount++
		}
	}

	return vowelCount >= 3
}

func doubleLetters(currStr string) bool {
	count := 0

	for i := 0; i < (len(currStr) - 1); i++ {

		if currStr[i] == currStr[i+1] {
			count++
		}
	}

	return count >= 1
}

func noForbiddenChars(currStr string) bool {
	if strings.Contains(currStr, "ab") || strings.Contains(currStr, "cd") || strings.Contains(currStr, "pq") || strings.Contains(currStr, "xy") {
		return false
	}
	return true
}

func main() {
	strList, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	niceStrings := 0

	for i := range strList {

		hasThreeVowels, hasDoubleLetters, hasNoForbiddenChars := false, false, false

		currStr := strList[i]

		hasThreeVowels = threeVowels(currStr)
		hasDoubleLetters = doubleLetters(currStr)
		hasNoForbiddenChars = noForbiddenChars(currStr)

		if hasThreeVowels && hasDoubleLetters && hasNoForbiddenChars {
			niceStrings++
		}

	}

	fmt.Println(niceStrings)

}
