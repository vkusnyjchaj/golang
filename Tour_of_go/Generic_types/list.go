package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Len() int {
	length := 0
	for current := l; current != nil; current = current.next {
		length++
	}
	return length
}

func (l *List[T]) Add(v T) *List[T] {
	// Create new node
	node := &List[T]{val: v}

	// Return new node if list is nil
	if l == nil {
		return node
	}

	// Go to the last element
	current := l
	for current.next != nil {
		current = current.next
	}

	// Add the new node to the end of the list
	current.next = node
	return l
}

func (l *List[T]) Remove() *List[T] {
	// Do nothing if l is nil or list has 1 element
	if l == nil || l.next == nil {
		return nil
	}

	// Go to the last element
	current := l
	for current.next.next != nil {
		current = current.next
	}

	// Remove the last element
	current.next = nil

	// Return pointer to current head
	return l
}

func (l *List[T]) GetAll() []T {
	var s []T
	for l != nil {
		s = append(s, l.val)
		l = l.next
	}

	return s
}

func main() {
	var list *List[int]
	list = list.Add(1)
	list = list.Add(2)
	list = list.Add(3)
	list = list.Remove()
	fmt.Printf("List: %v, Length: %d", list.GetAll(), list.Len())
}
