package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(math.Pi)
	//we can declare vars like this
	var (
		name, ape string = "poop", "scoop"
		age       int    = 12
	)
	println(name, age, ape)
	test()
	add(10000, 1000)
	fmt.Println(getDay(time.Now().Weekday().String()))
}

func test() {
	name, location := "JOE", "yomama"
	age := 12
	fmt.Printf("%s is %d years old and is from %s", name, age, location)
	//vars can be be functions
	f := func() {
		//exec code
	}
	f()
	const big = 2 << 10
	fmt.Println("BIG", big)
	fmt.Println("My fav Number is", rand.Intn(10))
}

//When declaring funcs the type comes after var name in the inputs
//The return type is then specified after func name and inputs
//example
func add(x int, y int) int { //returns a int
	return x + y
}

//We can declare the type of args for both

func add2(x, y int) int {
	return x + y
}

//A example of a func that returns 2 values

func getDay(day string) (string, int) {
	var weekday string
	var timeOfDay int
	switch time.Now().Weekday() {
	case time.Tuesday:
		weekday, timeOfDay = "Tuesday", time.Now().Hour()
	default:
		weekday, timeOfDay = "Null", 1
	}
	return weekday, timeOfDay
}
