package main

import (
	"reflect"
	"testing"
)

var (
	validPassportPart1 = passwordWithPolicy{
		min:      1,
		max:      13,
		letter:   "r",
		password: "gqdrspndrpsrjfjx",
	}
	validPassportPart2 = passwordWithPolicy{
		min:      1,
		max:      3,
		letter:   "a",
		password: "abcde",
	}
)

func TestConvertStringsToPasswordWithPolicy(t *testing.T) {
	var tests = []struct {
		name    string
		in      []string
		want    []passwordWithPolicy
		wantErr string
	}{
		{
			name: "one string to passportWithPolicy",
			in:   []string{"1-13 r: gqdrspndrpsrjfjx"},
			want: []passwordWithPolicy{
				{
					min:      1,
					max:      13,
					letter:   "r",
					password: "gqdrspndrpsrjfjx",
				},
			},
		},
		{
			name: "multiple strings to passportWithPolicy",
			in:   []string{"1-13 r: gqdrspndrpsrjfjx", "5-15 x: xxjxwshpxjxxxxsnxvz"},
			want: []passwordWithPolicy{
				{
					min:      1,
					max:      13,
					letter:   "r",
					password: "gqdrspndrpsrjfjx",
				},
				{
					min:      5,
					max:      15,
					letter:   "x",
					password: "xxjxwshpxjxxxxsnxvz",
				},
			},
		},
		{
			name:    "invalid min",
			in:      []string{"a-13 r: gqdrspndrpsrjfjx"},
			wantErr: `error converting min: strconv.Atoi: parsing "a": invalid syntax`,
		},
		{
			name:    "invalid max",
			in:      []string{"1-1a3 r: gqdrspndrpsrjfjx"},
			wantErr: `error converting max: strconv.Atoi: parsing "1a3": invalid syntax`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := convertStringsToPasswordWithPolicy(tt.in)
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("StringToIntSlice(%v) = error: %v, want %v", tt.in, err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(output, tt.want) {
					t.Errorf("convertStringsToPasswordWithPolicy(%v) = %v, want %v", tt.in, output, tt.want)
				}
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	var tests = []struct {
		name    string
		in      []passwordWithPolicy
		want    int
		wantErr string
	}{
		{
			name: "passport matches policy",
			in: []passwordWithPolicy{
				passwordWithPolicy{
					min:      5,
					max:      10,
					letter:   "a",
					password: "aahnaahgfda",
				},
			},
			want: 1,
		},
		{
			name: "passport has less than min number of the policy letter",
			in: []passwordWithPolicy{
				passwordWithPolicy{
					min:      5,
					max:      10,
					letter:   "a",
					password: "aa",
				},
			},
			want: 0,
		},
		{
			name: "passport has more than max number of the policy letter",
			in: []passwordWithPolicy{
				passwordWithPolicy{
					min:      5,
					max:      10,
					letter:   "a",
					password: "aaaaaaaaaaa",
				},
			},
			want: 0,
		},
		{
			name: "multiple valid passports",
			in: []passwordWithPolicy{
				validPassportPart1,
				validPassportPart1,
				validPassportPart1,
			},
			want: 3,
		},
		{
			name: "valid and invalid passports",
			in: []passwordWithPolicy{
				validPassportPart1,
				passwordWithPolicy{},
				validPassportPart1,
				passwordWithPolicy{},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, _ := partOne(tt.in)
			if output != tt.want {
				t.Errorf("partOne(%v) = %v, want %v", tt.in, output, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	var tests = []struct {
		name    string
		in      []passwordWithPolicy
		want    int
		wantErr string
	}{
		{
			name: "password matches policy at min but not at max",
			in: []passwordWithPolicy{
				passwordWithPolicy{
					min:      1,
					max:      2,
					letter:   "a",
					password: "aab",
				},
			},
			want: 1,
		},
		{
			name: "password matches policy at max but not at min",
			in: []passwordWithPolicy{
				passwordWithPolicy{
					min:      1,
					max:      2,
					letter:   "a",
					password: "aba",
				},
			},
			want: 1,
		},
		{
			name: "letters at min and max match policy letter",
			in: []passwordWithPolicy{
				passwordWithPolicy{
					min:      1,
					max:      2,
					letter:   "a",
					password: "aaa",
				},
			},
			want: 0,
		},
		{
			name: "multiple valid passports",
			in: []passwordWithPolicy{
				validPassportPart2,
				validPassportPart2,
				validPassportPart2,
			},
			want: 3,
		},
		{
			name: "valid and invalid passports",
			in: []passwordWithPolicy{
				validPassportPart2,
				passwordWithPolicy{},
				validPassportPart2,
				passwordWithPolicy{},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, _ := partOne(tt.in)
			if output != tt.want {
				t.Errorf("partTwo(%v) = %v, want %v", tt.in, output, tt.want)
			}
		})
	}
}
