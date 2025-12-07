package day5

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 5: Cafeteria ---
// As the forklifts break through the wall, the Elves are delighted to discover that there was a cafeteria on the other side after all.

// You can hear a commotion coming from the kitchen. "At this rate, we won't have any time left to put the wreaths up in the dining hall!" Resolute in your quest, you investigate.

// "If only we hadn't switched to the new inventory management system right before Christmas!" another Elf exclaims. You ask what's going on.

// The Elves in the kitchen explain the situation: because of their complicated new inventory management system, they can't figure out which of their ingredients are fresh and which are spoiled. When you ask how it works, they give you a copy of their database (your puzzle input).

// The database operates on ingredient IDs. It consists of a list of fresh ingredient ID ranges, a blank line, and a list of available ingredient IDs. For example:

// 3-5
// 10-14
// 16-20
// 12-18

// 1
// 5
// 8
// 11
// 17
// 32
// The fresh ID ranges are inclusive: the range 3-5 means that ingredient IDs 3, 4, and 5 are all fresh. The ranges can also overlap; an ingredient ID is fresh if it is in any range.

// The Elves are trying to determine which of the available ingredient IDs are fresh. In this example, this is done as follows:

// Ingredient ID 1 is spoiled because it does not fall into any range.
// Ingredient ID 5 is fresh because it falls into range 3-5.
// Ingredient ID 8 is spoiled.
// Ingredient ID 11 is fresh because it falls into range 10-14.
// Ingredient ID 17 is fresh because it falls into range 16-20 as well as range 12-18.
// Ingredient ID 32 is spoiled.
// So, in this example, 3 of the available ingredient IDs are fresh.

// Process the database file from the new inventory management system. How many of the available ingredient IDs are fresh?

func Run(input string) string {
	fmt.Println("This is Day 5 solution.")
	raw := strings.Split(input, "\n\n")
	rangesRaw := strings.Split(raw[0], "\n")
	idsRaw := strings.Split(raw[1], "\n")
	// freshIDs := map[int]bool{}
	freshIDs := map[int]struct{}{} // a bit more memory efficient :D

	fmt.Println("test1")
	for _, r := range rangesRaw {
		fmt.Println("test1a", r)

		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			freshIDs[i] = struct{}{}
		}
	}
	fmt.Println("test2")

	sum := 0
	for _, idStr := range idsRaw {
		fmt.Println("idStr", idStr)
		id, _ := strconv.Atoi(idStr)
		if _, ok := freshIDs[id]; ok {
			sum = sum + 1
		}
	}

	return strconv.Itoa(sum)
}

// Optimized version with Range struct
func Run2(input string) string {
	fmt.Println("This is Day 5 solution.")
	raw := strings.Split(input, "\n\n")
	rangesRaw := strings.Split(raw[0], "\n")
	idsRaw := strings.Split(raw[1], "\n")
	// freshIDs := map[int]bool{}
	freshIDs := map[int]struct{}{} // a bit more memory efficient :D

	ranges := make([]Range, 0, len(rangesRaw))
	for _, r := range rangesRaw {
		ranges = append(ranges, newRange(r))
	}

	sum := 0
	for _, idStr := range idsRaw {
		fmt.Println("idStr", idStr)
		id, _ := strconv.Atoi(idStr)
		for _, r := range ranges {
			if r.Contains(id) {
				freshIDs[id] = struct{}{}
				sum = sum + 1
				break
			}
		}
	}

	return strconv.Itoa(sum)
}

type Range struct {
	min int
	max int
}

func newRange(line string) Range {
	parts := strings.Split(line, "-")
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])
	return Range{min: min, max: max}
}
func (r *Range) Contains(id int) bool {
	return id >= r.min && id <= r.max
}

//go:embed input.txt
var input string

func RunEmbeded() {
	result := Run2(input)
	fmt.Println("Day 5 embedded result:", result)
}
