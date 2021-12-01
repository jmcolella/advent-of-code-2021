package main

import (
	"fmt"
	"main/dayOne"
)

func main() {
	err := dayOne.New().Run()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
