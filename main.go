package main

import (
	"fmt"
	"main/dayTen"
)

func main() {
	// err := dayOne.New().Run()
	// err := dayTwo.New().Run()
	// err := dayThree.New().Run()
	// err := dayFour.New().Run()
	// err := dayFive.New().Run()
	// err := daySix.New().Run()
	// err := daySeven.New().Run()
	// err := dayEight.New().Run()
	// err := dayNine.New().Run()

	err := dayTen.New().Run()

	if err != nil {
		fmt.Printf(err.Error())
	}
}
