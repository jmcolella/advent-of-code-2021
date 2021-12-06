package dayFive

import (
	"fmt"
	"main/readInput"
	"regexp"
	"strconv"
)

type DayFive struct{}

func makeLineCoordinates(records []string) ([][][]int64, int64, error) {
	coords := [][][]int64{}
	var max int64
	max = 0

	re, err := regexp.Compile(`(\d+),(\d+) -> (\d+),(\d+)`)
	if err != nil {
		return nil, max, err
	}

	for _, r := range records {
		submatch := re.FindStringSubmatch(r)
		x1, err := strconv.ParseInt(submatch[1], 10, 64)
		y1, err := strconv.ParseInt(submatch[2], 10, 64)
		x2, err := strconv.ParseInt(submatch[3], 10, 64)
		y2, err := strconv.ParseInt(submatch[4], 10, 64)
		if err != nil {
			return nil, max, err
		}

		if x1 > max {
			max = x1
		} else if x2 > max {
			max = x2
		} else if y1 > max {
			max = y1
		} else if y2 > max {
			max = y2
		}

		row := [][]int64{}

		xCount := x1
		xEnd := x2

		yCount := y1
		yEnd := y2

		for xCount != xEnd || yCount != yEnd {
			row = append(row, []int64{xCount, yCount})

			if y2 > y1 {
				yCount++
			} else if y1 > y2 {
				yCount--
			}

			if x2 > x1 {
				xCount++
			} else if x1 > x2 {
				xCount--
			}
		}

		row = append(row, []int64{xCount, yCount})

		coords = append(coords, row)
	}

	maxLen := max + 1

	return coords, maxLen, nil
}

func makeGrid(max int64) [][]int64 {
	var row int64
	var col int64
	row = 0
	col = 0

	grid := [][]int64{}

	for row < max {
		r := []int64{}
		for col < max {
			r = append(r, 0)

			col++
		}

		grid = append(grid, r)
		col = 0
		row++
	}

	return grid
}

func mark(grid [][]int64, coord [][]int64) [][]int64 {
	for _, c := range coord {
		x := c[0]
		y := c[1]

		grid[y][x] += 1
	}

	return grid
}

func (d *DayFive) Run() error {
	records, err := readInput.New().ReadInputTxt("/dayFive/input.txt")
	if err != nil {
		return err
	}

	coords, max, err := makeLineCoordinates(records)
	if err != nil {
		return err
	}

	grid := makeGrid(max)

	markedGrid := grid

	for _, coord := range coords {
		markedGrid = mark(markedGrid, coord)
	}

	dangerPoints := 0

	for _, m := range markedGrid {
		for _, i := range m {
			if i > 1 {
				dangerPoints++
			}
		}
	}

	fmt.Printf("danger points %v\n", dangerPoints)

	return nil
}

func New() *DayFive {
	return &DayFive{}
}
