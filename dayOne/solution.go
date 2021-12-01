package dayOne

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type DayOne struct{}

func (d *DayOne) Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	csvFile, err := os.Open(wd + "/dayOne/input.csv")
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	input := []int64{}

	for lineIdx, line := range records {
		if lineIdx == 0 {
			continue
		}

		lineInt, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			return err
		}

		input = append(input, lineInt)
	}

	depthMap := make(map[string]string)

	// Part One
	// for iIdx, i := range input {
	// 	depthMapKey := fmt.Sprintf("%v-%v", i, iIdx)
	// 	if iIdx == 0 {
	// 		depthMap[depthMapKey] = ""

	// 		continue
	// 	}

	// 	if i > input[iIdx-1] {
	// 		depthMap[depthMapKey] = "increased"
	// 	} else {
	// 		depthMap[depthMapKey] = "decreased"
	// 	}
	// }

	rollingDepths := []int64{}

	for iIdx, i := range input {
		if iIdx+2 <= (len(input) - 1) {
			value_two := input[iIdx+1]
			value_three := input[iIdx+2]

			sum := i + value_two + value_three
			rollingDepths = append(rollingDepths, sum)
		}
	}

	for jIdx, j := range rollingDepths {
		depthMapKey := fmt.Sprintf("%v-%v", j, jIdx)
		if jIdx == 0 {
			depthMap[depthMapKey] = ""

			continue
		}

		if j > rollingDepths[jIdx-1] {
			depthMap[depthMapKey] = "increased"
		} else {
			depthMap[depthMapKey] = "decreased"
		}
	}

	numOfIncreaseDepths := 0

	for _, depthChange := range depthMap {
		if depthChange == "increased" {
			numOfIncreaseDepths++
		}
	}

	fmt.Println(numOfIncreaseDepths)

	return nil
}

func New() *DayOne {
	return &DayOne{}
}
