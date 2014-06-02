package sworld

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

func (t *Neighbors) List() []int {
	l := make([]int, 0)
	cur := t.head
	for cur != nil {
		l = append(l, cur.region_id)
		cur = cur.next
	}
	return l
}
