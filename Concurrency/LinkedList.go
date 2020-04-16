package Concurrency

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

func LinkedList() {
	n := Node{
		data: 12,
		prev: nil,
		next: nil,
	}

	n1 := Node{
		data: 11,
		prev: &n,
		next: nil,
	}
	n2 := Node{
		data: 3,
		prev: &n1,
		next: nil,
	}

	fmt.Println(n2.prev.data)
}

func maisn() {
	LinkedList()
}
