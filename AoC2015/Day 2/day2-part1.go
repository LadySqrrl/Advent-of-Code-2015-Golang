//elves measuring wrapping paper
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {

		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	var wrappingPaper int

	for i := 0; i < len(nums); i += 3 {
		l, w, h, area1, area2, area3, smallest := 0, 0, 0, 0, 0, 0, 0

		l = nums[i]
		w = nums[i+1]
		h = nums[i+2]

		area1 = 2 * (l * w)
		area2 = 2 * (w * h)
		area3 = 2 * (l * h)

		if area1 <= area2 && area1 <= area3 {
			smallest = area1 / 2
		} else if area2 < area1 && area2 < area3 {
			smallest = area2 / 2
		} else {
			smallest = area3 / 2
		}

		wrappingPaper += (area1 + area2 + area3 + smallest)
	}

	fmt.Println(wrappingPaper)
}
