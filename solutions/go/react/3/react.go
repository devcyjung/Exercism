package react

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
    updatedCallbackList := c.cell.callbacks[:0]
	for _, cb := range c.cell.callbacks {
        if c.callback != cb {
            updatedCallbackList = append(updatedCallbackList, cb)
        }
    }
    c.cell.callbacks = updatedCallbackList
}

func (c *cell) Value() int {
	if len(c.deps) == 0 {
        return c.value
    }
    depValueList := make([]int, 0)
    for _, d := range(c.deps) {
        depValueList = append(depValueList, d.Value())
    }
    return (*c.computation)(depValueList...)
}

func (c *cell) SetValue(value int) {
	affectedChecked := make(map[*cell]bool)
    affectedChecked[c] = false
    for {
        addedCountPerCycle := 0
        for affected, checked := range affectedChecked {
            if checked {
                continue
            }
            for _, affectedByAffected := range (*c.reactor)[affected] {
                _, alreadyAdded := affectedChecked[affectedByAffected]
                if alreadyAdded {
                    continue
                }
                affectedChecked[affectedByAffected] = false
                addedCountPerCycle++
            }
            affectedChecked[affected] = true
        }
        if addedCountPerCycle == 0 {
            break
        }
    }

    prevValues := make(map[*cell]int)
    for affected := range affectedChecked {
        prevValues[affected] = affected.Value()
    }

    c.value = value

    for affected := range affectedChecked {
        if affected.Value() != prevValues[affected] {
            newValue := affected.Value()
            for _, cb := range affected.callbacks {
                (*cb)(newValue)
            }
        }
    }
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, &callback)
    var canclerInstance canceler
    canclerInstance.cell = c
    canclerInstance.callback = &callback
    return &canclerInstance
}

func New() Reactor {
    reactorInstance := make(reactor)
    return &reactorInstance
}

func (r *reactor) CreateInput(initial int) InputCell {
	var c cell
    c.value = initial
    c.deps = make([]*cell, 0)
    c.callbacks = make([]*func(int), 0)
    c.reactor = r
    return &c
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	var c cell
    c.deps = make([]*cell, 1)
    c.deps[0] = dep.(*cell)
    variadicCompute := func (deps ...int) int {
        if len(deps) != 1 {
            panic("Unexpected number of args. 1 expected.")
        }
        return compute(deps[0])
    }
    c.computation = &variadicCompute
    c.callbacks = make([]*func(int), 0)
    c.reactor = r

    for _, parent := range c.deps {
        siblings, siblingExists := (*r)[parent]
        if !siblingExists {
            (*r)[parent] = make([]*cell, 0)
        }
        (*r)[parent] = append(siblings, &c)
    }
    return &c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	var c cell
    c.deps = make([]*cell, 2)
    c.deps[0] = dep1.(*cell)
    c.deps[1] = dep2.(*cell)
    variadicCompute := func (deps ...int) int {
        if len(deps) != 2 {
            panic("Unexpected number of args. 2 expected.")
        }
        return compute(deps[0], deps[1])
    }
    c.computation = &variadicCompute
    c.callbacks = make([]*func(int), 0)
    c.reactor = r

    for _, parent := range c.deps {
        siblings, siblingExists:= (*r)[parent]
        if !siblingExists {
            (*r)[parent] = make([]*cell, 0)
        }
        (*r)[parent] = append(siblings, &c)
    } 
    return &c
}
