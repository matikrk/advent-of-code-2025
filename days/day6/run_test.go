package day6_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/matikrk/advent-of-code-2025/days/day6"
)

func Test_Run_Sample(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `

	result := day6.Run(input)
	assert.Equal(t, "4277556", result)

}
