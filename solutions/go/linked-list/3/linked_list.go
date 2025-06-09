package linkedlist

import "errors"

var (
    ErrNilList = errors.New("List is nil")
    ErrEmptyList = errors.New("List is empty")
)

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
    list := &List{}
    var prev *Node
    var node *Node
    for i, v := range elements {
        prev = node
        node = &Node{Value: v}
        switch {
        case i == 0:
            list.head = node
        default:
            node.prev = prev
            prev.next = node
        }
    }
    list.tail = node
    return list
}

func (n *Node) Next() *Node {
    switch n {
    case nil:
        return nil
    default:
        return n.next
    }
}

func (n *Node) Prev() *Node {
    switch n {
    case nil:
        return nil
    default:
        return n.prev
    }
}

func (l *List) Unshift(v any) {
    if l == nil {
        return
    }
    node := &Node{Value: v}
    first := l.First()
    switch first {
    case nil:
        l.head, l.tail = node, node
        l.tail, l.head = node, node
    default:
        first.prev, node.next, l.head = node, first, node
    }
}

func (l *List) Push(v any) {
    if l == nil {
        return
    }
    node := &Node{Value: v}
    last := l.Last()
    switch last {
    case nil:
        l.head = node
        l.tail = node
    default:
        last.next, node.prev, l.tail = node, last, node
    }
}

func (l *List) Shift() (any, error) {
    if l == nil {
        return nil, ErrNilList
    }
    first := l.First()
    switch first {
    case nil:
        return nil, ErrEmptyList
    default:
        l.head = first.next
        switch l.head {
        case nil:
            l.tail = nil
        default:
            l.head.prev = nil
        }
        first.next = nil
        return first.Value, nil
    }
}

func (l *List) Pop() (v any, e error) {
    if l == nil {
        return nil, ErrNilList
    }
    last := l.Last()
    switch last {
    case nil:
        return nil, ErrEmptyList
    default:
        l.tail = last.prev
        switch l.tail {
        case nil:
            l.head = nil
        default:
            l.tail.next = nil
        }
        last.prev = nil
        return last.Value, nil
    }
}

func (l *List) Reverse() {
    if l == nil {
        return
    }
    node := l.First()
    for node != nil {
        node.prev, node.next, node = node.next, node.prev, node.next
    }
    l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
    switch l {
    case nil:
        return nil
    default:
        return l.head
    }
}

func (l *List) Last() *Node {
    switch l {
    case nil:
        return nil
    default:
        return l.tail
    }
}