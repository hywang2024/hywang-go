package main

import "fmt"

func main() {
	bird := new(Bird)
	bird.Fly.canFly()
	bird.canFly()
	bird.canWalk()

	peoples := new(Peoples)
	peoples.canWalk()
	peoples.canSwim()
}

type Fly struct {
}

func (fly *Fly) canFly() {
	fmt.Println("可以飞")
}

type Walk struct {
}

func (walk *Walk) canWalk() {
	fmt.Println("可以走")
}

type Swim struct {
}

func (swim *Swim) canSwim() {
	fmt.Println("可以游泳")
}

type Bird struct {
	Fly
	Walk
}

type Peoples struct {
	Walk
	Swim
}
