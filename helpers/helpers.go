// Package helpers contains functions useful for manipulating data in aoc challenges. 
package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToIntSlice converts a string containing numbers separated by \n into a slice of int and any error encountered.
func StringToIntSlice(input string) ([]int, error) {
	strSlice := strings.Split(input, "\n")
	var numbers []int
	for _, x := range strSlice {
		i, err := strconv.Atoi(x)
		if err != nil {
			return nil, fmt.Errorf("could not convert to int: %v", x)
		}
		numbers = append(numbers, i)
	}
	return numbers, nil
}
