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
	copySchool := school
	maxDays := 80
	dayCount := 0

	for dayCount < maxDays {
		fmt.Printf("Day Count %v\n", dayCount)
		for idx, fish := range copySchool {
			nextFish := fish - 1

			if nextFish < 0 {
				copySchool[idx] = 6

				copySchool = append(copySchool, 8)

				continue
			}

			copySchool[idx] = fish - 1
		}

		dayCount++
	}

	return copySchool
}

func partTwo(school []int64) int {
	maxDays := 256
	dayCount := 0
	var fishCounts = map[int64]int{}

	var timer int64
	timer = 0

	for timer < 9 {
		fishCounts[timer] = 0

		timer++
	}

	for _, s := range school {
		fishCounts[s] += 1
	}

	for dayCount < maxDays {
		nextFish := fishCounts[0]

		fishCounts[0] = fishCounts[1]
		fishCounts[1] = fishCounts[2]
		fishCounts[2] = fishCounts[3]
		fishCounts[3] = fishCounts[4]
		fishCounts[4] = fishCounts[5]
		fishCounts[5] = fishCounts[6]
		fishCounts[6] = fishCounts[7] + nextFish
		fishCounts[7] = fishCounts[8]
		fishCounts[8] = nextFish

		dayCount++
	}

	val := 0

	for _, v := range fishCounts {
		val += v
	}

	return val
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

	infiniteSchoolVal := partTwo(school)

	fmt.Printf("Final school Part Two %v\n", infiniteSchoolVal)

	return nil
}

func New() *DaySix {
	return &DaySix{}
}
