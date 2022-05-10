package main

import "fmt"

func main() {
	/*g := Game{
		name:     "Hello World",
		id:       1,
		released: false,
	}
	n := &g
	n.name = "Hello World 123"
	fmt.Println(g.name)/*/
	//xddd1()
	sayHello("Hello")
	fmt.Println("HI")
}

func sayHello(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", name)
	}
}

func xddd1() {
	a := []int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v) //this is the index and value of the array
	}
	//we can omit the index or value
	for _, v := range a {
		fmt.Println(v) //this is the value of the array
	}
	//we can omit the index
	for v := range a {
		fmt.Println(v) //this is the index of the array
	}
	//we can omit the value
	for i := range a {
		fmt.Println(i) //this is the index of the array
	}
}
func learnArra() {

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
}

// Game just a collection of fields
type Game struct {
	name     string
	id       int
	released bool
}
