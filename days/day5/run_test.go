package day5_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/matikrk/advent-of-code-2025/days/day5"
)

func Test_Run_Sample(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	result := day5.Run(input)
	assert.Equal(t, "3", result)

}

func Test_Run2_Sample(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	result := day5.Run2(input)
	assert.Equal(t, "3", result)

}
