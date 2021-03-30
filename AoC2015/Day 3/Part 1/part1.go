package main

import (
	"fmt"
	"strconv"
)

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func main() {
	path := ">^^v^<>v<<<v<v^>>v^^^<v<>^^><^<<^vv>>>^<<^>><vv<<v^<^^><>>><>v<><>^^<^^^<><>>vv>vv>v<<^>v<>^>v<v^<>v>><>^v<<<<v^vv^><v>v^>>>vv>v^^^<^^<>>v<^^v<>^<vv^^<^><<>^>><^<>>><><vv><>v<<<><><>v><<>^^^^v>>^>^<v<<vv^^<v<^<^>^^v^^^^^v<><^v><<><^v^>v<<>^<>^^v^<>v<v^>v>^^<vv^v><^<>^v<><^><v^><><><<<<>^vv^>^vvvvv><><^<vv^v^v>v<<^<^^v^<>^<vv><v<v^v<<v<<^^>>^^^v^>v<><^vv<<^<>v<v><><v^^><v<>^^>^^>v^>^<<<<v><v<<>v><^v>^>><v^^<^>v<vvvv<>>>>>^v^^>v<v<^<vv>^>^vv^>vv^^v<<^<^^<>v>vv^v>><>>>v^>^>^^v<>^<v<<>^vv>v^<<v>v<<><v>^vvv<v<vvv^v<vv<v^^^>v><<^<>><v^^>^v^>>^v<^<><v<>>v^<>>v<>>v^^^><^>>vvvv>^v<^><<>>^<>^>vv><v<<>>^^>v^^^><^<<^^v>v<^<<>v>^^vvv^v^>v^<>^^<>v^v>v>v<v^>vv>^^v<>v>>^<>><>v>v^<<vvvv<vvv><v^<^>^v<>>^><v>><>^<v>v<v>vv^>>vvv<>v>v<v^>>^>>v<<>^<>^<>>>^v<<<^<^v>vv^>><<><v^>^v^^^v<>^^vv><>><>>^>v^<v<>v<>>^<<^v>^^^<>^v^><>v<<v>vv^>vv<<>>><<^v^<>v<vv>>>^^<>^><<^>vv>>^<<v^^vv<>>><v>v><^<v<<>>>^^<>>^<^v><>vv^^^v>vvv>^><<>^^>^<<v^<v<^v<<>vvv<^<<>^>^v<vv<^>vvv>v>vv^<v^><>>^vv<^^^vv><^vv<v^<><v^vvv><<^>^^><v<<vv^>v<vv<v>^<>^v<<>v<v^v^>^>^>v<<^vvv<<<v>^^>^<<<<>vv>>^<>^>>>v<v>^^<v^<v<>>>vv>^^v<<>>>^^v><<<v<v<^v<>^^><v<^v<<v^><><^<><v<^^v>>><v^^v<<v^><^<><<v^>><^<>v>v^<><^<v>^v^>^>^vv^>^^<<vv^>vv<^vvv<>>^^<^>v^>^>^<v^><v<v>>>v<<<><^v<<><^<vv^v^^^>v<^^<v^vvv<v<><v<vv<^vv<>vv<v^<>>vvvvv<<>^v^v>vv>>>vvv^^<^<^<><>v<v>><^v><^<<<>><<<v>^>v<>^>^v>>^<>v^<^>><<>^<v>^>^^^>^^<v>>>><>^v^v><<<<vv^<vv<>vv>v<>v^<v^>v><>>>v^<><^vvv>vv^<^<<^<^^v>^>>>v<^<^v^^<^<^>>><v>vv>^<<><>^>>v>^<<>><^<>v<>vv^^>^>vvv^v<<^<^^<vv<>^vvv<^^v^vv^>>v<^>^^<v^<>v<^<^vv>v<<vv>vv>^>vvv>>>^^>v<>^v>v^<^>>v>^^v>>>>v^<v>v<^>v<v<<>>^v<^^<v><^<>>^<<vv^>>v<<v>^v<>><^>vv<v<^>>^^<vvvvvvvvv>>>v<v<>v^<>>^vv<v^^v<<^vvv^<<^><>vv<><<>>v>vv^><>>^^v^>>v^v^><<<>>^^<^v<<^<>>>>^<^>v^><<^>v<^v<^>>^^<<<<><^<^v^v<>>^v<^<<vv^<><^^vv><v^v^v>^>>^>^vv^>^v<v^v<<vvv^><>>^v^^><>v>vv><^>>vv<vvv<<<<^<>vvv^v<v>^<v<^>^<^<v<><>v^^^^<<vv<^^vv<v>><<v^><>>><v^>^v><^>^><vv^<><^<v>><<^vv<>>v^<<v<>v><v<><><vv>^>>v^<^<v>^><>>><^><v^v<>>>^^<^>v<v>vvv<>^<<><v^^>^>>v<^v>^>v>>>vv>v>>v^^^<^<vvv^<>^>^<v^<v^v>v>^>vv>vvv<>v<^>v>^^>>^<vv^^v>v^^^^^v^vv><^<><>^>vv<^>>^vvvv^^^>^<vv>^v<<^><^^>^<>^^>^<<v<^>>>^><<^^>v^v>>^>vvvv>^^v><v>>vv><<<vv<^>v>^^^<v>v^vvv<^><<^>^<>^><<<<<v^<<vv^v>^<>v<v>^>^>><>v^v<^vv^^>vv<<v^v>vv^vvv<<<<>^v<v^^v^v>v<<v>^^<>^vv^^>^>^v^vv^>>v^vv^^<vv><<v^v^^v><vv<^vvv<vv^^<<v>v^v^^^^v<^<^>v>^>v>^vv^v^^<v<^vvvv<<<>^<^^^<^^<>^<><vv<^^<<^>>><v^vvvv>^<>>^^>v^^v^<<v^^^<<<><^<v^v^^v<v^<>v><<v<>^v>v<^><^>vv^^<vvv<^v>>v>^<><v^><^^^<v^>>vv<<<<<^<>^v^v>^vv^<>v>v<^>vv<<^vv>vv<v<><>>v>><v<^<^^>><<v^v<<^><v<^<vv<v<<vv^>^<<><^^>^<^>>^<vv>><v<<vvv<^^v^>^^<^v>^v<v<>v><v^v^<<^<><<v<<^v>v<<>>^>v>>v>>v<^<<^<^>>>v>^^^v><^>^^>>v<<>^v><v>vvv^vv<<<>vvv<<>^>>>v<v<v^<^<^>^<^>v^^v<^^<v<>v<>>^^>^v^>v<<<<^<>v^><<<v>>>><<v^<^vv>v>><>>^<<<^<^^>v<>>v<>vv<<^<<><<^>v^^^vv^>vvvv>>v>v^><<v<>vv^<<><<vvv>^>>>^<<<^<^<<v>^>v<>>v>>vv^^><<<<^^^v>><<^><v><v^^><v<<v^^v^^v>>v<><><<>^><v><^<vv>><^v<>v<vvv<>^>><v>>v<^><<v>^<>^v><^><^^<v>^><^^v^<<><>>^>v^<^v^vv<><^>vv^>v^vvv^<>>^><^<^<>^<<v^v<^v><>^v<v>>^>>^v^vv>><vv><v^^<<^v^<>^v<<>^><^>><v>>v<<<v^^vv<>^^v>>><><><<v^<<<v^<^^><v^>v^^vv<v^<>>vv^<^v<>^v>>v^v>v<^^vv><>^v<<>v^<>v^>>v>vvv<^><><^^>^vv^>>v^>^<^^<><>><<>^^^><^v^v><<<><<^v^vv>v>><^>>><v^>v<v><><v^v<>v^^>>v<<>v>v<v<v<^^<><>v^^<>>v<^v<v>v<><v<v>^<<>v>vv^^<>>^^^<>^^>^v>v>>>^v^v><v^^<><v>^^v^v<^<^^><<v<^<^<>^<>><<>^>>^>^^><v><>v<><>><<<>>>>vv>>>^>>^v<^>v^^^v<<vv>><<<^<<<>>>>>^>vv<^v^<>^<v^>^v><v>vvv<>>>^v^^^v<<<<>>^^<vv<^<^^>^<>v<^<<<>><>>v<^<>^<vvv<^<>><><<v>^^^>^^<<v<v^>^^v^>><<^vv><v>^v>>^<v>v>^^>^v>^vvv<>v^v^^<><vv>vv^>>><>v<^><v<v^<><<<>^v>^v<<<^>^>^>v^v<<><vvv<<v^^<><v>^>>><vv>><v>>v^<vv>>vv<<^v^v<<><^v<vv>>>vv<>>>>^vv>v^<>vv>v^v<v^><v<^^^^^>vv<><<vvv^<v><^<vv><^^^vv^<>^^^^<^><^<>v^<v^v<<^v<<^^<>>^<v^^>>>vv<vvv<>v<<>><^vvv^<<^^<<>>>^<>>>v^^><>><<>><v^v>>>>>><>>><v^<<vvv^>v<>>v^<>vv<><^^^^v^<<^<v^vv><<^^>v<^vvv^v>>v>^>>v>^^><<v^<>v<>vv<^v^vv><v><<vv^v>>v^>>v<^^^>^><<v<>^><>v>>>vvv<v<vv<^>>^v<v>^<^^^^^v><>v><>v^v^v<v^vv^v>vvvv<>vv<<<vv<v<<>^<^>^^v^<<>^<v><^><v<v<><<>v^<<^<><vv>v<<^v>>^v<><v>^>>^^><>v^<^<vvv^>^>^<<<<>vv>^v^v<^^^<vv>><>^^<<v<^<^^>>>v^v<<^^^<v<v<^<>^v<v><v^vv^^v^^v^^<vv<>^<><vv^<^v^<<^><<vvv>^^<^^^<^v>^>^vv><<<^v<v>vv>v<>v^v<v^>v^>>>v^v<>^v<<>^vv>v>v>v^<^>v^^<^>^^^^vv>^^><^>vv^>>^^v>><<<<^><>v<>^<v<vv^>^^><<^><v>v^>^^<^>>><>><v^v<v^<v<vv^v^<<^<vvv>>><vv<^^>>^>^><<v^<>>v>v^v^^><<>vv^v>v^<v><^<>^^<^>v>^<><<<v>^<^<^>^>^>^^v^<<^^v^^<^<>><^>v>>^^<>^^^<<<<v^>^v<^vv>^<<<v<><<v<>vv>>>v><>>><>>v<<<vv><>^v>v<^>><^><><v<>^v^>^v>^v<<><<^<>>v>^><>^>><>><^<v^><v^^<><v><^^>^v^^<>v^<v^<^v<v^^^^^v^<<^>^^^<^v><>^^<<<><<<<<^^>v^vvvv>v<>>vv<^>^v^>v<^vv^v<<><<v>v^v>^^><><^<v^>v><vv><>>><<>^vv<>v>>v<^v>>>v<v>v>v>^vv<<>^^vv<v<^v^<v<v>vv<>^<^<vv<v^<^v^^><<>^>><^v>vv^^v<<^^><<>v^^<><><v^^<v^v>^>^>^>v<^<v>^v^^>v<>vvv<^v<v^v><<v^><<^^><^<<v^v^>v<>^>v><><v>^<v<v>^<^^^>^v<<><<><>vv>v^<>v^><v^v<v><><<v>v<vv><<v>>v>^<<<>vv>>vvv>^^vv^v^^<^^<>v^^<>v>>^^>^>^>v>><^>><>>^<<>><^>v<<<<<<<^v^v<v^<v^^>^<><<v<^>v^>v^vv<<^^vv^>>>>^<>v<^v<>v<vv<^>>v^vv>vv><vv<<^>v>><vv>>>vv^<<<<vv^>v<<<<^^>^^v^><<^<v^>v^>^^<v<>vvv^>^<>vvv<v<^^>v^<<v>><>v<v<>^^<vvv>^>vv><><<<^^vv<v^<v<>v<>><<v><^vv^>^<^>^^^<<<v>vv^<^<<>^>^<vv>v><v<<^><^>^^<vv^v^^>>>>vv^><^^vv><>^<v^v>v<vv>v><<<v>v<v>^><v^^><v>v<^v^>>^^<v^>^^>vv>>vv^><^vv^vv<<^>vv>^v<v><vv><v<vvvvv>^^v^v><v>>>^vv<>v>^^^^<^>><>^v^^^>v<^^<<^^v<vv<>vvv<^>><><^>>^><^<>v<v<<><<v><v^v<>><^>v><<v^<v>v<^<vv^v^v^>vvv^^>v>^<vv^>v^v^<>v>^>>vv>><^^<v<<>^vv<><><<^v<v>v<<vv><>><^v<v>>v^>vvv^v^<<^><v<>^vv^>v^<v<^>>v<v><v><v>>^<<<v^<><<>v>^>^^<v<>>^<>^>^><<<^<<^<<^>^v>>><vvv>><<<<v>>>>>>>^<^v<^>v<>vv<><>v>>^>>^>vv^^><<^<v<v>>^^<<^>v<^>>vv>^<>v><^>v<vv>>>>>>^v<^<<<v^><vv<<>>vv<<><v<><<<v<^<v<>>v<^^^^v^^<^^^<^<vv><<^>><>v<<>v<v<>>>><>v^vv>^>^>>vv^v<v<<><^v>vv^><v<<>v^v<^>vv<<^^v><^>>^^vv<^<>>v^^>><v>^v>>>^>>v>v<>v<^vv><>^<<^>vv>>><><>v^><>v^>v>v><^v<><v<v>^v<<^vv^><^^>><^^^<<<^>v>^v>>><^>><^>>>^^^<^>vv<><<<v^>^<^^>>^^^v^v^v>v<v>>>><^>>>v>^vv<<^^^<^^vv>v<<><v<<^^>v>><<v^^><^>^<^>^v^>v><^<^vv>v>><>^<<vv<<v>v<vv<v>^>^>><^^<v>^v^v<><<>vvv<^<v>^><>^>vvv>>>^><<>><v^^<^<<^v>>^v<v<vv>vv^v^>v<<vvv<^^v^v>^<^>>^>v<^>^v<<><<<^>^<^^^>vv<^^^^vv<v<^^v<<<<v<^v^<><v<<^><<>vv>>><^<^<>>>^>^>>^<<<<<^^v>^>^<>vvv^^<^><^>^^v>^vv^><v^<^<<v^<vvv<<^v<><^><^>>>v>^v>^>^v<vv^v>><v><^><v^^>v^>^<><<><>v<v^>vvv^>^>>v<>^><^>^><vvv>^^v^v>v<>^v^><^>>v>v^><<<^>>^<>^<>>v><>>v^>^>^^<>>v^>^<vvvv<^vvvv^>>vv^<v^v>^vv<>v<>^<v<v>v>^^><^>vv^<^v^<<^<^<><vv<^v<^v><>>>^v^<<^><^>vv<v>v<^>vv^>v<<<>^<><v<^^^>v><^^<>^<^<v^vv^<<^>><<v^v<^vvv<<<>>vvvv^v^^^>v<>>><<>vvv<<^^^>v>v>>v<<v<v^v^>^^v>^><^<><<v^<v<v^^^><>v^^^<v>vv<>^>^^vv>^<<^v<^v><v>>>^>>><^<<>^v>>^>vv<<<v<>^<v><v^<^<>v>v^^v^>><<^v<<<<>v>v>v^^<^><>^^<<<v>vv<>>>^>>v<><v^>^<><vv>v>v^v<v^<^>>^>><<^^<^^v<vv<>><<<v<^<<^^^>vvv^<vvv<^>vv><>><<<^<v^v^^<<^vvv^^<^<><<>^<^<>>vvv<>^<>v^v<><>>v^v><<>>>vvv>v<>^>>^><^>vv<<>>v<<^><>v>>^^<v>^>^<<>><^<<vv<^<vv^vv><>>>><^<v>^>vv<v><>^<>vvvvv^vv<<v<>>>^<<><>^^vvv>>>vv<<^^><^v^^v<>^^>^><^>v^^^^v<^<<vv<vv<>vv^^>v^vv>v><>>vv>^<^<v^v^>>v^v^^v>^>vv^>v<vvvv<^v<^v>^v>^^v<<^>^^<<>^><^v>>>vv^>^^>vvvv>>v<^<v>^>>>v^<><^<^^<v>vv^^><v>v^<>^^^>>><^^v>v>^<<>^<v^>vvv^>^^^><v<^>>v<v>>^v><<><<>v<^<<>^><>^>vv>^<v>^^v<<^v^vvv^^>^vv^<^>^>^^v>v^>^<<><<^>v>>vv^vv><v>>^<<^<v^^<^<v^^vv^><^^<^^><v^^>v^^^<^<>^<>>^v<^vvv^^v^<><^>>>>>v><><<<>vv<^v>><<>vvv<><<vv<<<^>v^^>>^>^v>><><^^v<>><>>v^>^<vv><<<>><><<v>^^<>>v<><^<vv>vv<^v>^<<<<v<^<<^^>>^<><^>><<>^>v>^^^v>>^<^^v><v^v>^><<><>>^>>^<<v<>^v<>^>^<v>>vv>^vvv<<v<<^>^>^<<^^<>^^^^vvv<>^vv<vvvvv^^>^^<^>>><>v^<><^<<^>v^^v<>>^vv<>v^^<>>v^vvvvv<<v^<v^^>>><vvvvv>><^>vv>v^v^<v<^>^^><^>^^^^v<><^v<<>v^>v>>vv<<>^<v^^>vvv>^^<v^<>vv^><>><v^^v<>^>>^>v><>>^^v>^>^>>>^>v<^v>v>^<^^^^^>>v<v<>>v<<^>^<v<<>^^>><<^><>v<>^^^vv<>^^>><<^^>v>vv>vv>v^>^v>v^^<>>><<v><v<<>>v><>vvv^^v>^^>^vvvv^>^<>^vvvv><v><v<>>><>^<^vv<>^v<^v<>^vvv<<>><vvv^>>^><<vv^<v^>^<v<<^^>^^<^^v^>v<>v^v><>><v^^>>^vvv><^vv>v^<^<^v>>v^^>^vvv^<v^^v^^>v<^<>>^<>>>^^<><^^vv<>^vv^<>>>>^^<<^^<>vv^^><>^^<v<<v>^<v^^>^v<><><>vvv>^v^>>vv<<^v<<>><v>^><^>>>^<^<^^>vv^<<^<>>^^><><<v>^^<v>>v<<vvvv>^v^vv>><^^<<^>>v>v<^^^<^><^^vv>^vv<^<vv<>v><^<><v><^^^>>^<><^<v>>>>v^<v>>>>>v<><^^>v<^<^>><v<>^>vv>^^v^v^<<v<><<<^v^><<^<><<<<v<^>><<<>v>>vv><vv<><<^<^<><vv>^^^^<>v<<<<v>vv<>vv^^^>><>vv^><>>^vv<<><^^vv<>v^>>^<<>^<v^<^>v<"
	tempStr := ""
	xPos, yPos := 0, 0

	var coordsVisited []string
	coordsVisited = append(coordsVisited, "0, 0")

	for i := 0; i < len(path); i++ {

		switch path[i] {
		case '<':
			xPos -= 1
		case '>':
			xPos += 1
		case '^':
			yPos += 1
		case 'v':
			yPos -= 1
		}

		tempStr = strconv.Itoa(xPos) + ", " + strconv.Itoa(yPos)
		coordsVisited = append(coordsVisited, tempStr)

	}

	var minusDuplicates []string
	minusDuplicates = removeDuplicatesUnordered(coordsVisited)

	atLeastOnce := len(minusDuplicates)

	fmt.Println(atLeastOnce)
}
