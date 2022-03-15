package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
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
	//fmt.Println(getDay(time.Now().Weekday().String()))
	fmt.Println(getDay2(time.Now().Weekday().String()))
	client := &http.Client{}
	resp, err := client.Get("https://google.com") //if no error then err becomes nil aka NULL
	fmt.Println("Error", err)                     //basically we can get an error if any
	fmt.Println("Response", resp)                 // we ca get the response if any
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

//functions can take named params like
//they will act as variables

func getDay2(day string) (weekday string, hour int) {
	switch day {
	case time.Tuesday.String():
		weekday, hour = "Tuesday", time.Now().Hour()
	default:
		weekday, hour = "Null", 1
	}
	return weekday, hour
}
