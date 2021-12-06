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

func partTwo(school []int64) int {
	maxDays := 2
	dayCount := 0
	var thing = map[int64]int{}

	var timer int64
	timer = 0

	for timer < 9 {
		thing[timer] = 0

		timer++
	}

	for _, s := range school {
		thing[s] += 1
	}

	for dayCount < maxDays {
		fmt.Printf("day count %v\n", dayCount)
		fmt.Printf("thing map %v\n", thing)

		// var changeThing = map[int64]int{}

		for k := range thing {
			modEight := dayCount % 8
			// changeThing[k] = 0
			times := thing[k]

			if modEight == 0 && k == 8 && times > 0 {
				thing[6] += times
				thing[8] -= times
			} else if k == 0 && times > 0 {
				thing[6] += times
				thing[8] += times
				thing[0] -= times
			} else if times != 0 {
				fmt.Printf("k %v\n", k)

				thing[k] -= times
				thing[k-1] += times

				if thing[k] < 0 {
					thing[k] = 0
				}

				// fmt.Printf("change thing map %v\n", changeThing)

			}
		}

		// fmt.Printf("change thing map after change %v\n", changeThing)

		// for k, v := range changeThing {
		// 	thing[k]["Times"] = v
		// }

		fmt.Printf("thing map after change %v\n", thing)

		dayCount++

	}

	fmt.Printf("thing map after %v\n", thing)

	// val := 0

	// for _, v := range thing {
	// 	val += v["Times"] * v["NewFish"]
	// }

	return 0
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

	// finalSchool := partOne(school)

	// fmt.Printf("Final school Part One %v\n", len(finalSchool))

	finalSchool := partTwo(school)

	fmt.Printf("Final school Part Two %v\n", finalSchool)

	return nil
}

func New() *DaySix {
	return &DaySix{}
}
