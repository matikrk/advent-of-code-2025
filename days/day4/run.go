package day4

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 4: Printing Department ---
// You ride the escalator down to the printing department. They're clearly getting ready for Christmas; they have lots of large rolls of paper everywhere, and there's even a massive printer in the corner (to handle the really big print jobs).

// Decorating here will be easy: they can make their own decorations. What you really need is a way to get further into the North Pole base while the elevators are offline.

// "Actually, maybe we can help with that," one of the Elves replies when you ask for help. "We're pretty sure there's a cafeteria on the other side of the back wall. If we could break through the wall, you'd be able to keep moving. It's too bad all of our forklifts are so busy moving those big rolls of paper around."

// If you can optimize the work the forklifts are doing, maybe they would have time to spare to break through the wall.

// The rolls of paper (@) are arranged on a large grid; the Elves even have a helpful diagram (your puzzle input) indicating where everything is located.

// For example:

// ..@@.@@@@.
// @@@.@.@.@@
// @@@@@.@.@@
// @.@@@@..@.
// @@.@@@@.@@
// .@@@@@@@.@
// .@.@.@.@@@
// @.@@@.@@@@
// .@@@@@@@@.
// @.@.@@@.@.
// The forklifts can only access a roll of paper if there are fewer than four rolls of paper in the eight adjacent positions.
// If you can figure out which rolls of paper the forklifts can access, they'll spend less time looking and more time breaking down the wall to the cafeteria.

// In this example, there are 13 rolls of paper that can be accessed by a forklift (marked with x):

// ..xx.xx@x.
// x@@.@.@.@@
// @@@@@.x.@@
// @.@@@@..@.
// x@.@@@@.@x
// .@@@@@@@.@
// .@.@.@.@@@
// x.@@@.@@@@
// .@@@@@@@@.
// x.x.@@@.x.
// Consider your complete diagram of the paper roll locations. How many rolls of paper can be accessed by a forklift?

func Run(input string) string {
	fmt.Println("This is Day 4 solution.")
	raw := strings.Split(input, "\n")

	sum := 0
	var grid [][]bool
	for _, r := range raw {
		row := newRow(r)
		grid = append(grid, row)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] && isAccesible(grid, i, j) {
				sum = sum + 1
			}
		}
	}

	return strconv.Itoa(sum)
}

func isAccesible(grid [][]bool, i, j int) bool {
	maxI := len(grid) - 1
	maxJ := len(grid[0]) - 1

	isExistingPlace := func(i, j int) bool {
		return i >= 0 && i <= maxI && j >= 0 && j <= maxJ
	}
	isOccupiedPlace := func(i, j int) int {
		if isExistingPlace(i, j) && grid[i][j] {
			return 1
		}
		return 0
	}

	sumOfOccupiedFields :=
		isOccupiedPlace(i-1, j-1) +
			isOccupiedPlace(i-1, j) +
			isOccupiedPlace(i-1, j+1) +
			isOccupiedPlace(i, j-1) +
			isOccupiedPlace(i, j+1) +
			isOccupiedPlace(i+1, j-1) +
			isOccupiedPlace(i+1, j) +
			isOccupiedPlace(i+1, j+1)

	return sumOfOccupiedFields < 4
}

func newRow(line string) []bool {
	batteries := make([]bool, len(line))
	for i, ch := range line {
		if ch == '@' {
			batteries[i] = true
		}
	}
	return batteries
}

//go:embed input.txt
var input string

func RunEmbeded() {
	result := Run(input)
	fmt.Println("Day 4 embedded result:", result)
}
