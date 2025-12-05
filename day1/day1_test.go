package day1_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/matikrk/advent-of-code-2025/day1"
)

func TestDay1_sample(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	result := day1.Day1(input)
	assert.Equal(t, "3", result)

}
