package list

import "errors"

type LinkedList struct {
	head     *Node
	inserted int
}

type Node struct {
	val  int
	next *Node
}

func (list *LinkedList) Add(val int) {
	newNode := &Node{val, nil}
	if list.inserted == 0 {
		list.head = newNode
	} else {
		aux := list.head
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = newNode
	}
	list.inserted++
}

func (list *LinkedList) AddOnIndex(val int, index int) error {
	if index >= 0 && index <= list.inserted {
		newNode := &Node{val, nil}
		if list.inserted == 0 {
			list.head = newNode
		} else if index == 0 {
			newNode.next = list.head
			list.head = newNode
		} else {
			prev := list.head
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
			newNode.next = prev.next
			prev.next = newNode
		}
		list.inserted++
		return nil
	} else {
		return errors.New("error msg")
	}
}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < list.inserted {
		if index == 0 {
			list.head = list.head.next
		} else {
			prev := list.head
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
			prev.next = prev.next.next
		}
		list.inserted--
		return nil
	} else {
		//error
		return errors.New("invalid index")
	}

}

func (list *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		cur := list.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		return cur.val, nil
	} else {
		if index < 0 {
			return index, errors.New("can't get value from linkedlist on negative index")
		} else {
			return index, errors.New("can't get value from linkedlist on index out of upper bounds")
		}
	}
}

func (list *LinkedList) Set(value, index int) error {
	if index >= 0 && index < list.inserted {
		target := list.head
		for i := 0; i < index; i++ {
			target = target.next
		}
		target.val = value
		return nil
	} else {
		return errors.New("error msg")
	}
}

func (list *LinkedList) Size() int {
	return list.inserted
}
