package dayNine

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"
)

type DayNine struct{}

func parsePoints(records []string) ([][]int, error) {
	points := [][]int{}

	for _, r := range records {
		ps := strings.Split(r, "")

		pInts := []int{}
		for _, p := range ps {
			i, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}

			pInts = append(pInts, i)
		}

		points = append(points, pInts)
	}

	return points, nil
}

func partOne(points [][]int) int {
	lowPoints := []int{}

	row := 0
	col := 0

	for row < len(points) {
		for col < len(points[0]) {
			p := points[row][col]

			if p == 0 {
				lowPoints = append(lowPoints, p)
				col++
				continue
			}

			if p == 9 {
				col++
				continue
			}

			top := row - 1
			bottom := row + 1
			left := col - 1
			right := col + 1

			if row == 0 && col == 0 {
				if points[row][right] > p && points[bottom][row] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if row == len(points)-1 && col == 0 {
				if points[row][right] > p && points[top][col] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if row == 0 && col == len(points[0])-1 {
				if points[row][right] > p && points[bottom][col] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if row == len(points)-1 && col == len(points[0])-1 {
				if points[row][left] > p && points[top][col] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if row < len(points) && col == 0 {
				if points[top][col] > p && points[bottom][col] > p && points[row][right] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if row < len(points) && col == len(points[0])-1 {
				if points[top][col] > p && points[bottom][col] > p && points[row][left] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if row == 0 && col < len(points[0])-1 {
				if points[bottom][col] > p && points[row][left] > p && points[row][right] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if row == len(points)-1 && col < len(points[0])-1 {
				if points[top][col] > p && points[row][left] > p && points[row][right] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			if top >= 0 && left >= 0 && bottom < len(points) && right < len(points[0]) {
				if points[top][col] > p && points[row][left] > p && points[row][right] > p && points[bottom][col] > p {
					lowPoints = append(lowPoints, p)
				}

				col++
				continue
			}

			col++
		}

		col = 0
		row++
	}

	fmt.Printf("lowPoints %v\n", lowPoints)

	risk := 0

	for _, lp := range lowPoints {
		risk += (lp + 1)
	}

	return risk
}

func partTwo(points [][]int) int {
	lowPoints := [][]int{}

	row := 0
	col := 0

	for row < len(points) {
		for col < len(points[0]) {
			p := points[row][col]

			if p == 0 {
				lowPoints = append(lowPoints, []int{p, row, col})
				col++
				continue
			}

			if p == 9 {
				col++
				continue
			}

			top := row - 1
			bottom := row + 1
			left := col - 1
			right := col + 1

			if row == 0 && col == 0 {
				if points[row][right] > p && points[bottom][row] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if row == len(points)-1 && col == 0 {
				if points[row][right] > p && points[top][col] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if row == 0 && col == len(points[0])-1 {
				if points[row][right] > p && points[bottom][col] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if row == len(points)-1 && col == len(points[0])-1 {
				if points[row][left] > p && points[top][col] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if row < len(points) && col == 0 {
				if points[top][col] > p && points[bottom][col] > p && points[row][right] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if row < len(points) && col == len(points[0])-1 {
				if points[top][col] > p && points[bottom][col] > p && points[row][left] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if row == 0 && col < len(points[0])-1 {
				if points[bottom][col] > p && points[row][left] > p && points[row][right] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if row == len(points)-1 && col < len(points[0])-1 {
				if points[top][col] > p && points[row][left] > p && points[row][right] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			if top >= 0 && left >= 0 && bottom < len(points) && right < len(points[0]) {
				if points[top][col] > p && points[row][left] > p && points[row][right] > p && points[bottom][col] > p {
					lowPoints = append(lowPoints, []int{p, row, col})
				}

				col++
				continue
			}

			col++
		}

		col = 0
		row++
	}

	fmt.Printf("lowPoints %v\n", lowPoints)

	risk := 0

	// for _, lp := range lowPoints {
	// 	risk += (lp + 1)
	// }

	return risk
}

func (d *DayNine) Run() error {
	tests, err := readInput.New().ReadInputTxt("/dayNine/test.txt")
	if err != nil {
		return err
	}
	inputs, err := readInput.New().ReadInputTxt("/dayNine/input.txt")
	if err != nil {
		return err
	}

	testPoints, err := parsePoints(tests)
	if err != nil {
		return err
	}

	points, err := parsePoints(inputs)
	if err != nil {
		return err
	}

	testRisk := partOne(testPoints)
	risk := partOne(points)

	fmt.Printf("PART ONE test %v\n", testRisk)
	fmt.Printf("PART ONE %v\n", risk)

	partTwo(testPoints)

	return nil
}

func New() *DayNine {
	return &DayNine{}
}
