package dayTen

import (
	"fmt"
	"main/readInput"
	"sort"
	"strings"
)

type DayTen struct{}

func partOne(records []string) int {
	pairs := map[string]int{
		"(": 1,
		")": -1,
		"[": 2,
		"]": -2,
		"{": 3,
		"}": -3,
		"<": 4,
		">": -4,
	}

	parsedRecords := [][]string{}
	for _, r := range records {
		rSlice := strings.Split(r, "")
		c := 0

		for c < len(rSlice) {
			if c == len(rSlice)-1 {
				c++
				continue
			}

			if pairs[rSlice[c]] > 0 && pairs[rSlice[c+1]] == -pairs[rSlice[c]] {
				rSlice = append(rSlice[0:c], rSlice[c+2:]...)
				c = 0
			}

			c++
		}

		parsedRecords = append(parsedRecords, rSlice)
	}

	incorrect := []string{}
	for _, pr := range parsedRecords {
		for idx, r := range pr {
			if idx == len(pr)-1 {
				continue
			}
			if pairs[pr[idx+1]] < 0 && pairs[r] > 0 && pairs[r] != -pairs[pr[idx+1]] {
				incorrect = append(incorrect, pr[idx+1])
			}
		}
	}

	scores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	total := 0
	for _, i := range incorrect {
		total += scores[i]
	}

	return total
}

func partTwo(records []string) int {
	pairs := map[string]int{
		"(": 1,
		")": -1,
		"[": 2,
		"]": -2,
		"{": 3,
		"}": -3,
		"<": 4,
		">": -4,
	}

	parsedRecords := [][]string{}
	for _, r := range records {
		fmt.Printf("R %v\n", r)

		rSlice := strings.Split(r, "")
		c := 0

		for c < len(rSlice) {
			if c == len(rSlice)-1 {
				c++
				continue
			}

			if pairs[rSlice[c]] > 0 && pairs[rSlice[c+1]] == -pairs[rSlice[c]] {
				rSlice = append(rSlice[0:c], rSlice[c+2:]...)
				c = 0
			}

			c++
		}

		parsedRecords = append(parsedRecords, rSlice)
	}

	incomplete := [][]string{}
	for _, pr := range parsedRecords {
		incorrect := false
		for idx, r := range pr {
			if idx == len(pr)-1 {
				continue
			}
			if pairs[pr[idx+1]] < 0 && pairs[r] > 0 && pairs[r] != -pairs[pr[idx+1]] {
				incorrect = true
			}

		}

		if incorrect == false {
			incomplete = append(incomplete, pr)
		}
	}

	scores := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	totals := make([]int, len(incomplete))
	for idx, line := range incomplete {
		c := len(line) - 1

		for c >= 0 {
			byFive := totals[idx] * 5
			i := line[c]
			t := byFive + scores[i]

			totals[idx] = t

			c--
		}
	}

	sort.Ints(totals)

	return totals[len(totals)/2]
}

func (d *DayTen) Run() error {
	tests, err := readInput.New().ReadInputTxt("/dayTen/tests.txt")
	if err != nil {
		return err
	}
	records, err := readInput.New().ReadInputTxt("/dayTen/input.txt")
	if err != nil {
		return err
	}

	totalTests := partOne(tests)
	total := partOne(records)

	fmt.Printf("PART ONE tests %v\n", totalTests)
	fmt.Printf("PART ONE %v\n", total)

	middleScoreTests := partTwo(tests)
	middleScore := partTwo(records)

	fmt.Printf("PART TWO tests %v\n", middleScoreTests)
	fmt.Printf("PART TWO %v\n", middleScore)

	return nil
}

func New() *DayTen {
	return &DayTen{}
}
