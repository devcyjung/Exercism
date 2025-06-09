package react

type reactor struct {
    cells		[]Cell
    runner		func(int)
}

type inputCell struct {
    id			int
    value		int
    runner		func(int)
}

type computeCell struct {
    id			int
    value		int
    compute		func() int
    registry	map[Canceler]func(int)
}

type canceler struct {
    cancel		func()
}

func (c *canceler) Cancel() {
    c.cancel()
}

func New() Reactor {
    r := &reactor{
        cells:	make([]Cell, 0, 128),
    }
    r.runner = func(id int) {
        for _, cell := range r.cells[id:] {
            _ = cell.Value()
        }
    }
    return r
}

func (r *reactor) CreateInput(initial int) InputCell {
    input := &inputCell{
        id:		len(r.cells),
        value:	initial,
        runner:	r.runner,
    }
    r.cells = append(r.cells, input)
    return input
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
    com := &computeCell{
        id:			len(r.cells),
        value:		compute(dep.Value()),
        compute:	func() int { return compute(dep.Value()) },
        registry:	make(map[Canceler]func(int)),
    }
    r.cells = append(r.cells, com)
    return com
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
    com := &computeCell{
        id:			len(r.cells),
        value:		compute(dep1.Value(), dep2.Value()),
        compute:	func() int { return compute(dep1.Value(), dep2.Value()) },
        registry:	make(map[Canceler]func(int)),
    }
    r.cells = append(r.cells, com)
    return com
}

func (c *inputCell) Value() int {
    return c.value
}

func (c *inputCell) SetValue(value int) {
    if c.value == value {
        return
    }
    c.value = value
    c.runner(c.id)
}

func (c *computeCell) Value() int {
    newValue := c.compute()
    if c.value == newValue {
        return c.value
    }
    c.value = newValue
    for _, f := range c.registry {
        f(c.value)
    }
    return c.value
}

func (c *computeCell) AddCallback(callback func(int)) Canceler {
    newCanceler := &canceler{}
    newCanceler.cancel = func() {
        delete(c.registry, newCanceler)
    }
    c.registry[newCanceler] = callback
    return newCanceler
}