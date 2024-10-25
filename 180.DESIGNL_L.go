package main

type nodeL struct {
	val  int
	next *nodeL
}

type MyLinkedList struct {
	head, tail *nodeL
}

func ConstructorL() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this.head == nil {
		return -1
	}

	i := -1

	for c := this.head; c != nil; c = c.next {
		i++
		if index == i {
			return c.val
		}
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	l := &nodeL{val: val, next: nil}
	if this.head == nil {
		this.head, this.tail = l, l
		return
	}

	l.next = this.head
	this.head = l
}

func (this *MyLinkedList) AddAtTail(val int) {
	l := &nodeL{val: val, next: nil}

	if this.tail == nil {
		this.head, this.tail = l, l
		return
	}

	this.tail.next = l
	this.tail = l
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	i := -1
	n := &nodeL{val: val, next: nil}
	var prev *nodeL

	if this.head == nil && index == 0 {
		this.head, this.tail = n, n
		return
	} else if this.head == nil && index > 0 {
		return
	}

	for c := this.head; c != nil; c = c.next {
		i++

		if i == index {
			n.next = c

			if index == 0 {
				this.head = n
			}

			if prev != nil {
				prev.next = n
			}
			return
		}

		prev = c
	}

	if i+1 == index {
		this.tail.next = n
		this.tail = n
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	i := -1
	var prev *nodeL

	for c := this.head; c != nil; c = c.next {
		i++

		if i == index {
			if prev == nil {
				this.head = this.head.next
			} else {
				prev.next = c.next
			}

			if c.next == nil {
				this.tail = prev
			}

			return
		}

		prev = c
	}
}
