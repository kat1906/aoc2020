package main
import (
	"testing"
)

var input = []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#.", "..#.##.....", ".#.#.#....#", ".#........#", "#.##...#...", "#...##....#", ".#..#...#.#"}

func TestCountTreesHit(t *testing.T) {
	tests := []struct{
		name string
		in []string
		gradient gradient
		want int
	}{
		{
			name: "gradient 1, 1",
			in: input,
			gradient: gradient{1, 1},
			want: 2,
		},
		{
			name: "gradient 3, 1",
			in: input,
			gradient: gradient{3, 1},
			want: 7,
		},
		{
			name: "gradient 5, 1",
			in: input,
			gradient: gradient{5, 1},
			want: 3,
		},
		{
			name: "gradient 7, 1",
			in: input,
			gradient: gradient{7, 1},
			want: 4,
		},
		{
			name: "gradient 1, 2",
			in: input,
			gradient: gradient{1, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := countTreesHit(tt.in, tt.gradient)
			if output != tt.want {
				t.Errorf("countTreesHit(%v, %v) = %v, want %v", tt.in, tt.gradient, output, tt.want)
			}
		})
	}
}
func TestPartOne(t *testing.T) {
	test := struct{
		in []string
		want int
	}{
		in: input,
		want: 7,
	}

	output := partOne(test.in)
	if output != test.want {
		t.Errorf("partOne(%v) = %v, want %v", test.in, output, test.want)
	}
}

func TestPartTwo(t *testing.T) {
	test := struct{
		in []string
		want int
	}{
		in: input,
		want: 336,
	}
	output := partTwo(test.in)
	if output != test.want {
		t.Errorf("partTwo(%v) = %v, want %v", test.in, output, test.want)
	}
}