package main

import (
	"fmt"

	"aoc2020/helpers"
)

func main() {
	expenses, err := helpers.StringToIntSlice(data)
	if err != nil {
		fmt.Printf("error converting data to []int: %v\n", err)
	}

	part1, err := partOne(expenses)
	if err != nil {
		fmt.Printf("error in part 1: %v\n", err)
	}
	part2, err := partTwo(expenses)
	if err != nil {
		fmt.Printf("error in part 2: %v\n", err)
	}

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func partOne(expenses []int) (int, error) {
	for i, numb := range expenses {
		for j := i + 1; j < len(expenses); j++ {
			if numb + expenses[j] == 2020 {
				return numb * expenses[j], nil
			}
		}
	}
	return 0, fmt.Errorf("no two numbers adding up to 2020 found")
}

func partTwo(expenses []int) (int, error) {
	for i, numb := range expenses {
		for j := i + 1; j < len(expenses); j++ {
			for k := i + 1; k < len(expenses); k++ {
				if numb + expenses[j] + expenses[k] == 2020 {
					return numb * expenses[j] * expenses[k], nil
				}
			}
		}
	}
	return 0, fmt.Errorf("no three numbers adding up to 2020 found")
}