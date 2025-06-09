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
	affectedIsCheckedMap := map[*cell]bool{c: false}
    for {
        addedCountPerCycle := 0
        for affectedCell, isChecked := range affectedIsCheckedMap {
            if isChecked {
                continue
            }
            for _, affectedByAffected := range (*c.reactor)[affectedCell] {
                _, alreadyAdded := affectedIsCheckedMap[affectedByAffected]
                if alreadyAdded {
                    continue
                }
                affectedIsCheckedMap[affectedByAffected] = false
                addedCountPerCycle++
            }
            affectedIsCheckedMap[affectedCell] = true
        }
        if addedCountPerCycle == 0 {
            break
        }
    }

    affectedCellPrevValueMap := make(map[*cell]int)
    for affectedCell := range affectedIsCheckedMap {
        affectedCellPrevValueMap[affectedCell] = affectedCell.Value()
    }

    c.value = value

    for affectedCell := range affectedIsCheckedMap {
        if affectedCell.Value() != affectedCellPrevValueMap[affectedCell] {
            affectedCellNewValue := affectedCell.Value()
            for _, affectedCellCallback := range affectedCell.callbacks {
                (*affectedCellCallback)(affectedCellNewValue)
            }
        }
    }
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, &callback)
    var cancelerInstance canceler
    cancelerInstance.cell = c
    cancelerInstance.callback = &callback
    return &cancelerInstance
}

func New() Reactor {
    return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	var newCell cell
    newCell.value = initial
    newCell.deps = make([]*cell, 0)
    newCell.callbacks = make([]*func(int), 0)
    newCell.reactor = r
    return &newCell
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	var newCell cell
    newCell.deps = make([]*cell, 1)
    newCell.deps[0] = dep.(*cell)
    variadicCompute := func (deps ...int) int {
        if len(deps) != 1 {
            panic("Unexpected number of varargs. 1 expected.")
        }
        return compute(deps[0])
    }
    newCell.computation = &variadicCompute
    newCell.callbacks = make([]*func(int), 0)
    newCell.reactor = r

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
