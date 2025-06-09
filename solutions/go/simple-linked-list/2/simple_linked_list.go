package linkedlist

import "errors"

type Element struct {
    value	int
    prev	*Element
}

type List struct {
    size	int
    tail	*Element
}

var (
    ErrNilList = errors.New("List is nil")
    ErrEmptyList = errors.New("List is empty")
)

func New(elements []int) *List {
    list := &List{}
	for _, element := range elements {
        list.Push(element)
    }
    return list
}

func (l *List) Size() int {
    switch l {
    case nil:
        return 0
    default:
        return l.size
    }
}

func (l *List) Push(element int) {
	if l == nil {
        return
    }
    node := &Element{value: element}
    switch l.tail {
    case nil:
        l.tail = node
    default:
        l.tail, node.prev = node, l.tail
    }
    l.size++
}

func (l *List) Pop() (int, error) {
	if l == nil {
        return 0, ErrNilList
    }
    last := l.tail
    switch last {
    case nil:
        return 0, ErrEmptyList
    default:
        l.tail, last.prev = last.prev, nil
        l.size--
        return last.value, nil
    }
}

func (l *List) Array() []int {
	result := make([]int, l.size)
    node := l.tail
    for idx := l.size - 1; idx >= 0; idx-- {
    	node, result[idx] = node.prev, node.value
    }
    return result
}

func (l *List) Reverse() *List {
	result := &List{}
    node := l.tail
    for node != nil {
        result.Push(node.value)
        node = node.prev
    }
    return result
}