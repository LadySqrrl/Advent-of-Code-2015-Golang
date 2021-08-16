package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func getAction(input string) string {
	action := ""

	if strings.Contains(input, "toggle") {
		action = "toggle"
	}

	if strings.Contains(input, "turn on") {
		action = "turn on"
	}

	if strings.Contains(input, "turn off") {
		action = "turn off"
	}

	return action
}

func getCoords(input string) string {

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	coords := reg.ReplaceAllString(input, " ")

	return coords
}

func main() {

	input := []string{"turn on 537,781 through 687,941", "toggle 223,39 through 454,511", "toggle 188,228 through 332,465"}
	var actions [300]string
	var coordStr [300]string

	for i := range input {
		currStr := input[i]
		actions[i] = getAction(currStr)
		coordStr[i] = getCoords(currStr)
	}

	var coordSplit []string
	var coords [1200]int

	for i := range coordStr {
		split := strings.Split(coordStr[i], " ")
		for j := range split {
			curr := split[j]
			coordSplit = append(coordSplit, curr)
		}
	}

	for i := 0; i < len(coordSplit); i++ {
		currStr := coordSplit[i]

		//fmt.Println(currStr)

		currInt, err := strconv.Atoi(currStr)
		if err != nil {
			fmt.Println("error")
		}

		coords[i] = currInt
	}

	//fmt.Println(coords)

	var lightArr [1000][1000]bool

	for i := 0; i < len(input); i++ {
		startX := coords[i+1]
		endX := coords[i+4]
		startY := coords[i+2]
		endY := coords[i+5]

		switch actions[i] {
		case "turn on":
			for k := startX; k < endX; k++ {
				for j := startY; j < endY; j++ {
					lightArr[k][j] = true
				}
			}

		case "turn off":
			for k := startX; k < endX; k++ {
				for j := startY; j < endY; j++ {
					lightArr[k][j] = false
				}
			}

		case "toggle":
			for k := startX; k < endX; k++ {
				for j := startY; j < endY; j++ {
					if lightArr[k][j] {
						lightArr[k][j] = false
					} else {
						lightArr[i][j] = true
					}
				}
			}

		}

		fmt.Println(lightArr)
	}

}
