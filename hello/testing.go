package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	println(sum)
	switchWithoutCondition()
}

func switchWithoutCondition() {
	fmt.Println("Executing function")
	defer fmt.Println("Function executed")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("gg")
	default:
		fmt.Println("IDK")
	}
}
