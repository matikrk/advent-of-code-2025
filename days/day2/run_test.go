package day2_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/matikrk/advent-of-code-2025/days/day2"
)

func Test_Run_Sample(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	result := day2.Run(input)
	assert.Equal(t, "1227775554", result)

}
