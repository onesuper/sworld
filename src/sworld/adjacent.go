package sworld

import (
	"fmt"
)

type Neighbors struct {
	head *RegionNode
	tail *RegionNode
}

func CreateNeighbors() *Neighbors {
	n := new(Neighbors)
	n.head = nil
	n.tail = nil
	return n
}

func (n *Neighbors) Show() {
	if n == nil {
		return
	}

	p := n.head
	for p != nil {
		fmt.Printf("->%d ", p.region_id)
		p = p.next
	}
}

func (t *Neighbors) Add(id int) {
	if t == nil {
		return
	}

	new_node := new(RegionNode)
	new_node.region_id = id
	new_node.next = nil

	if t.head != nil {
		t.tail.next = new_node
		t.tail = new_node
	} else {
		t.head = new_node
		t.tail = new_node

	}
}
