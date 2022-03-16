package main

import (
	"fmt"
	"time"
)

func main() {
	//switchWithoutCondition()
	//workWithPointers()
	//lol()
	//createAp()
	//doSomething()
	//array()
	//sliceLiterals()
	//nilSlices()
	sliceWithMake()
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

type Vertex struct {
	x, y int
}

func increaseVertexX(x int, vertex Vertex) int {
	vx := vertex.x
	return vx + 1
}

func doSomething() {
	vertex := Vertex{10, 1}
	//increaseVertexX(vertex.x, vertex) wont increase
	p := &vertex
	p.x = increaseVertexX(vertex.x, vertex)
	fmt.Println(vertex.x + vertex.y)
}

func array() {

	var a [10]string
	//a = array name 10 = size int = type
	a[0] = "xd"
	a[1] = "ok"
	strings := [3]string{"lol", "xd", "pov"}
	fmt.Println(strings)
	fmt.Println("_------------------------------")

	//Slices
	var s []string = strings[1:3]
	fmt.Println(s)
}

func nilSlices() {
	var s []int // slice
	s = append(s, 123)
	//fmt.Println("Info: ", s, len(s), cap(s))
	if s == nil {
		fmt.Println("Slice is nil")
	} else {
		fmt.Println(s)
	}
}

func sliceWithMake() {
	// 0 elements
	a := make([]int, 3, 5) // this is an array everything is zero valued
	length := len(a)
	capacity := cap(a)
	fmt.Println(length)   // 3
	fmt.Println(capacity) // 5
	fmt.Println(a)        // Prints 0,0,0,0 etc
	xddd()
	//In depth explanation:

	//len is just the amount of elements inside the array
	//capacity is how much elements it can hold
}

func xddd() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceLiterals() {
	s := []int{1, 2, 3, 4}
	//Slice literal
	fmt.Println(s)

	k := []struct {
		i int
		b bool
	}{{2, true}, {1, false}}
	fmt.Println(k)
}
