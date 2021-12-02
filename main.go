package main

import (
	"fmt"
	"main/dayTwo"
)

func main() {
	// err := dayOne.New().Run()

	err := dayTwo.New().Run()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
