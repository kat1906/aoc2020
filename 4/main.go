package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string // optional
}

func yearValidation(yr string, start, end int) bool {
	if len(yr) != 4 {
		return false
	}
	year, err := strconv.Atoi(yr)
	if err != nil {
		return false
	}
	if year < start || year > end {
		return false
	}
	return true
}

// define validation function for each passport field
var vf = map[string]func(string) bool{
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	"byr": func(byr string) bool {
		return yearValidation(byr, 1920, 2002)
	},
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	"iyr": func(iyr string) bool {
		return yearValidation(iyr, 2010, 2020)
	},
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	"eyr": func(eyr string) bool {
		return yearValidation(eyr, 2020, 2030)
	},
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	"hgt": func(hgt string) bool {
		if len(hgt) < 4 {
			return false
		}
		last2 := string(hgt[len(hgt)-2:])
		if last2 != "in" && last2 != "cm" {
			return false
		}
		numbs := string(hgt[:len(hgt)-2])
		numb, err := strconv.Atoi(numbs)
		if err != nil {
			return false
		}
		if last2 == "cm" {
			if numb > 193 || numb < 150 {
				return false
			}
		}
		if last2 == "in" {
			if numb > 76 || numb < 59 {
				return false
			}
		}
		return true
	},
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	"hcl": func(hcl string) bool {
		if len(hcl) != 7 {
			return false
		}
		firstChar := string(hcl[0:1])
		if firstChar != "#" {
			return false
		}
		chars := string(hcl[1:])
		isValid := regexp.MustCompile(`^[a-f0-9]+$`).MatchString
		return isValid(chars)
	},
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	"ecl": func(ecl string) bool {
		colours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, colour := range colours {
			if ecl == colour {
				return true
			}
		}
		fmt.Println(ecl)
		return false
	},
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	"pid": func(pid string) bool {
		_, err := strconv.Atoi(pid)
		if err != nil {
			return false
		}
		return len(pid) == 9
	},
}

func main() {
	passports := formatPassports(data)

	part1 := partOne(passports)
	fmt.Printf("Part 1: %v\n", part1)

	part2 := partTwo(passports)
	fmt.Printf("Part 2: %v\n", part2)
}

func partOne(passports []passport) int {
	count := 0
	for _, p := range passports {
		// check for non default values for all required passport fields
		if p.byr != "" &&
			p.iyr != "" &&
			p.eyr != "" &&
			p.hgt != "" &&
			p.hcl != "" &&
			p.ecl != "" &&
			p.pid != "" {
			count++
		}
	}
	return count
}

func partTwo(passports []passport) int {
	count := 0
	for _, p := range passports {
		// call validation functions for all passport fields
		if vf["byr"](p.byr) &&
			vf["iyr"](p.iyr) &&
			vf["eyr"](p.eyr) &&
			vf["hgt"](p.hgt) &&
			vf["hcl"](p.hcl) &&
			vf["ecl"](p.ecl) &&
			vf["pid"](p.pid) {
			count++
		}
	}
	return count
}

func formatPassports(data string) []passport {
	passports := []passport{}

	// split into data for each passport - separated by a blank line
	dataStrings := strings.Split(data, "\n\n")
	for _, pass := range dataStrings {
		formatted := passport{}

		// separate out each field - separated by whitespace
		passFields := strings.Fields(pass)
		for _, field := range passFields {
			keyValues := strings.Split(field, ":")
			value := keyValues[1]

			switch keyValues[0] {
			case "byr":
				formatted.byr = value
			case "iyr":
				formatted.iyr = value
			case "eyr":
				formatted.eyr = value
			case "hgt":
				formatted.hgt = value
			case "ecl":
				formatted.ecl = value
			case "pid":
				formatted.pid = value
			case "cid":
				formatted.cid = value
			case "hcl":
				formatted.hcl = value
			}
		}
		passports = append(passports, formatted)
	}
	return passports
}
