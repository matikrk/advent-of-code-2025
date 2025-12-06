package day3_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/matikrk/advent-of-code-2025/days/day3"
)

func Test_Run_Sample(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	result := day3.Run(input)
	assert.Equal(t, "357", result)

}
