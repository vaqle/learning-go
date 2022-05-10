package main

import "fmt"

func main() {
	//swap2("HEY", "YOU")
	//varsAgain()
	//forLoopsagain()
	//relearnStuff()
	//exampe()
	exampe2()
}

func forLoopsagain() {
	amount := 1
	for amount < 10 {
		amount += 1
		if amount == 3 {
			fmt.Println("BROKE AT 3")
			break
		}
	}
}

func varsAgain() {
	var a, b, c = 100, 1, 1
	added := a + b + c

	fmt.Println(added)

}

func swap2(lol, poop string) {
	fmt.Println(lol + poop)
}

func relearnStuff() {

	var lol string = "hey"
	var pointer = &lol   // pointer is a memory address
	fmt.Println(lol)     // prints "hey"
	fmt.Println(pointer) // prints 0xc000088220
	//if we wanted to get the value of pointer we would need to dereference it
	//adding *before the pointer will dereference it
	fmt.Println(*pointer) // prints "hey"
	//now if we wanted to change the value of the variable we would need to dereference it as well
	*pointer = "you"
	fmt.Println(lol) // prints "you"
	//now if we print lol again it will print "you" cuz we changed the value of the variable
	fmt.Println(lol)

}

/**


When you write a function,
you can define arguments to be passed ether by value, or by reference.
Passing by value means that a copy of that value is sent to the function,
and any changes to that argument within that function only effect that variable within that function,
and not where it was passed from. However, if you pass by reference, meaning you pass a pointer to that argument,
you can change the value from within the function, and also change the value of the original variable that was passed in.

*/

type Object struct {
	name string
}

type List struct {
	objects []Object
}

func exampe() {
	var lol Object = Object{name: "lol"}
	fmt.Println("LOL: ", lol.name)
	//NOW if we w
	//changeValueOfObject(lol)
}

//THIS will not change the value of the original variable lol
func changeValueOfObject(name Object) {
	name.name = "you"
	fmt.Println("CHANGED: ", name.name)
}

//Now if we wanted to change the value of the original variable lol
//we would need to dereference it as well

func exampe2() {
	var lol Object = Object{name: "lol"}
	fmt.Println("FIRST: ", lol.name)
	changeIt(&lol)
	fmt.Println("THIRD", lol)
}

//CHANGES TGHE ORIGINAL VARIABLE LOL
func changeIt(name *Object) {
	name.name = "HELLO"
	fmt.Println("LAST: ", name)
	do()
}

func do() {
	//we can do this
	var name *Object
	fmt.Println("LOL: ", name)
	name.name = "HEY"
	fmt.Println("HEY: ", name)
	//we are passing a pointer to the variable name

}
