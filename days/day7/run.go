package day7

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// --- Day 7: Laboratories ---
// You thank the cephalopods for the help and exit the trash compactor, finding yourself in the familiar halls of a North Pole research wing.

// Based on the large sign that says "teleporter hub", they seem to be researching teleportation; you can't help but try it for yourself and step onto the large yellow teleporter pad.

// Suddenly, you find yourself in an unfamiliar room! The room has no doors; the only way out is the teleporter. Unfortunately, the teleporter seems to be leaking magic smoke.

// Since this is a teleporter lab, there are lots of spare parts, manuals, and diagnostic equipment lying around. After connecting one of the diagnostic tools, it helpfully displays error code 0H-N0, which apparently means that there's an issue with one of the tachyon manifolds.

// You quickly locate a diagram of the tachyon manifold (your puzzle input).
// A tachyon beam enters the manifold at the location marked S; tachyon beams always move downward.
// Tachyon beams pass freely through empty space (.).
// However, if a tachyon beam encounters a splitter (^), the beam is stopped; instead, a new tachyon beam continues from the immediate left and from the immediate right of the splitter.

type Operation rune

const (
	Start    Operation = 'S'
	Beam     Operation = '|'
	Splitter Operation = '^'
	Empty    Operation = '.'
)

func Run(input string) string {
	fmt.Println("This is day7 solution placeholder.")
	raw := strings.Split(input, "\n")

	usedSpliters := 0
	results := make([]Operation, len(raw[0]))
	for _, r := range raw {
		for i, e := range r {
			switch Operation(e) {
			case Start:
				results[i] = Beam
			case Empty:
				// results[i] = results[i]
			case Splitter:
				if results[i] == Beam {
					usedSpliters++
					results[i] = Empty
					results[i-1] = Beam
					results[i+1] = Beam // to powoduje blad za proste podejscie - powinienm zrobic nowa zmienna
				}
			case Beam: // not needed
				results[i] = Beam
			}
		}
	}

	return strconv.Itoa(usedSpliters)
}

//go:embed input.txt
var input string

func RunEmbeded() {
	result := Run(input)
	fmt.Println("day7 embedded result:", result)
}
