package dayThree

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"
)

type DayThree struct{}

func (d *DayThree) Run() error {
	records, err := readInput.New().ReadInputCsv("/dayThree/input.csv")
	if err != nil {
		return err
	}

	input := [][]string{}

	for rIdx, record := range records {
		if rIdx == 0 {
			continue
		}

		thing := strings.Split(record[0], "")

		input = append(input, thing)
	}

	rowIdx := 0
	colIdx := 0

	gamma := ""
	epsilon := ""

	oxyBits := []string{}
	coBits := []string{}

	for colIdx < len(input[0]) {
		zeroCount := 0
		oneCount := 0

		for rowIdx < len(input) {
			val, err := strconv.ParseInt(input[rowIdx][colIdx], 10, 64)
			if err != nil {
				return err
			}

			if val == 0 {
				zeroCount += 1
			} else {
				oneCount += 1
			}

			rowIdx++
		}

		if zeroCount > oneCount {
			gamma += "0"
			epsilon += "1"
		} else if zeroCount == oneCount {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}

		colIdx++
		rowIdx = 0
	}

	for gIdx, g := range strings.Split(gamma, "") {
		if gIdx == 0 {
			for _, i := range input {
				if strings.Split(i[0], "")[gIdx] == g {
					oxyBits = append(oxyBits, strings.Join(i, ""))
				}
			}

			continue
		}

		oxyTemp := []string{}

		zeroCount := 0
		oneCount := 0

		for _, ob := range oxyBits {
			if strings.Split(ob, "")[gIdx] == "0" {
				zeroCount++
			} else {
				oneCount++
			}

		}

		for _, ob := range oxyBits {
			if zeroCount > oneCount {
				if strings.Split(ob, "")[gIdx] == "0" {
					oxyTemp = append(oxyTemp, ob)
				}
			} else if zeroCount == oneCount {
				if strings.Split(ob, "")[gIdx] == "1" {
					oxyTemp = append(oxyTemp, ob)
				}
			} else {
				if strings.Split(ob, "")[gIdx] == "1" {
					oxyTemp = append(oxyTemp, ob)
				}
			}

		}

		oxyBits = oxyTemp
	}

	for gIdx, g := range strings.Split(gamma, "") {
		if gIdx == 0 {
			zeroCount := 0
			oneCount := 0

			for _, i := range input {
				if strings.Split(i[0], "")[gIdx] == "0" {
					zeroCount++
				} else {
					oneCount++
				}
				if strings.Split(i[0], "")[gIdx] == g {
				}
			}

			for _, i := range input {
				if zeroCount < oneCount {
					if strings.Split(i[0], "")[gIdx] == "0" {
						coBits = append(coBits, strings.Join(i, ""))
					}
				} else if zeroCount == oneCount {
					if strings.Split(i[0], "")[gIdx] == "0" {
						coBits = append(coBits, strings.Join(i, ""))
					}
				} else {
					if strings.Split(i[0], "")[gIdx] == "1" {
						coBits = append(coBits, strings.Join(i, ""))
					}
				}
			}

			continue
		}

		if len(coBits) == 1 {
			continue
		}

		coTemp := []string{}

		zeroCount := 0
		oneCount := 0

		for _, co := range coBits {
			if strings.Split(co, "")[gIdx] == "0" {
				zeroCount++
			} else {
				oneCount++
			}
		}

		for _, co := range coBits {
			if zeroCount < oneCount {
				if strings.Split(co, "")[gIdx] == "0" {
					coTemp = append(coTemp, co)
				}
			} else if zeroCount == oneCount {
				if strings.Split(co, "")[gIdx] == "0" {
					coTemp = append(coTemp, co)
				}
			} else {
				if strings.Split(co, "")[gIdx] == "1" {
					coTemp = append(coTemp, co)
				}
			}
		}

		coBits = coTemp
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return err
	}

	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return err
	}

	oxyBitsDec, err := strconv.ParseInt(oxyBits[0], 2, 64)
	if err != nil {
		return err
	}

	coBitsDec, err := strconv.ParseInt(coBits[0], 2, 64)
	if err != nil {
		return err
	}

	fmt.Printf("gamma %v \n", gamma)
	fmt.Printf("epsilon %v \n", epsilon)
	fmt.Printf("gammaDec %v \n", gammaDec)
	fmt.Printf("epsilonDec %v \n", epsilonDec)
	fmt.Printf("solution %v \n", gammaDec*epsilonDec)

	fmt.Printf("oxyBits %v \n", oxyBits)
	fmt.Printf("coBits %v \n", coBits)
	fmt.Printf("oxyBitsDec %v \n", oxyBitsDec)
	fmt.Printf("coBitsDec %v \n", coBitsDec)
	fmt.Printf("solution 2 %v \n", oxyBitsDec*coBitsDec)

	return nil
}

func New() *DayThree {
	return &DayThree{}
}
