package daySix

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"
)

type DaySix struct{}

func initialSchool(records []string) ([]int64, error) {
	school := []int64{}

	for _, r := range strings.Split(records[0], ",") {
		num, err := strconv.ParseInt(r, 10, 64)
		if err != nil {
			return nil, err
		}

		school = append(school, num)
	}

	return school, nil
}

func partOne(school []int64) []int64 {
	maxDays := 80
	dayCount := 0

	for dayCount < maxDays {
		fmt.Printf("Day Count %v\n", dayCount)
		for idx, fish := range school {
			nextFish := fish - 1

			if nextFish < 0 {
				school[idx] = 6

				school = append(school, 8)

				continue
			}

			school[idx] = fish - 1
		}

		dayCount++
	}

	return school
}

func partTwo(school []int64) []int64 {
	maxDays := 80
	dayCount := 0
	newFish := []int64{}

	for dayCount < maxDays {
		fmt.Printf("Day Count %v\n", dayCount)
		for idx, fish := range school {
			nextFish := fish - 1

			if nextFish < 0 {
				school[idx] = 6

				newFish = append(newFish, 8)

				continue
			}

			school[idx] = fish - 1
		}

		dayCount++
	}

	dayCount = 0

	for dayCount < (maxDays / 8) {
		nextFish := make([]int64, len(newFish))
		newFish = append(newFish, nextFish...)

		dayCount++
	}

	fmt.Printf("new fish %v\n", len(newFish))

	school = append(school, newFish...)

	return school
}

func (d *DaySix) Run() error {
	records, err := readInput.New().ReadInputTxt("/daySix/input.txt")
	if err != nil {
		return err
	}

	school, err := initialSchool(records)
	if err != nil {
		return err
	}

	finalSchool := partOne(school)

	fmt.Printf("Final school Part One %v\n", len(finalSchool))

	finalSchool = partTwo(school)

	fmt.Printf("Final school Part Two %v\n", len(finalSchool))

	return nil
}

func New() *DaySix {
	return &DaySix{}
}
