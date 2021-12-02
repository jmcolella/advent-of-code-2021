package dayTwo

import (
	"fmt"
	"main/readInput"
	"regexp"
	"strconv"
)

type DayTwo struct{}

func (d *DayTwo) Run() error {
	records, err := readInput.New().ReadInputCsv("/dayTwo/input.csv")
	if err != nil {
		return err
	}

	input := []string{}

	for lineIdx, line := range records {
		if lineIdx == 0 {
			continue
		}

		input = append(input, line[0])
	}

	movement := []int64{0, 0, 0}

	re := regexp.MustCompile(`(\D+) (\d+)`)

	for _, i := range input {
		submatch := re.FindStringSubmatch(i)
		direction := submatch[1]
		value, err := strconv.ParseInt(submatch[2], 10, 64)
		if err != nil {
			return err
		}

		if direction == "forward" {
			movement[0] += value
			movement[1] += (movement[2] * value)
		}

		if direction == "down" {
			movement[2] += value
		}

		if direction == "up" {
			movement[2] -= value
		}
	}

	fmt.Printf("%v", movement[0]*movement[1])

	return nil
}

func New() *DayTwo {
	return &DayTwo{}
}
