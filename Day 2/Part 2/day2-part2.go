//elves measuring ribbon
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
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

	var ribbon int

	for i := 0; i < len(nums); i += 3 {
		l, w, h, shortest1, shortest2, volume := 0, 0, 0, 0, 0, 0

		l = nums[i]
		w = nums[i+1]
		h = nums[i+2]

		if l <= w || l <= h {
			shortest1 = l
			if w < h {
				shortest2 = w
			} else {
				shortest2 = h
			}
		} else {
			shortest1 = w
			shortest2 = h
		}

		volume = l * w * h

		ribbon += ((shortest1 * 2) + (shortest2 * 2) + volume)
	}

	fmt.Println(ribbon)
}
