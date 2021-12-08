package dayEight

import (
	"fmt"
	"main/readInput"
	"sort"
	"strconv"
	"strings"
)

type DayEight struct{}

func splitInput(records []string) [][][]string {
	splitDelim := [][][]string{}

	for _, r := range records {
		tempSplit := strings.Split(r, "|")
		first := strings.Split(strings.Trim(tempSplit[0], " "), " ")
		second := strings.Split(strings.Trim(tempSplit[1], " "), " ")

		splitDelim = append(splitDelim, [][]string{first, second})
	}

	return splitDelim
}

func partOne(entries [][][]string) int {
	uniqLens := "2 4 3 7"

	uniqInstances := 0

	for _, e := range entries {
		for _, output := range e[1] {
			outputStr := fmt.Sprintf("%v", len(output))
			if strings.Index(uniqLens, outputStr) != -1 {
				uniqInstances += 1
			}
		}
	}

	return uniqInstances
}

func partTwo(entry [][]string) (int, error) {
	numberMap := map[int]string{0: "", 1: "", 2: "", 3: "", 4: "", 5: "", 6: "", 7: "", 8: "", 9: ""}

	foundValues := 0

	for foundValues < 10 {
		foundValues = 0

		for _, segment := range entry[0] {
			segmentLen := len(segment)
			splitSeg := strings.Split(segment, "")
			sort.Strings(splitSeg)
			sortedSeg := strings.Join(splitSeg, "")

			if segmentLen == 2 {
				numberMap[1] = sortedSeg
				continue
			}

			if segmentLen == 4 {
				numberMap[4] = sortedSeg
				continue
			}

			if segmentLen == 3 {
				numberMap[7] = sortedSeg
				continue
			}

			if segmentLen == 7 {
				numberMap[8] = sortedSeg
				continue
			}

			if segmentLen == 6 {
				if len(leftOver(numberMap[4], sortedSeg)) == 0 {
					numberMap[9] = sortedSeg
				} else {
					if len(leftOver(numberMap[7], sortedSeg)) == 0 {
						numberMap[0] = sortedSeg
					} else {
						numberMap[6] = sortedSeg
					}
				}
			}

			if segmentLen == 5 {
				if len(leftOver(numberMap[7], sortedSeg)) == 0 {
					numberMap[3] = sortedSeg
					continue
				}

				if len(leftOver(numberMap[6], sortedSeg)) == 1 && len(leftOver(numberMap[1], sortedSeg)) == 1 {
					numberMap[5] = sortedSeg
					continue
				}

				if len(leftOver(numberMap[5], sortedSeg)) == 2 && len(leftOver(numberMap[1], sortedSeg)) == 1 {
					numberMap[2] = sortedSeg
					continue
				}
			}
		}

		for _, v := range numberMap {
			if v != "" {
				foundValues += 1
			}
		}
	}

	output := []string{}

	for _, o := range entry[1] {
		for k, v := range numberMap {
			splitO := strings.Split(o, "")
			sort.Strings(splitO)
			sortedO := strings.Join(splitO, "")

			if sortedO == v {
				str := fmt.Sprintf("%v", k)
				output = append(output, str)
			}
		}
	}

	return strconv.Atoi(strings.Join(output, ""))
}

func leftOver(one string, two string) []string {
	if len(one) == 0 && len(two) == 0 {
		return []string{}
	}

	oneSplit := strings.Split(one, "")

	common := []string{}
	uniq := []string{}

	for _, i := range oneSplit {
		if strings.Index(two, i) != -1 {
			common = append(common, i)
		} else {
			uniq = append(uniq, i)
		}
	}

	return uniq
}

func (d *DayEight) Run() error {
	tests, err := readInput.New().ReadInputTxt("/dayEight/test.txt")
	if err != nil {
		return err
	}
	records, err := readInput.New().ReadInputTxt("/dayEight/input.txt")
	if err != nil {
		return err
	}

	entriesTest := splitInput(tests)
	entries := splitInput(records)

	uniqueTestInstances := partOne(entriesTest)
	uniqueInstances := partOne(entries)

	fmt.Printf("Part One TEST %v\n", uniqueTestInstances)
	fmt.Printf("Part One %v\n", uniqueInstances)

	totalTest := 0

	for _, entry := range entriesTest {
		total, err := partTwo(entry)
		if err != nil {
			return err
		}

		totalTest += total
	}

	total := 0

	for _, entry := range entries {
		t, err := partTwo(entry)
		if err != nil {
			return err
		}

		total += t
	}

	fmt.Printf("Part Two TEST %v\n", totalTest)
	fmt.Printf("Part Two %v\n", total)

	return nil
}

func New() *DayEight {
	return &DayEight{}
}
