package main

import "fmt"

func main() {
	node := new(Node)
	node.data = 1

	node2 := new(Node)
	node2.data = 2
	node.next = node2

	showNode(node)
}

type Node struct {
	data int
	next *Node
}

func showNode(node *Node) {
	for node != nil {
		fmt.Println(*node)
		node = node.next
	}
}
