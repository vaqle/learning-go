package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//learnMutuability()
	create()
}

func main1() {
	//REMEBER the type of the variable must be declared last
	//var name  string = "XD"
	learnPointers()
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

func zeroval(ival int) {
	ival = 0 //gets a copy of ival distinct from the one in the caling func
	// println("IVAL =", ival)
}

func zeroptr(iptr *int) { //takes in a int pointer
	*iptr = 0
}

func learnPointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) //this does not change the value of i so i is still = 1
	fmt.Println("zeroval:", i)

	zeroptr(&i)                //& is basically the memory of i when we do this it sets i to 0 we pass in the pointer of i
	fmt.Println("zeroptr:", i) // i is now 0
	//now &i = to its memory adress
	fmt.Println("pointer:", &i)

}

type Artist struct {
	name, genre string
	songs       int
}

func newRelease(a Artist) int {
	a.songs++
	return a.songs
}

func learnMutuability() {
	//Only constants are immutable
	//once it is created it cannot be changed
	me := Artist{name: "kaleb", genre: "RNB", songs: 100}
	fmt.Printf("%s released %dth song \n", me.name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.name, me.songs) //The Total Amount should be 101 but it isnt we need to use a pointer
}

//example

func newReleases(a *Artist) int {
	//* takes a pointer from our me variable
	a.songs++
	return a.songs
}

func create() {
	me := &Artist{name: "kaleb", genre: "RNB", songs: 100}           //& gets a pointer to the value
	fmt.Printf("%s released %dth song \n", me.name, newReleases(me)) //we need to mutuate it meaning copy
	fmt.Printf("%s has a total of %d songs", me.name, me.songs)      //The Total Amount is now 101
}
