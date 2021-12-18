package dayFifteen

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"
)

type DayFifteen struct{}

func parseInput(input []string) ([][]int, error) {
	parsed := [][]int{}

	for _, i := range input {
		split := strings.Split(i, "")

		temp := []int{}

		for _, s := range split {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			temp = append(temp, num)
		}

		parsed = append(parsed, temp)
	}

	return parsed, nil
}

func partOne(input [][]int) int {
	row := 0
	col := 0

	risk := 0

	for row < len(input) {
		for col < len(input[0]) {
			nextBottom := input[row+1][col]
			nextLeft := input[row][col+1]

			// if nextBottom == 1 {
			// 	fmt.Printf("nextBottom %v\n\n", nextBottom)
			// 	risk += nextBottom

			// 	row++
			// 	continue
			// } else if nextLeft == 1 {
			// 	fmt.Printf("nextLeft %v\n\n", nextLeft)

			// 	risk += nextLeft

			// 	col++
			// 	continue
			// }

			bottomTot := nextBottom + input[row+2][col]
			leftTot := nextLeft

			if col+2 < len(input[0]) {
				leftTot += input[row][col+2]
			}

			if bottomTot < leftTot {
				fmt.Printf("nextBottom %v\n\n", nextBottom)
				risk += nextBottom

				row++
			} else {
				fmt.Printf("nextLeft %v\n\n", nextLeft)

				risk += nextLeft

				col++
			}
		}
	}

	return risk
}

func (d *DayFifteen) Run() error {
	tests, err := readInput.New().ReadInputTxt("/dayFifteen/tests.txt")
	if err != nil {
		return err
	}

	parsedTest, err := parseInput(tests)
	if err != nil {
		return err
	}

	testsRisk := partOne(parsedTest)

	fmt.Printf("Part One tests %v", testsRisk)

	return nil
}

func New() *DayFifteen {
	return &DayFifteen{}
}
