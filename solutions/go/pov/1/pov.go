package pov

import "slices"

type Tree struct {
    size		int
	value		string
    children	[]*Tree
    parent		*Tree
}

func New(value string, children ...*Tree) *Tree {
    size := 1
    tr := &Tree{0, value, children, nil}
    for _, child := range children {
        size += child.size
        child.parent = tr
    }
    tr.size = size
    return tr
}

func (tr *Tree) Value() string {
	return tr.value
}

func (tr *Tree) Children() []*Tree {
	return tr.children
}

func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

func (tr *Tree) getPath(to string) []*Tree {
    if tr == nil || tr.size == 0 {
        return nil
    }
    dfs := make([]*Tree, 0, tr.size)
    path := make([]*Tree, 0, tr.size)
    dfs = append(dfs, tr)
    var node *Tree
    var found bool
    for len(dfs) > 0 {
        node = dfs[len(dfs) - 1]
        dfs = dfs[:len(dfs) - 1]
        if node.value == to {
            found = true
            break
        }
        for _, child := range node.children {
            dfs = append(dfs, child)
        }
    }
    if !found {
        return nil
    }

    for node != nil {
        path = append(path, node)
        node = node.parent
    }
    slices.Clip(path)
    slices.Reverse(path)
    return path
}

func (tr *Tree) swap(child *Tree) {
    tr.size, child.size = tr.size - child.size, tr.size
    tr.parent, child.parent = child, tr.parent
    child.children = append(child.children, tr)
    tr.children = slices.DeleteFunc(tr.children, func(t *Tree) bool {
        return t == child
    })
}

func (tr *Tree) FromPov(from string) *Tree {
	path := tr.getPath(from)
    if path == nil || len(path) == 0 {
        return nil
    }
    if len(path) == 1 {
        return tr
    }
    for i, node := range path[:len(path) - 1] {
        node.swap(path[i + 1])
    }
    return path[len(path) - 1]
}

func (tr *Tree) PathTo(from, to string) []string {
	newTree := tr.FromPov(from)
    if newTree == nil {
        return []string{}
    }
    path := newTree.getPath(to)
    if path == nil || len(path) == 0 {
        return []string{}
    }
    result := make([]string, len(path))
    for i := range result {
        result[i] = path[i].value
    }
    return result
}