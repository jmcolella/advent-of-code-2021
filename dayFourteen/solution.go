package dayFourteen

import (
	"fmt"
	"main/readInput"
	"strings"
)

type DayFourteen struct{}

func parse(input []string) (string, map[string]string) {
	polymer := ""
	pairs := map[string]string{}

	for idx, line := range input {
		if idx == 0 {
			polymer = line

			continue
		}

		if strings.Index(line, "->") != -1 {
			p := strings.Split(line, " -> ")

			pairs[p[0]] = p[1]
		}
	}

	return polymer, pairs
}

func partOne(polymer string, pairs map[string]string) int {
	nextPolymer := strings.Split(polymer, "")

	c := 0

	for c < 10 {
		temp := []string{}

		for idx, p := range nextPolymer {
			if idx == len(nextPolymer)-1 {
				temp = append(temp, p)

				continue
			}
			pair := fmt.Sprintf("%v%v", p, nextPolymer[idx+1])

			temp = append(temp, []string{p, pairs[pair]}...)
		}

		fmt.Printf("next polyer %v\n\n\n", nextPolymer)

		nextPolymer = temp

		c++
	}

	occurs := map[string]int{}

	for _, p := range nextPolymer {
		if _, ok := occurs[p]; ok {
			occurs[p] += 1
		} else {
			occurs[p] = 1
		}
	}

	min := -1
	max := -1

	for _, v := range occurs {
		if v > max {
			max = v
		}

		if min < 0 {
			min = v
		} else if v < min {
			min = v
		}
	}

	return max - min
}

func (d *DayFourteen) Run() error {
	tests, err := readInput.New().ReadInputTxt("/dayFourteen/tests.txt")
	if err != nil {
		return err
	}
	input, err := readInput.New().ReadInputTxt("/dayFourteen/input.txt")
	if err != nil {
		return err
	}

	testPolymer, testPairs := parse(tests)
	polymer, pairs := parse(input)

	testsDiff := partOne(testPolymer, testPairs)
	diff := partOne(polymer, pairs)

	fmt.Printf("Part One tests %v\n", testsDiff)
	fmt.Printf("Part One %v\n", diff)

	return nil
}

func New() *DayFourteen {
	return &DayFourteen{}
}
