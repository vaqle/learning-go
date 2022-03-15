package main

import (
	"fmt"
	"math/cmplx"
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
}

func learn() {

}
