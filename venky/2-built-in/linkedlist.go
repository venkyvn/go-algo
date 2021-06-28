package main

import (
	"fmt"
)

func main() {

	//myLinkedList := NewLinkedList(3)
	//myLinkedList.append(5)
	//myLinkedList.append(7)
	//myLinkedList.append(9)
	//myLinkedList.append(9)
	//
	//myLinkedList.prepend(11)
	//myLinkedList.prepend(22)
	//myLinkedList.print()
	//myLinkedList.toPrint()
	//
	//myLinkedList.delete(11)
	//myLinkedList.delete(7)
	//myLinkedList.print()
	//myLinkedList.toPrint()
	//fmt.Println("---------------")
	//myLinkedList.print()
	//myLinkedList.add(0, 1110)
	//myLinkedList.print()
	//myLinkedList.toPrint()

	testDelete := NewLinkedList(1)
	testDelete.append(3)
	testDelete.append(5)
	testDelete.append(6)
	testDelete.toPrint()
	testDelete.delete(1)
	testDelete.delete(7)
	testDelete.delete(6)
	testDelete.toPrint()

}

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

func NewLinkedList(value int) *linkedList {
	return &linkedList{
		head:   NewNode(value),
		length: 1,
	}
}

func NewNode(value int) *node {
	return &node{
		value: value,
		next:  nil,
	}
}

type isLinkedList interface {
	prepend(value int)
	append(value int)
	print()
	delete(value int)
	add(index, value int)
}

func (l *linkedList) prepend(value int) {
	oldHead := l.head

	newHead := NewNode(value)
	l.head = newHead
	l.head.next = oldHead
	l.length++
}

func (l *linkedList) append(value int) {
	tail := l.head

	for tail.next != nil {
		tail = tail.next
	}

	tail.next = NewNode(value)
	l.length++
}

func (l *linkedList) print() {
	if l.length == 0 {
		fmt.Println("empty linkedlist")
	}
	tail := l.head
	for tail.next != nil {
		fmt.Printf("%d ---> ", tail.value)
		tail = tail.next
	}
	fmt.Printf("%d ---> \n", tail.value)
	fmt.Println("length: ", l.length)
}

func (l linkedList) toPrint() {
	if l.length == 0 {
		fmt.Println("empty linkedlist")
	}
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ---> ", toPrint.value)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Println()
}

func (l *linkedList) delete(value int) {
	if l.length == 0 {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.value != value {
		// last node
		if previousToDelete.next.next == nil {
			return
		}

		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l *linkedList) add(index, value int) {
	if index == 0 {
		l.prepend(value)
		return
	}

	previousNode := l.head
	for i := 0; i < index-1; i++ {
		previousNode = previousNode.next
	}
	tempNode := previousNode.next
	previousNode.next = NewNode(value)
	previousNode.next.next = tempNode
}
