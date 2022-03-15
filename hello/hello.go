package main

import (
	"fmt"
	"time"
)

const TEST = "TEST"

func main() {
	fmt.Println("hello world")
	//constantsandloops()
	//main2()
	//ifstatesandswitches()
	//learnArrays()
}

func learnArrays() {
	//we can specify the amount of values an array can hold
	var a [5]int // assigning it as int means default values are 0
	fmt.Println("ok: ", a)
	//we can update the key in the array
	a[4] = 10
	fmt.Println("ok: ", a)
	//We can get the length of the array
	length := len(a)
	println("Length: ", length)
	//We can create an array in one line
	b := [5]int{1, 2, 3, 4, 5} //we create an array and set the values
	fmt.Println("The Values ", b)
	//We will see slices rather then arrays more
	var twoD [2][6]int //first int is the amount of arrays second is the amount of values inside the array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = 1
		}
	}
	fmt.Println("Array Test", twoD) // outputs Array Test [[1 1 1] [1 1 1]]
}

func constantsandloops() {
	//Constans can be declared inside functions and outside of functions
	//for example
	//Also constants dont have a specific type unless declared
	const n = 3
	const a = 1
	p := n + a
	println(p) // prints 3 + 1 = 4
	//for loops with no condition will loop until u declare break
	//ex of for loops
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ // 1 - 10
	}
	//var a = 1000
	//var lol int = a <= 100 ? 50 : a
	//println(lol)
	b := 1
	for b <= 15 {
		b++
		println(b)
		if b == 12 {
			//if it hits it will stop looping
			break
		}
	}
	//example of continue
	println("-----------------------------------------------------")
	h := 1
	for h <= 12 {
		h++
		println(h)
		if h*10 > 2*10 {
			continue //so it will print some bc they r greater then 20
		}
	}
}

func ifstatesandswitches() {
	const value = 10
	i := 1
	if i+1 < 10 {
		println("Statement is true")
	} else {
		println("Statement is false")
	}
	//we can declare vars in if statements
	if n := 10; n < 0 {
		fmt.Println("Value is less than 0")
	} else if n < 10 {
		fmt.Println("Value less then 10")
	} else {
		fmt.Println("Value is 10 or greater")
	}
	fmt.Println("-----------------------------------------------------")
	v := 100
	switch v {
	case 100:
		fmt.Println("V = 100")
	case 96:
		fmt.Println("v = 96")
	}
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Today is Monday")
	//example of commas for multi expressions
	case time.Saturday, time.Friday:
		fmt.Println("Today is Sat or Friday")
	default:
		fmt.Println("Day could not be determined")
	}
	//type switch statements they compare values
	function := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("i is int")
		case string:
			fmt.Println("I is string")
		}
	}
	function(10)
	function("XD")
}

func learnSlices() {
	//create a slice

}

func main2() {
	a := 1
	//Strings Can be Added Together with +
	//for instance in php it  would be like "hello" . "GO" in go it would be "YO" + "go"
	fmt.Println("go" + "lang")
	//prints 1+1 = 2
	fmt.Println("1+1 =", 1+1)
	//we can declare variables by using var
	fmt.Println(a)
	//We can declare multiple variables like this
	var b, c int = 1, 2 // 1 = b 2 = c
	fmt.Println(b + c)  // we add both variables
	//Variables that are not initalized are represented as 0 for example:
	var d int // d = 0
	println(d)
	//We can also declare variables like this
	f := "apple"
	println(f) // prints apple
}
