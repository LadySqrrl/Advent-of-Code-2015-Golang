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
		vowelCount := 0

		if currStr[j] == 'a' || currStr[j] == 'e' || currStr[j] == 'i' || currStr[j] == 'o' || currStr[j] == 'u' {
			vowelCount++
		}
	}

	if vowelCount < 3 {
		return false
	}

	return true
}

func doubleLetters(currStr string) bool {
	count := 0

	for i := 0; i < len(currStr); i++ {

		if currStr[i] == currStr[i+1] {
			count++
		}
	}

	if count < 1 {
		return false
	}

	return true
}

func forbiddenChars(currStr string) bool {
	if strings.Contains(currStr, "ab") || strings.Contains(currStr, "cd") || strings.Contains(currStr, "pq") || strings.Contains(currStr, "xy") {
		return true
	}
	return false
}

func main() {
	strList, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	niceStrings := 0

	for i := range strList {

		hasThreeVowels, hasDoubleLetters, hasForbiddenChars := false, false, true

		currStr := strList[i]

		hasThreeVowels = threeVowels(currStr)
		hasDoubleLetters = doubleLetters(currStr)
		hasForbiddenChars = forbiddenChars(currStr)

		if hasThreeVowels && hasDoubleLetters && !hasForbiddenChars {
			niceStrings++
		}

	}

	fmt.Println(niceStrings)

}
