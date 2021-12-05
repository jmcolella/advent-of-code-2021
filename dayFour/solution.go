package dayFour

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"
)

type DayFour struct{}

func parseInput() ([]string, [][][]string, error) {
	records, err := readInput.New().ReadInputCsv("/dayFour/input.csv")
	if err != nil {
		return nil, nil, err
	}

	drawings := []string{}
	boards := [][][]string{}

	board := [][]string{}

	for rIdx, r := range records {
		if rIdx == 0 {
			continue
		}

		if rIdx == 1 {
			for _, d := range strings.Split(r[0], ",") {
				drawings = append(drawings, d)
			}

			continue
		}

		sanitizeR := []string{}
		for _, i := range r {
			for _, j := range strings.Split(i, " ") {
				if len(j) == 0 {
					continue
				}

				sanitizeR = append(sanitizeR, j)
			}
		}

		board = append(board, sanitizeR)

		if len(board) == 5 {
			boards = append(boards, board)

			board = [][]string{}
		}
	}

	return drawings, boards, nil
}

func draw(drawings []string) ([]string, []string) {
	nextDrawings := []string{}
	remainingDrawings := []string{}

	i := 0

	for i < len(drawings) {
		if i < 5 {
			nextDrawings = append(nextDrawings, drawings[i])
		} else {
			remainingDrawings = append(remainingDrawings, drawings[i])
		}
		i++
	}

	return nextDrawings, remainingDrawings
}

func markBoards(boards [][][]string, drawing string) [][][]string {
	markedBoards := [][][]string{}

	for _, b := range boards {
		markedBoard := [][]string{}

		row := 0
		col := 0

		for row < 5 {
			markedRow := []string{}

			for col < 5 {
				if b[row][col] == drawing {
					markedRow = append(markedRow, "X")
				} else {
					markedRow = append(markedRow, b[row][col])
				}

				col++
			}

			markedBoard = append(markedBoard, markedRow)
			col = 0

			row++
		}

		markedBoards = append(markedBoards, markedBoard)

		row = 0
	}

	return markedBoards
}

func sumBoard(markedBoard [][]string) (int64, error) {
	var sum int64

	for _, row := range markedBoard {
		for _, item := range row {
			if item == "" {
				continue
			}

			if item != "X" {
				num, err := strconv.ParseInt(item, 10, 64)

				if err != nil {
					return 0, err
				}

				sum += num
			}
		}
	}

	return sum, nil
}

func (d *DayFour) Run() error {
	drawings, boards, err := parseInput()
	if err != nil {
		return err
	}

	fmt.Printf("drawings %v \n", drawings)
	fmt.Printf("boards %v \n", boards)

	nextBoards := boards

	won := false
	winnerIdx := 0
	winnerDraw := "0"

	for len(drawings) > 0 && won == false {
		nextDrawings, remainingDrawings := draw(drawings)

		d := 0

		for d < len(nextDrawings) && won == false {
			fmt.Printf("NEXT DRAWING %v\n", nextDrawings[d])

			markedBoards := markBoards(nextBoards, nextDrawings[d])
			yetToWinBoards := [][][]string{}

			for bIdx, b := range markedBoards {
				row := 0
				col := 0
				boardWon := false

				for row < 5 {
					rowWin := []string{}
					colWin := []string{}

					for col < 5 {
						if b[row][col] == "X" {
							rowWin = append(rowWin, "X")
						}

						if b[col][row] == "X" {
							colWin = append(colWin, "X")
						}

						if strings.Join(rowWin, "") == "XXXXX" || strings.Join(colWin, "") == "XXXXX" {
							winnerDraw = nextDrawings[d]
							winnerIdx = bIdx
							boardWon = true

							if len(markedBoards) == 1 {
								won = true
							}
						}

						col++
					}

					col = 0

					row++
				}

				if boardWon == false || won == true {
					yetToWinBoards = append(yetToWinBoards, b)
				}

				row = 0
			}

			nextBoards = yetToWinBoards
			d++
		}

		drawings = remainingDrawings
	}

	fmt.Printf("WINNER IDX %v\n", winnerIdx)
	fmt.Printf("WINNER DRAW %v\n", winnerDraw)

	winSum, err := sumBoard(nextBoards[winnerIdx])
	if err != nil {
		return err
	}

	fmt.Printf("WINNER SUM %v\n", winSum)

	winnerDrawNum, err := strconv.ParseInt(winnerDraw, 10, 64)
	if err != nil {
		return err
	}

	fmt.Printf("SOLUTION %v\n", winSum*winnerDrawNum)

	return nil
}

func New() *DayFour {
	return &DayFour{}
}
