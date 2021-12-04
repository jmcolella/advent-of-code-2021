package main

import (
	"fmt"
	"main/dayThree"
)

func main() {
	// err := dayOne.New().Run()

	// err := dayTwo.New().Run()

	err := dayThree.New().Run()

	if err != nil {
		fmt.Printf(err.Error())
	}
}
