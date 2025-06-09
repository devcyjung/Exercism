package react

import "slices"

type canceler struct {
    cell *cell
    callback *func(int)
}

type cell struct {
    value int
    computation *func(deps ...int) int
    deps []*cell
    callbacks []*func(int)
    reactor *reactor
}

type reactor map[*cell][]*cell

func (c *canceler) Cancel() {
    c.cell.callbacks = slices.DeleteFunc(c.cell.callbacks, func(cb *func(int)) bool {
        return c.callback == cb
    })
}

func (c *cell) Value() int {
	if len(c.deps) == 0 {
        return c.value
    }
    depValueList := make([]int, len(c.deps))
    for i, dep := range(c.deps) {
        depValueList[i] = dep.Value()
    }
    return (*c.computation)(depValueList...)
}

func (c *cell) SetValue(value int) {
	subscribers := map[*cell]bool{c: false}
    for i, init := 0, false; !(init && i == 0); i = 0 {
        init = true
        for child, isChecked := range subscribers {
            if isChecked {
                continue
            }
            for _, descendant := range (*c.reactor)[child] {
                _, alreadyAdded := subscribers[descendant]
                if alreadyAdded {
                    continue
                }
                subscribers[descendant] = false
                i++
            }
            subscribers[child] = true
        }
    }

    oldValues := make(map[*cell]int)
    for child := range subscribers {
        oldValues[child] = child.Value()
    }

    c.value = value

    for child := range subscribers {
        if child.Value() != oldValues[child] {
            newValue := child.Value()
            for _, cb := range child.callbacks {
                (*cb)(newValue)
            }
        }
    }
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, &callback)
    return &canceler{cell: c, callback: &callback}
}

func New() Reactor {
    return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
    return &cell{value: initial, reactor: r}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
    f := func(deps ...int) int {
        if len(deps) != 1 {
            panic("Unexpected number of varargs. 1 expected.")
        }
        return compute(deps[0])
    }
    newCell := cell{
        deps: []*cell{dep.(*cell)},
        reactor: r,
        computation: &f,
    }

    for _, parent := range newCell.deps {
        siblings, siblingExists := (*r)[parent]
        if !siblingExists {
            (*r)[parent] = make([]*cell, 0)
        }
        (*r)[parent] = append(siblings, &newCell)
    }
    return &newCell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	var newCell cell
    newCell.deps = make([]*cell, 2)
    newCell.deps[0] = dep1.(*cell)
    newCell.deps[1] = dep2.(*cell)
    variadicCompute := func (deps ...int) int {
        if len(deps) != 2 {
            panic("Unexpected number of varargs. 2 expected.")
        }
        return compute(deps[0], deps[1])
    }
    newCell.computation = &variadicCompute
    newCell.callbacks = make([]*func(int), 0)
    newCell.reactor = r

    for _, parent := range newCell.deps {
        siblings, siblingExists:= (*r)[parent]
        if !siblingExists {
            (*r)[parent] = make([]*cell, 0)
        }
        (*r)[parent] = append(siblings, &newCell)
    } 
    return &newCell
}
