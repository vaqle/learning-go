package main

import (
	"fmt"
	"time"
)

func main() {
	//switchWithoutCondition()
	//workWithPointers()
	//lol()
	createAp()
}

func switchWithoutCondition() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	println(sum)
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

func workWithPointers() {

	var p *int
	i := 42
	p = &i //generate its operand
	//* denotes the pointers underlying value
	//so if i did this
	fmt.Println(p) //output would be 0xc000016088
	//however if i did this
	fmt.Println(*p) //output is 42

	//if we wanted to update p we would do this
	*p = 21     //this is dereferencing or inderecting
	println(*p) //outputs 21
	println(&i) //outputs addresss
}

func createAp() {
	var i, j int = 20, 25
	k := &j
	p := &i    //establish a pointer
	println(p) //should output adresss
	println(k) //should output adresss
	//now if we want to get i we would do this
	result := *p
	fmt.Println("Result of I:", result)
	*p = 10
	println(*p) //should be 10

}
func lol() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
