package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pinkahd/aoc-2021/utils"
)

func firstPart(s []int) int {
	count := 0
	for i, element := range s {
		if i != 0 {
			prev := s[i-1]
			current := element
			if current > prev {
				count++
			}
		}
	}
	return count
}

func secondPart(s []int) int {
	count := 0
	prevSum := 0
	for i := 0; i < len(s)-2; i++ {
		currentSum := utils.Reduce(0, utils.Itoifa(s[i:i+3]), utils.Sum).(int)
		if i != 0 && currentSum > prevSum {
			count++
		}
		prevSum = currentSum
	}
	return count
}

func main() {
	if len(os.Args) != 2 {
		utils.Bail("Require one file as argument.")
	}
	file := utils.ReadFile(os.Args[1])
	iarr := utils.MakeIntArr(strings.Split(file, "\n"))
	fmt.Printf("Hello from day1.\n%v", file)
	fmt.Printf("Part 1 => %v\n", firstPart(iarr))
	fmt.Printf("Part 2 => %v\n", secondPart(iarr))
}
