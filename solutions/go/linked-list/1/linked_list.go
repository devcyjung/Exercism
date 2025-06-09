package linkedlist

import "errors"

type Node struct {
    Value	any
    next	*Node
    prev	*Node
}

type List struct {
    head	*Node
    tail	*Node
}

func NewList(elements ...any) *List {
    size := len(elements)
    if size == 0 {
        return &List{}
    }
    var prev *Node
    var node *Node
    var head *Node
    for i, v := range elements {
        if i == 0 {
            prev = &Node{
                Value:	v,
            }
            head = prev
        } else {
            node = &Node{
                Value:	v,
                prev:	prev,
            }
            prev.next = node
            prev = node
        }
    }
    return &List {
        head:	head,
        tail:	prev,
    }
}

func (n *Node) Next() *Node {
    if n == nil {
        return nil
    }
	return n.next
}

func (n *Node) Prev() *Node {
    if n == nil {
        return nil
    }
	return n.prev
}

func (l *List) Unshift(v any) {
    if l == nil {
        *l = List{}
    }
	first := l.head
    n := &Node{
        Value:	v,
        next:	first,
    }
    if first != nil {
    	first.prev = n   
    } else {
        l.tail = n
    }
    l.head = n
    return
}

func (l *List) Push(v any) {
    if l == nil {
        *l = List{}
    }
	last := l.tail
    n := &Node{
        Value:	v,
        prev:	last,
    }
    if last != nil {
    	last.next = n   
    } else {
        l.head = n
    }
    l.tail = n
    return
}

func (l *List) Shift() (v any, e error) {
    if l == nil {
        *l = List{}
    }
	first := l.head
    if first == nil {
        e = errors.New("shift on empty list")
        return
    }
    if first.next != nil {
        first.next.prev = nil
    } else {
        l.tail = nil
    }
    l.head = first.next
    v = first.Value
    return
}

func (l *List) Pop() (v any, e error) {
    if l == nil {
        *l = List{}
    }
	last := l.tail
    if last == nil {
        e = errors.New("pop on empty list")
        return
    }
    if last.prev != nil {
        last.prev.next = nil
    } else {
        l.head = nil
    }
    l.tail = last.prev
    v = last.Value
    return
}

func (l *List) Reverse() {
    if l == nil {
        *l = List{}
    }
    current := l.head
    for current != nil {
        current.prev, current.next, current = current.next, current.prev, current.next
    }
	l.head, l.tail = l.tail, l.head
    return
}

func (l *List) First() *Node {
    if l == nil {
        return nil
    }
	return l.head
}

func (l *List) Last() *Node {
    if l == nil {
        return nil
    }
	return l.tail
}
