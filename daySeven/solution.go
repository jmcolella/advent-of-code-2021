package daySeven

import (
	"fmt"
	"main/readInput"
	"math"
	"strconv"
	"strings"
)

type DaySeven struct{}

func makePositions(records []string) ([]int, error) {
	positions := []int{}

	splitRecords := strings.Split(records[0], ",")

	for _, r := range splitRecords {
		val, err := strconv.Atoi(r)
		if err != nil {
			return nil, err
		}

		positions = append(positions, val)
	}

	return positions, nil
}

func partOne(positions []int) int {
	posFreqs := make(map[int]int)

	for _, p := range positions {
		if posFreqs[p] != 0 {
			posFreqs[p] += 1
		} else {
			posFreqs[p] = 1
		}
	}

	fuel := 0

	for k := range posFreqs {
		checkFuel := 0
		for _, p := range positions {
			f := float64(p - k)

			checkFuel += int(math.Abs(f))
		}

		if fuel == 0 {
			fuel = checkFuel
		} else if fuel > checkFuel {
			fuel = checkFuel
		}
	}

	return fuel
}

func partOneMode(positions []int) int {
	posFreqs := make(map[int]int)

	for _, p := range positions {
		if posFreqs[p] != 0 {
			posFreqs[p] += 1
		} else {
			posFreqs[p] = 1
		}
	}

	mode := 0
	mostFreq := 0

	for k, v := range posFreqs {
		if v > mostFreq {
			mostFreq = v
			mode = k
		}
	}

	fuel := 0

	for _, p := range positions {
		f := float64(p - mode)

		fuel += int(math.Abs(f))
	}

	return fuel
}

func partTwo(positions []int) int {
	posFreqs := make(map[int]int)

	for _, p := range positions {
		if posFreqs[p] != 0 {
			posFreqs[p] += 1
		} else {
			posFreqs[p] = 1
		}
	}

	maxPos := 0

	for _, p := range positions {
		if p > maxPos {
			maxPos = p
		}
	}

	fuel := 0

	for i := 0; i < maxPos; i++ {
		checkFuel := 0
		for k, v := range posFreqs {
			fStep := 0
			f := int(math.Abs(float64(i - k)))

			for j := 1; j <= f; j++ {
				fStep += j
			}

			checkFuel += (fStep * v)
		}

		if fuel == 0 {
			fuel = checkFuel
		} else {
			fuel = int(math.Min(float64(fuel), float64(checkFuel)))
		}
	}

	return fuel
}

func partTwoMean(positions []int) int {
	total := 0

	for _, p := range positions {
		total += p
	}

	mean := math.Ceil(float64(total) / float64(len(positions)))

	fuel := 0

	for _, p := range positions {
		d := int(math.Abs(float64(float64(p) - mean)))

		fuel += ((d * (d + 1)) / 2)
	}

	return fuel
}

func (d *DaySeven) Run() error {
	tests, err := readInput.New().ReadInputTxt("/daySeven/test.txt")
	if err != nil {
		return err
	}
	records, err := readInput.New().ReadInputTxt("/daySeven/input.txt")
	if err != nil {
		return err
	}

	testPositions, err := makePositions(tests)
	if err != nil {
		return err
	}
	positions, err := makePositions(records)
	if err != nil {
		return err
	}

	lowestFuelTest := partOne(testPositions)
	lowestFuel := partOne(positions)
	fmt.Printf("Part One TEST %v\n", lowestFuelTest)
	fmt.Printf("Part One INPUT %v\n", lowestFuel)

	lowestFuelTest = partOneMode(testPositions)
	lowestFuel = partOneMode(positions)
	fmt.Printf("Part One TEST Mode %v\n", lowestFuelTest)
	fmt.Printf("Part One INPUT Mode %v\n", lowestFuel)

	lowestFuelTest = partTwo(testPositions)
	lowestFuel = partTwo(positions)
	fmt.Printf("Part Two TEST %v\n", lowestFuelTest)
	fmt.Printf("Part Two INPUT %v\n", lowestFuel)

	lowestFuelTest = partTwoMean(testPositions)
	lowestFuel = partTwoMean(positions)
	fmt.Printf("Part Two TEST Mean %v\n", lowestFuelTest)
	fmt.Printf("Part Two INPUT Mean %v\n", lowestFuel)

	return nil
}

func New() *DaySeven {
	return &DaySeven{}
}
