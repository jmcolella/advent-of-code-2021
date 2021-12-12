package dayEleven

import (
	"main/readInput"
	"strconv"
	"strings"
)

type DayEleven struct{}

func parseInput(records []string) ([][]int, error) {
	input := [][]int{}

	for _, r := range records {
		split := strings.Split(r, "")

		temp := []int{}

		for _, s := range split {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			temp = append(temp, num)
		}

		input = append(input, temp)
	}

	return input, nil
}

func partOne(input [][]int) int {
	flashes := 0

	for r, row := range input {
		for o, oct := range row {
			nextOct := oct + 1

			if nextOct > 9 {
				flashes++
				input[r][o] = 0
				input = flash(input, r, o)
			} else {
				input[r][o] = nextOct
			}
		}
	}

	return flashes
}

func flash(input [][]int, row int, oct int) [][]int {
	flahses := 0

	countRow := row + 1
	countOct := oct + 1

	for countRow > 0 {
		for countOct > 0 {
			nextOct := input[countRow][countOct] + 1

			if nextOct > 9 {
				// maybe recursive? i don't know
				return flash(input, countRow, countOct)
			}
		}
	}
}

func (d *DayEleven) Run() error {
	tests, err := readInput.New().ReadInputTxt("/dayEleven/tests.txt")
	if err != nil {
		return err
	}
	// records, err := readInput.New().ReadInputTxt("/dayTen/input.txt")
	// if err != nil {
	// 	return err
	// }

	parsedTests, err := parseInput(tests)
	if err != nil {
		return err
	}

	partOne(parsedTests)

	// fmt.Printf("PART ONE tests %v\n", totalTests)

	return nil
}

func New() *DayEleven {
	return &DayEleven{}
}
