package linkedlist

import (
	"errors"
	"strconv"
)

type node[K any] struct {
	Val  K
	Prev *node[K]
	Next *node[K]
}

type LinkedList[K any] struct {
	head   *node[K]
	tail   *node[K]
	length int
}

func New[K any]() LinkedList[K] {
	return LinkedList[K]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *LinkedList[K]) Add(items ...K) {
	for _, item := range items {
		l.length++
		node := node[K]{
			Val:  item,
			Prev: nil,
			Next: nil,
		}

		if l.length == 1 {
			l.head = &node
			l.tail = &node
		} else {
			if l.length == 2 {
				l.tail = &node
				l.head.Next = l.tail
				l.tail.Prev = l.head
			} else {
				l.tail.Next = &node
				node.Prev = l.tail
				l.tail = &node
			}
		}
	}
}

func (l *LinkedList[K]) Get(index int) K {
	counter := 0
	currentNode := l.head

	for counter < index {
		currentNode = currentNode.Next
		counter++
	}

	return currentNode.Val
}

func (l *LinkedList[K]) Remove(index int) {
	l.length--
	if index == 0 {
		if l.length == 0 {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.Next
			l.head.Prev = nil
		}
	} else if index == l.length {
		l.tail = l.tail.Prev
		l.tail.Next = nil
	} else {
		counter := 0
		currentNode := l.head

		for counter < index {
			currentNode = currentNode.Next
			counter++
		}

		currentNode.Prev.Next = currentNode.Next
		currentNode.Next.Prev = currentNode.Prev
	}
}

func (l *LinkedList[K]) AddAll(anotherList []K) {
	for _, element := range anotherList {
		l.Add(element)
	}
}

func (l *LinkedList[K]) Size() int {
	return l.length
}

func (l *LinkedList[K]) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkedList[K]) Insert(index int, val K) error {
	counter := 0
	currentNode := l.head

	if index > l.length-1 {
		return errors.New("Index " + strconv.FormatInt(int64(index), 10) + " was out of bounds of list")
	}

	for counter < index {
		currentNode = currentNode.Next
		counter++
	}

	newNode := &node[K]{
		Val:  val,
		Prev: currentNode.Prev,
		Next: currentNode,
	}

	if currentNode.Prev == nil {
		l.head = newNode
	} else {
		currentNode.Prev.Next = newNode
	}
	l.length++

	return nil
}

func (l *LinkedList[K]) ForEach(predicate func(index int, item K)) {
	currentNode := l.head
	counter := 0

	for currentNode != nil {
		predicate(counter, currentNode.Val)
		counter++
		currentNode = currentNode.Next
	}
}

func (l *LinkedList[K]) Reverse() {
	var tmp *node[K]
	current := l.head

	for current != nil {
		tmp = current.Prev
		current.Prev = current.Next
		current.Next = tmp
		current = current.Prev
	}

	if l.length > 1 {
		l.head = tmp.Prev
	}
}

func (l *LinkedList[K]) ToSlice() []K {
	var listSlice []K
	l.ForEach(func(index int, item K) {
		listSlice = append(listSlice, item)
	})
	return listSlice
}

func (l *LinkedList[K]) GetFirst() K {
	return l.head.Val
}

func (l *LinkedList[K]) GetLast() K {
	return l.tail.Val
}
