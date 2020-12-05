package main

import (
	"strings"
	"fmt"
)

type gradient struct {
	accross int
	down int
}

func main() {
	treeGrid := strings.Split(data, "\n")

	part1 := partOne(treeGrid);
	fmt.Printf("Part 1: %v\n", part1)

	part2 := partTwo(treeGrid);
	fmt.Printf("Part 2: %v\n", part2)
}

func countTreesHit(treeGrid []string, slope gradient) int {
	treesHitCount := 0
	// keep track of the horizontal position
	positionCount := 0
	for i := 0; i < len(treeGrid); i = i + slope.down {
		line := treeGrid[i]
		// calculate the position accounting for the repitition if positionCount*slope.accross is longer than line length
		position := (positionCount*slope.accross) % len(line)
		
		// tree is #
		if string(line[position]) == "#" {
			treesHitCount++
		}
		positionCount++
	}
	return treesHitCount
}

func partOne(treeGrid []string) int{
	return countTreesHit(treeGrid, gradient{3, 1})
}

func partTwo(treeGrid []string) int {
	r3d1 := countTreesHit(treeGrid, gradient{3, 1})
	r1d1 := countTreesHit(treeGrid, gradient{1, 1})
	r5d1 := countTreesHit(treeGrid, gradient{5, 1})
	r7d1 := countTreesHit(treeGrid, gradient{7, 1})
	r1d2 := countTreesHit(treeGrid, gradient{1, 2})

	return r1d1 * r3d1 * r5d1 *  r7d1 * r1d2
}