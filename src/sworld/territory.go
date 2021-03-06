package sworld

type Territory struct {
	head *RegionNode
}

func CreateTerritory() *Territory {
	t := new(Territory)
	t.head = nil
	return t
}

func (t *Territory) Add(id int) {
	if t == nil {
		return
	}

	new_node := new(RegionNode)
	new_node.region_id = id
	new_node.next = nil

	if t.head != nil {
		new_node.next = t.head
		t.head = new_node
	} else {
		t.head = new_node
	}
}

func (t *Territory) Remove(id int) bool {
	if t == nil {
		return false
	}

	cur := t.head
	var prev *RegionNode = nil

	for cur.next != nil {
		if cur.region_id == id {
			if prev != nil {
				prev.next = cur.next
			} else {
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

func (t *Territory) Has(id int) bool {
	cur := t.head
	for cur != nil {
		if cur.region_id == id {
			return true
		}
		cur = cur.next
	}
	return false
}

func (t *Territory) IsEmpty() bool {
	if t.head == nil {
		return true
	} else {
		return false
	}
}

func (t *Territory) List() []int {
	l := make([]int, 0)
	cur := t.head
	for cur != nil {
		l = append(l, cur.region_id)
		cur = cur.next
	}
	return l
}
