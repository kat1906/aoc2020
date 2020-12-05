package main

import (
	"fmt"
	"strconv"
	"strings"
)

type passwordWithPolicy struct {
	min      int
	max      int
	letter   string
	password string
}

func main() {
	dataStringSlice := strings.Split(data, "\n")
	passports, err := convertStringsToPasswordWithPolicy(dataStringSlice)
	if err != nil {
		fmt.Printf("error convertStringsToPasswordWithPolicy: %v", err)
	}

	part1 := partOne(passports)
	part2 := partTwo(passports)

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func convertStringsToPasswordWithPolicy(dataStrings []string) ([]passwordWithPolicy, error) {
	var passports []passwordWithPolicy
	for _, line := range dataStrings {
		passportInfo := strings.Split(line, " ")
		minMax := strings.Split(string(passportInfo[0]), "-")
		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			return passports, fmt.Errorf("error converting min: %v", err)
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			return passports, fmt.Errorf("error converting max: %v", err)
		}
		info := passwordWithPolicy{
			min:      min,
			max:      max,
			letter:   strings.TrimRight(passportInfo[1], ":"),
			password: passportInfo[2],
		}
		passports = append(passports, info)
	}
	return passports, nil
}

func partOne(passwords []passwordWithPolicy) int {
	valid := 0
	for _, p := range passwords {
		n := strings.Count(p.password, p.letter)
		if n >= p.min && n <= p.max {
			valid++
		}
	}

	return valid
}

func partTwo(passwords []passwordWithPolicy) int {
	valid := 0
	for _, p := range passwords {
		letterAtMin := string(p.password[p.min-1])
		letterAtMax := string(p.password[p.max-1])
		if letterAtMin == p.letter {
			if letterAtMax != p.letter {
				valid++
			}
		}
		if letterAtMax == p.letter {
			if letterAtMin != p.letter {
				valid++
			}
		}
	}

	return valid
}
