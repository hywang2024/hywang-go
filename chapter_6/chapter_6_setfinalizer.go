package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	setF()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}

type Myint int

func findMyint(r *Myint) {
	log.Print("road: ", r)
}
func setF() {
	var my Myint = Myint(99)
	myi := &my
	runtime.SetFinalizer(myi, findMyint)
}
