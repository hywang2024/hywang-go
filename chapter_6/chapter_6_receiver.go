package main

import "fmt"

//接收器
func main() {
	bag := &Bag{}
	addItemToBag(bag, "apple")
	fmt.Println(bag)
	ba := new(Bag)
	ba.AddItemToBagReceiver("banana")
	fmt.Println(ba)
}

type Bag struct {
	item []string
}

func addItemToBag(bag *Bag, item string) {
	bag.item = append(bag.item, item)
}

//(bag *Bag) 表示接收器
func (bag *Bag) AddItemToBagReceiver(items string) {
	bag.item = append(bag.item, items)
}
