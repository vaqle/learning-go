package main

import "fmt"

func main() {
	//kk()
	fmt.Println(countString("Yay"))
	//So we created a method called tooString
	L := Lol{
		Name: "Yay",
		Age:  18,
	}
	fmt.Println(L.tooString())
}

type Lol struct {
	Name string
	Age  int
}

func countString(s string) (amount int) {
	lenght := len(s)
	for i := 0; i < lenght; i++ {
		amount++
	}
	return amount
}

func (l Lol) tooString() string {
	return l.Name
}

func kk() {
	m := make(map[string]Lol)
	m["lol"] = Lol{
		Name: "lol",
		Age:  1,
	}
	fmt.Println(m["lol"])
}
