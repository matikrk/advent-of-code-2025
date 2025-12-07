package main

import (
	"fmt"

	"github.com/matikrk/advent-of-code-2025/days/day1"
	"github.com/matikrk/advent-of-code-2025/days/day2"
	"github.com/matikrk/advent-of-code-2025/days/day3"
	"github.com/matikrk/advent-of-code-2025/days/day4"
	"github.com/matikrk/advent-of-code-2025/days/day5"
	"github.com/matikrk/advent-of-code-2025/days/day6"
)

func main() {
	fmt.Println("Advent of Code 2025!")

	fmt.Println("Run dedicated solutions for each day. by running tests")
	fmt.Println("Later maybe add a CLI to run specific days")

	day1.RunEmbeded()
	day2.RunEmbeded()
	day3.RunEmbeded()
	day4.RunEmbeded()
	day5.RunEmbeded()
	day6.RunEmbeded()
}
