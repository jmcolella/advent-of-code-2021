package main

import (
	"fmt"
	"main/dayFour"
)

func main() {
	// err := dayOne.New().Run()

	// err := dayTwo.New().Run()

	// err := dayThree.New().Run()

	err := dayFour.New().Run()

	if err != nil {
		fmt.Printf(err.Error())
	}
}
