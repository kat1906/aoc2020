package main

import (
	"reflect"
	"testing"
)

const mockData = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestYearValidation(t *testing.T) {
	tests := []struct {
		name  string
		yr    string
		start int
		end   int

		want bool
	}{
		{
			name:  "year between start and end",
			yr:    "2020",
			start: 2010,
			end:   2030,

			want: true,
		},
		{
			name:  "year before start",
			yr:    "2000",
			start: 2010,
			end:   2030,

			want: false,
		},
		{
			name:  "year efter end",
			yr:    "2040",
			start: 2010,
			end:   2030,

			want: false,
		},
		{
			name:  "year not number string",
			yr:    "year",
			start: 2010,
			end:   2030,

			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := yearValidation(tt.yr, tt.start, tt.end)
			if output != tt.want {
				t.Errorf("countTreesHit(%v, %v, %v) = %v, want %v", tt.yr, tt.start, tt.end, output, tt.want)
			}
		})
	}
}

func TestVfByr(t *testing.T) {
	tests := []struct {
		name string
		yr   string

		want bool
	}{
		{
			name: "year between 1920 and 2002",
			yr:   "2001",

			want: true,
		},
		{
			name: "year after 2002",
			yr:   "2020",

			want: false,
		},
		{
			name: "year before 1920",
			yr:   "1901",

			want: false,
		},
		{
			name: "year not number string",
			yr:   "year",

			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := vf["byr"](tt.yr)
			if output != tt.want {
				t.Errorf(`vf["byr"](%v) = %v, want %v`, tt.yr, output, tt.want)
			}
		})
	}
}

func TestVfIyr(t *testing.T) {
	tests := []struct {
		name string
		yr   string

		want bool
	}{
		{
			name: "year between 2010 and 2020",
			yr:   "2011",

			want: true,
		},
		{
			name: "year after 2020",
			yr:   "2025",

			want: false,
		},
		{
			name: "year before 2010",
			yr:   "1901",

			want: false,
		},
		{
			name: "year not number string",
			yr:   "year",

			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := vf["iyr"](tt.yr)
			if output != tt.want {
				t.Errorf(`vf["iyr"](%v) = %v, want %v`, tt.yr, output, tt.want)
			}
		})
	}
}

func TestVfEyr(t *testing.T) {
	tests := []struct {
		name string
		yr   string

		want bool
	}{
		{
			name: "year between 2020 and 2030",
			yr:   "2021",

			want: true,
		},
		{
			name: "year after 2030",
			yr:   "2035",

			want: false,
		},
		{
			name: "year before 2020",
			yr:   "1901",

			want: false,
		},
		{
			name: "year not number string",
			yr:   "year",

			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := vf["eyr"](tt.yr)
			if output != tt.want {
				t.Errorf(`vf["eyr"](%v) = %v, want %v`, tt.yr, output, tt.want)
			}
		})
	}
}

func TestVfHgt(t *testing.T) {
	tests := []struct {
		name   string
		height string

		want bool
	}{
		{
			name:   "valid height in cm",
			height: "155cm",
			want:   true,
		},
		{
			name:   "valid height in cm",
			height: "65in",
			want:   true,
		},
		{
			name:   "height too short in cm",
			height: "55cm",
			want:   false,
		},
		{
			name:   "height too short in in",
			height: "40in",
			want:   false,
		},
		{
			name:   "height too tall in cm",
			height: "255cm",
			want:   false,
		},
		{
			name:   "height too tall in in",
			height: "80in",
			want:   false,
		},
		{
			name:   "unit not cm or in",
			height: "638m",
			want:   false,
		},
		{
			name:   "no valid number",
			height: "5e5cm",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := vf["hgt"](tt.height)
			if output != tt.want {
				t.Errorf(`vf["hgt"](%v) = %v, want %v`, tt.height, output, tt.want)
			}
		})
	}
}

func TestVfHcl(t *testing.T) {
	tests := []struct {
		name   string
		colour string

		want bool
	}{
		{
			name:   "valid colour",
			colour: "#123abc",
			want:   true,
		},
		{
			name:   "no #",
			colour: "123abc",
			want:   false,
		},
		{
			name:   "too many chars",
			colour: "#123abcd",
			want:   false,
		},
		{
			name:   "not enough chars",
			colour: "#123ab",
			want:   false,
		},
		{
			name:   "invalid char",
			colour: "#123ab.",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := vf["hcl"](tt.colour)
			if output != tt.want {
				t.Errorf(`vf["hcl"](%v) = %v, want %v`, tt.colour, output, tt.want)
			}
		})
	}
}

func TestVfEcl(t *testing.T) {
	tests := []struct {
		name   string
		colour string

		want bool
	}{
		{
			name:   "valid colour",
			colour: "amb",
			want:   true,
		},
		{
			name:   "invalid colour",
			colour: "grt",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := vf["ecl"](tt.colour)
			if output != tt.want {
				t.Errorf(`vf["ecl"](%v) = %v, want %v`, tt.colour, output, tt.want)
			}
		})
	}
}

func TestVfPid(t *testing.T) {
	tests := []struct {
		name           string
		passportNumber string

		want bool
	}{
		{
			name:           "valid number leading 0s",
			passportNumber: "000000000",
			want:           true,
		},
		{
			name:           "valid number no leading 0s",
			passportNumber: "123456789",
			want:           true,
		},
		{
			name:           "passportNumber too short",
			passportNumber: "12345678",
			want:           false,
		},
		{
			name:           "passportNumber too long",
			passportNumber: "1234567890",
			want:           false,
		},
		{
			name:           "passportNumber has non number characters",
			passportNumber: "1234s67d9",
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := vf["pid"](tt.passportNumber)
			if output != tt.want {
				t.Errorf(`vf["pid"](%v) = %v, want %v`, tt.passportNumber, output, tt.want)
			}
		})
	}
}

func TestFormatPassports(t *testing.T) {
	tests := []struct {
		name           string
		passportString string

		want []passport
	}{
		{
			name:           "formats one passport correctly",
			passportString: "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
			want: []passport{
				{
					hcl: "#ae17e1",
					iyr: "2013",
					eyr: "2024",
					ecl: "brn",
					pid: "760753108",
					byr: "1931",
					hgt: "179cm",
				},
			},
		},
		{
			name: "formats multiple passports correctly",
			passportString: "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm" +
				"\n\n" + "hcl:#123abc iyr:2014 eyr:2023 ecl:amb pid:760753100 byr:1961 hgt:159cm",
			want: []passport{
				{hcl: "#ae17e1", iyr: "2013", eyr: "2024", ecl: "brn", pid: "760753108", byr: "1931", hgt: "179cm"},
				{hcl: "#123abc", iyr: "2014", eyr: "2023", ecl: "amb", pid: "760753100", byr: "1961", hgt: "159cm"},
			},
		},
		{
			name:           "formats passport if fields are missing",
			passportString: "\n\n",
			want:           []passport{{}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := formatPassports(tt.passportString)
			if !reflect.DeepEqual(output, tt.want) {
				t.Errorf(`vf["pid"](%v) = %v, want %v`, tt.passportString, output, tt.want)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		name      string
		passports []passport

		want int
	}{
		{
			name:      "one valid passport",
			passports: formatPassports(`hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm`),
			want:      1,
		},
		{
			name: "multiple valid passports",
			passports: formatPassports("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm" +
				"\n\n" + "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:169cm"),
			want: 2,
		},
		{
			name: "multiple passports with one invalid",
			passports: formatPassports("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm" +
				"\n\n" + "iyr:2013 eyr:2024 ecl:rn pid:760753108 byr:1931 hgt:179cm" +
				"\n\n" + "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:169cm"),
			want: 2,
		},
		{
			name:      "multiple passports with multiple invalid",
			passports: formatPassports(mockData),
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := partOne(tt.passports)
			if output != tt.want {
				t.Errorf(`partOne(%v) = %v, want %v`, tt.passports, output, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name      string
		passports []passport

		want int
	}{
		{
			name:      "one valid passport",
			passports: formatPassports(`hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm`),
			want:      1,
		},
		{
			name: "multiple valid passports",
			passports: formatPassports("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm" +
				"\n\n" + "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:169cm"),
			want: 2,
		},
		{
			name: "multiple passports with one invalid",
			passports: formatPassports("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm" +
				"\n\n" + "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:rn pid:760753108 byr:1931 hgt:179cm" +
				"\n\n" + "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:169cm"),
			want: 2,
		},
		{
			name:      "multiple passports with multiple invalid",
			passports: formatPassports(mockData),
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := partTwo(tt.passports)
			if output != tt.want {
				t.Errorf(`partTwo(%v) = %v, want %v`, tt.passports, output, tt.want)
			}
		})
	}
}
