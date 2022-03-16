package main

import (
	"fmt"
	"math/cmplx"
	"time"
)

var (
	goIsFun  bool       = true                 //declaring a variable of type bool
	maxInt   uint64     = 1<<64 - 1            //declaring a variable of type uint64
	complexx complex128 = cmplx.Sqrt(-5 + 12i) //declaring a variable of type complex128
)

func main() {
	/*const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complexx, complexx)
	*/
	//how we can convert values in go
	lol := 10
	floated := float64(lol)
	u := uint(floated)
	fmt.Println(lol, floated, u)
	//learnAssertion()
	//learnstructs()
	ok()
}

func learnAssertion() {
	//if we have a value that you want to convert to a specific type we can use type assertion.
	//Takes a value and creates another version of it
	foo := map[string]interface{}{
		"Matt": 42,
	}
	timeMap(foo)
	fmt.Println(foo)
	//prints Matt:42: updated at time
}

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{}) //so we make y a map of interfaces essentially just making it  matt:42
	fmt.Println("Z =", z)               //z = map[Matt:42]
	fmt.Println("ok =", ok)             // ok = true
	if ok {                             //if ok is true
		z["iscool"] = false //we add a value called updated at
	}
}

type BootCamp struct {
	Lat  float64
	Lon  float64
	Date time.Time
}

type Point struct {
	X, Y int
}

func learnstructs() {
	fmt.Println(BootCamp{
		Lat:  34.123,
		Lon:  -112.23,
		Date: time.Now(),
	})

	//struct is a collection of fields and properties.
	//No need to define getters and setters in struct fields
	//structs can be used to define additional types in go

	//Example of Struct Literals
	//A struct literal sets a newly allowcated struct value by listing the values in its fields.
	var (
		a = Point{1, 2}  //has type point
		b = &Point{1, 2} // has type of *point
		c = Point{X: 1}  // X = 1 Y = 0
		d = Point{}      // X = 0 Y = 0
	)
	fmt.Println(a, b, c, d)

	//We can access fields using the dot notation

	event := BootCamp{
		Lat: 32.533,
		Lon: 243.3233,
	}
	event.Date = time.Now()
	fmt.Printf("Event on %s location (%f %f)",
		event.Date, event.Lat, event.Lon)
}

func ok() {
	var b int
	println(b)
}
