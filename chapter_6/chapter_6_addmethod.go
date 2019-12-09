package main

import "fmt"

func main() {
	var i MyInt = 100
	s := i.isZero()
	add := i.add(1)
	fmt.Println(s, add)

	var delegate func(i int)
	cla := new(class)
	delegate = cla.do
	delegate(100)

	delegate = do
	delegate(101)
}

type MyInt int

func (myint MyInt) isZero() bool {
	return myint == 0
}

func (myint MyInt) add(other int) int {
	return other + int(myint)
}

type class struct {
}

func (cla class) do(i int) {
	fmt.Println("do param:", i)
}

func do(i int) {
	fmt.Println("func do param:", i)
}
