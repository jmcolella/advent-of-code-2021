package readInput

import (
	"encoding/csv"
	"os"
)

type ReadInput struct{}

func (ri *ReadInput) ReadInputCsv(filePath string) ([][]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	csvFile, err := os.Open(wd + filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(csvFile)
	return csvReader.ReadAll()
}

func New() *ReadInput {
	return &ReadInput{}
}
