package main

import (
	"fmt"
	"main/dayFive"
)

func main() {
	// err := dayOne.New().Run()

	// err := dayTwo.New().Run()

	// err := dayThree.New().Run()

	// err := dayFour.New().Run()

	err := dayFive.New().Run()

	if err != nil {
		fmt.Printf(err.Error())
	}
}
