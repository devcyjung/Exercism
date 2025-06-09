package tree

import (
    "slices"
	"fmt"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (root *Node, e error) {
    size := len(records)
    if size == 0 {
        return
    }
    slices.SortFunc(records, func(a, b Record) int {
        if a.ID == b.ID {
            e = fmt.Errorf("Duplicate nodes: %v %v", a, b)
        }
        return a.ID - b.ID
    })
    if e != nil {
        return
    }
	var nodes []*Node
    for _, r := range records {
        if r.ID < 0 || size <= r.ID || r.Parent < 0 || size <= r.Parent || r.Parent > r.ID {
            e = fmt.Errorf("Invalid Range in Record - ID: %d Parent: %d", r.ID, r.Parent)
            return
        }
        if r.ID != 0 && r.Parent == r.ID {
            e = fmt.Errorf("Non-root Record cycles to itself: %d %d", r.ID, r.Parent)
        }
        nodes = append(nodes, &Node{
            ID:			r.ID,
            Children:	[]*Node{},
        })
    }
    for i, r := range records {
        if r.Parent != i {
        	nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[i])
        }
    }
    root = nodes[0]
    return
}
