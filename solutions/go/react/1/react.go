package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.
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

var react reactor

func (c *canceler) Cancel() {
    newcbs := c.cell.callbacks[:0]
	for _, cb := range c.cell.callbacks {
        if c.callback != cb {
            newcbs = append(newcbs, cb)
        }
    }
    c.cell.callbacks = newcbs
}

func (c *cell) Value() int {
	if len(c.deps) == 0 {
        return c.value
    }
    depInts := make([]int, 0)
    for _, d := range(c.deps) {
        depInts = append(depInts, d.Value())
    }
    return (*c.computation)(depInts...)
}

func (c *cell) SetValue(value int) {
	affected := make(map[*cell]bool)
    affected[c] = false
    for {
        added := 0
        for affectedcell, checked := range affected {
            if checked {
                continue
            }
            for _, dep := range react[affectedcell] {
                _, ok := affected[dep]
                if ok {
                    continue
                }
                affected[dep] = false
                added++
            }
            affected[affectedcell] = true
        }
        if added == 0 {
            break
        }
    }

    affectedValues := make(map[*cell]int)
    for i := range affected {
        affectedValues[i] = i.Value()
    }

    c.value = value

    for i := range affected {
        if i.Value() != affectedValues[i] {
            t := i.Value()
            for _, cb := range i.callbacks {
                (*cb)(t)
            }
        }
    }
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, &callback)
    var can canceler
    can.cell = c
    can.callback = &callback
    return &can
}

func New() Reactor {
	react = make(map[*cell][]*cell)
    return &react
}

func (r *reactor) CreateInput(initial int) InputCell {
	var c cell
    c.value = initial
    c.deps = make([]*cell, 0)
    c.callbacks = make([]*func(int), 0)
    c.reactor = &react
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
    c.reactor = &react

    for _, d := range c.deps {
        arr, ok := (*r)[d]
        if !ok {
            (*r)[d] = make([]*cell, 0)
        }
        (*r)[d] = append(arr, &c)
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
    c.reactor = &react

    for _, d := range c.deps {
        arr, ok := (*r)[d]
        if !ok {
            (*r)[d] = make([]*cell, 0)
        }
        (*r)[d] = append(arr, &c)
    } 
    return &c
}
