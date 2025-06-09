package tree

import (
    "cmp"
    "errors"
    "fmt"
    "slices"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

const (
    errDuplicateNodeFmt = "Duplicate nodes: %v %v"
    errInvalidRangeFmt = "Invalid Range in Record - ID: %d Parent: %d"
    errCycleFmt = "Non-root Record cycles to itself: %d %d"
)

func Build(records []Record) (*Node, error) {
	size := len(records)
    if size == 0 {
        return nil, nil
    }
    var err error
    slices.SortFunc(records, func(a, b Record) int {
        if cmp.Compare(a.ID, b.ID) == 0 {
            err = errors.Join(err, fmt.Errorf(errDuplicateNodeFmt, a, b))
        }
        return cmp.Compare(a.ID, b.ID)
    })
    if err != nil {
        return nil, err
    }
	nodes := make([]*Node, 0, size)
    for _, r := range records {
        if r.ID < 0 || size <= r.ID || r.Parent < 0 || size <= r.Parent || r.Parent > r.ID {
            return nil, fmt.Errorf(errInvalidRangeFmt, r.ID, r.Parent)
        }
        if r.ID != 0 && r.Parent == r.ID {
            return nil, fmt.Errorf(errCycleFmt, r.ID, r.Parent)
        }
        nodes = append(nodes, &Node{
            r.ID, nil,
        })
    }
    for i, r := range records {
        if r.Parent != i {
        	nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[i])
        }
    }
    return nodes[0], nil
}