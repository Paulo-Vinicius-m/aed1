package list

import (
	"errors"
	"fmt"
)

type DoubleLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
	size int
}

type DoubleNode struct {
	val  int
	next *DoubleNode
	prev *DoubleNode
}

func (list *DoubleLinkedList) Add(value int) {
	fmt.Print("Add", list.size, "\n")
	if list.size == 0 {
		list.tail = &DoubleNode{val: value, prev: nil, next: nil}
		list.head = list.tail
	} else {
		list.tail.next = &DoubleNode{val: value, prev: list.tail, next: nil}
		list.tail = list.tail.next
	}
	list.size++
}

func (list *DoubleLinkedList) AddOnIndex(value int, index int) error {
	fmt.Print("AddOnIndex ", index, list.size, "\n")
	if index < 0 || index > list.size {
		return errors.New("list add index out of bounds")
	} else if list.size == 0 {
		list.tail = &DoubleNode{val: value, prev: nil, next: nil}
		list.head = list.tail
	} else if index == 0 {
		list.head.prev = &DoubleNode{val: value, next: list.head, prev: nil}
		list.head = list.head.prev
	} else if index == list.size || index == -1 {
		list.tail.next = &DoubleNode{val: value, next: nil, prev: list.tail}
		list.tail = list.tail.next
	} else if index <= (list.size / 2) {
		target := list.head

		for i := 1; i < index; i++ {
			target = target.next
		}
		newNode := &DoubleNode{val: value, next: target.next, prev: target}
		target.next = newNode

	} else {
		target := list.tail

		for i := index + 1; i < list.size; i++ {
			target = target.prev
		}
		newNode := &DoubleNode{val: value, next: target, prev: target.prev}
		target.prev = newNode

	}
	list.size++
	return nil
}

func (list *DoubleLinkedList) RemoveOnIndex(index int) error {
	fmt.Print("RemoveOnIndex ", index, list.size, "\n")
	if index < 0 || index >= list.size {
		return errors.New("list remove index out of bounds or empty list")
	} else if list.size == 1 {
		list.tail = nil
		list.head = nil
	} else if index == 0 {
		list.head = list.head.next
		list.head.prev = nil
	} else if index == list.size-1 || index == -1 {
		list.tail = list.tail.prev
		list.tail.next = nil
	} else if index <= (list.size / 2) {
		target := list.head

		for i := 0; i < index; i++ {
			target = target.next
		}
		target.next.prev = target.prev
		target.prev.next = target.next

	} else {
		target := list.tail

		for i := index; i < list.size; i++ {
			target = target.prev
		}
		target.next.prev = target.prev
		target.prev.next = target.next

	}
	list.size--
	return nil
}

func (list *DoubleLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.size {
		return -1, errors.New("list get index out of bounds")
	} else if index <= (list.size / 2) {
		target := list.head

		for i := 0; i < index; i++ {
			target = target.next
		}
		return target.val, nil
	} else {
		target := list.tail

		for i := index + 1; i < list.size; i++ {
			target = target.prev
		}
		return target.val, nil
	}
}

func (list *DoubleLinkedList) Set(value, index int) error {
	fmt.Print("Set ", index, value, list.size, "\n")
	if index < 0 || index >= list.size {
		return errors.New("list set index out of bounds")
	} else if index <= (list.size / 2) {
		target := list.head

		for i := 0; i < index; i++ {
			target = target.next
		}
		target.val = value

	} else {
		target := list.tail

		for i := index + 1; i < list.size; i++ {
			target = target.prev
		}
		target.val = value

	}
	return nil
}

func (list *DoubleLinkedList) Size() int {
	return list.size
}
