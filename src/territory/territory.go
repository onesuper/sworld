package territory

import (
	"fmt"
)

type Territory struct {
	head *Node
	tail *Node
}

type Node struct {
	region_id int
	next      *Node
}

func CreateTerritory() *Territory {
	t := new(Territory)
	t.head = nil
	t.tail = nil
	return t
}

func (t *Territory) Add(id int) {
	if t == nil {
		return
	}

	new_node := new(Node)
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

func (t *Territory) Remove(id int) bool {
	if t == nil {
		return false
	}

	cur := t.head
	var prev *Node = nil
	for cur.next != nil {
		if cur.region_id == id {
			if prev != nil {
				prev.next = cur.next // delete cur
			} else { // if cur is the head, need to modify t.head
				t.head = cur.next
			}
			return true
		} else {
			prev = cur
		}
		cur = cur.next
	}

	return false
}

func (t *Territory) HasRegion(id int) bool {
	p := t.head
	for p != nil {
		if p.region_id == id {
			return true
		}
		p = p.next
	}
	return false
}

func (t *Territory) Show() {
	p := t.head
	for p != nil {
		fmt.Printf("->%d ", p.region_id)
		p = p.next
	}
}
