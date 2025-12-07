package day2

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// --- Day 2: Gift Shop ---
// You get inside and take the elevator to its only other stop: the gift shop. "Thank you for visiting the North Pole!" gleefully exclaims a nearby sign. You aren't sure who is even allowed to visit the North Pole, but you know you can access the lobby through here, and from there you can access the rest of the North Pole base.

// As you make your way through the surprisingly extensive selection, one of the clerks recognizes you and asks for your help.
// As it turns out, one of the younger Elves was playing on a gift shop computer and managed to add a whole bunch of invalid product IDs to their gift shop database! Surely, it would be no trouble for you to identify the invalid product IDs for them, right?

// They've even checked most of the product ID ranges already; they only have a few product ID ranges (your puzzle input) that you'll need to check. For example:

// 11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
// 1698522-1698528,446443-446449,38593856-38593862,565653-565659,
// 824824821-824824827,2121212118-2121212124
// (The ID ranges are wrapped here for legibility; in your input, they appear on a single long line.)

// The ranges are separated by commas (,); each range gives its first ID and last ID separated by a dash (-).

// Since the young Elf was just doing silly patterns, you can find the invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice. So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.

// None of the numbers have leading zeroes; 0101 isn't an ID at all. (101 is a valid ID that you would ignore.)

// Your job is to find all of the invalid IDs that appear in the given ranges. In the above example:

// 11-22 has two invalid IDs, 11 and 22.
// 95-115 has one invalid ID, 99.
// 998-1012 has one invalid ID, 1010.
// 1188511880-1188511890 has one invalid ID, 1188511885.
// 222220-222224 has one invalid ID, 222222.
// 1698522-1698528 contains no invalid IDs.
// 446443-446449 has one invalid ID, 446446.
// 38593856-38593862 has one invalid ID, 38593859.
// The rest of the ranges contain no invalid IDs.
// Adding up all the invalid IDs in this example produces 1227775554.
func Run(input string) string {
	fmt.Println("This is Day 2 solution.")
	rawRanges := strings.Split(input, ",")
	ranges := make([]Range, 0, len(rawRanges))
	for _, rawLine := range rawRanges {
		line, err := newRange(rawLine)
		if err != nil {
			log.Fatal("parsing range failed:", err)
		}
		ranges = append(ranges, line)
	}
	sum := 0
	for _, r := range ranges {
		ids := r.GetInvalidIds()
		fmt.Println("ids", ids)
		for _, id := range ids {
			// fmt.Println("sum", sum, id)
			sum = sum + id
		}
	}

	return strconv.Itoa(sum)
}

type Range struct {
	min int
	max int
}

func isInvalidID(id int) bool {
	// some sequence of digits repeated twice. So, 55 (5 twice), 6464 (64 twice)
	if id == 0 {
		return false
	}

	str := strconv.Itoa(id)

	if len(str)%2 != 0 {
		return false
	}

	a := str[:len(str)/2]
	b := str[len(str)/2:]

	// fmt.Println(id, a, b)
	if a == b {
		return true
	}

	return false
}
func (r *Range) GetInvalidIds() []int {
	var invalidIds []int
	for i := r.min; i <= r.max; i++ {
		if isInvalidID(i) {
			invalidIds = append(invalidIds, i)
			// fmt.Println(i)
		}
	}

	return invalidIds
}

func newRange(s string) (Range, error) {
	// 11-22
	rs := strings.Split(s, "-")
	if len(rs) != 2 {
		return Range{}, fmt.Errorf("invalid range string: %s", s)
	}
	min, err := strconv.Atoi(rs[0])
	if err != nil {
		return Range{}, err
	}
	max, err := strconv.Atoi(rs[1])
	if err != nil {
		return Range{}, err
	}
	return Range{min: min, max: max}, nil
}

//go:embed input.txt
var input string

func RunEmbeded() {
	result := Run(input)
	fmt.Println("Day 2 embedded result:", result)
}
