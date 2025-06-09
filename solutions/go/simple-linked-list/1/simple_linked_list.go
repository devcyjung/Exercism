package linkedlist

import "errors"

type value = int
type index = int

type List struct {
    table 		map[index]value
    left, right index
}

func New(elements []int) *List {
    list := &List{
        left: 0,
        right: -1,
        table: make(map[index]value),
    }
    for _, v := range elements {
        list.Push(v)
    }
	return list
}

func (l *List) Size() int {
	return l.right - l.left + 1
}

func (l *List) Push(element int) {
	l.right++
    l.table[l.right] = element
}

func (l *List) Pop() (int, error) {
	if l.Size() == 0 {
        return 0, errors.New("pop() called on an empty list")
    }
    popped, _ := l.table[l.right]
    delete(l.table, l.right)
    l.right--
    return popped, nil
}

func (l *List) Array() []int {
	a := make([]int, 0, l.Size())
    var v int
    for k := l.left; k <= l.right; k++ {
        v, _ = l.table[k]
        a = append(a, v)
    }
    return a
}

func (l *List) Reverse() *List {
    n := New(nil)
	a := l.Array()
    for i := l.Size() - 1; i >= 0; i-- {
        n.Push(a[i])
    }
    return n
}