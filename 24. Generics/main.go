package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head *Element[T]
	tail *Element[T]
}

type Element[T any] struct {
	next *Element[T]
	val  T
}

func (list *List[T]) Push(value T) {
	if list.tail == nil {
		list.head = &Element[T]{val: value}
		list.tail = list.head
	} else {
		list.tail.next = &Element[T]{val: value}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	var elements []T

	for elem := list.head; elem != nil; elem = elem.next {
		elements = append(elements, elem.val)
	}

	return elements
}

func main() {
	m := map[int]string{
		1: "2",
		2: "4",
		4: "8",
	}
	fmt.Println("Keys:", MapKeys(m))

	_ = MapKeys(m)
	list := List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)
	fmt.Println("List:", list.GetAll())
}
