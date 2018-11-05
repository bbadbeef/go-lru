package cache

import (
	"fmt"
	"reflect"
)

type listNode struct {
	item interface{}
	prev *listNode
	next *listNode
}

type List struct {
	head *listNode
	tail *listNode
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
	}
}

func (l *List) add(v interface{}) {
	n := &listNode{
		item: v,
		prev: nil,
		next: l.head,
	}
	if l.head != nil {
		l.head.prev = n
	}
	l.head = n
	if l.tail == nil {
		l.tail = n
	}
}

func (l *List) del(v interface{}) {
	for n := l.head; n != nil; n = n.next {
		if reflect.DeepEqual(n.item, v) {
			if n.prev != nil {
				n.prev.next = n.next
			}
			if n.next != nil {
				n.next.prev = n.prev
			}
			return
		}
	}
}

func (l *List) cut() interface{} {
	if l.head == nil {
		return nil
	}
	n := l.tail
	if l.head == l.tail {
		l.head, l.tail = nil, nil
		return n
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	return n.item
}

func (l *List) update(v interface{}) {
	for n := l.head; n != nil; n = n.next {
		if reflect.DeepEqual(n.item, v) {
			if n.prev == nil {
				return

			}
			n.prev.next = n.next
			if n.next != nil {
				n.next.prev = n.prev
			} else {
				l.tail = n.prev
			}
			n.prev = nil
			l.head.prev = n
			n.next = l.head
			l.head = n
			return
		}
	}
}

func (l *List) debug() {
	for n := l.head; n != nil; n = n.next {
		fmt.Printf("%v-->", n.item)
	}
	fmt.Printf("\n")
	for n := l.tail; n != nil; n = n.prev {
		fmt.Printf("%v<--", n.item)
	}
	fmt.Printf("\n")
}
