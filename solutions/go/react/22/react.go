package react

func New() Reactor {
	updateFns := make([]func(), 0)
	updateAfter := func(index int) {
		for i := index; i < len(updateFns); i++ {
			updateFns[i]()
		}
	}
	return &reactor{
		inputCell: func(initial int) InputCell {
			value := initial
			index := len(updateFns)
			return &inputCell{
				getValue: func() int {
					return value
				},
				setValue: func(newValue int) {
					if value != newValue {
						value = newValue
						updateAfter(index)
					}
				},
			}
		},
		computeCell: func(compute func() int) ComputeCell {
			value := compute()
			callbacks := make(map[*func(int)]struct{})
			updateFns = append(updateFns, func() {
				newValue := compute()
				if value != newValue {
					value = newValue
					for cb := range callbacks {
						(*cb)(value)
					}
				}
			})
			return &computeCell{
				getValue: func() int {
					return value
				},
				addCallback: func(callback func(int)) Canceler {
					cb := &callback
					callbacks[cb] = struct{}{}
					return canceler(func() {
						delete(callbacks, cb)
					})
				},
			}
		},
	}
}

type reactor struct {
	inputCell   func(int) InputCell
	computeCell func(func() int) ComputeCell
}

func (r *reactor) CreateInput(initial int) InputCell {
	return r.inputCell(initial)
}

func (r *reactor) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	return r.computeCell(func() int {
		return compute(cell.Value())
	})
}

func (r *reactor) CreateCompute2(cell1, cell2 Cell, compute func(int, int) int) ComputeCell {
	return r.computeCell(func() int {
		return compute(cell1.Value(), cell2.Value())
	})
}

type inputCell struct {
	getValue func() int
	setValue func(int)
}

func (i *inputCell) Value() int {
	return i.getValue()
}

func (i *inputCell) SetValue(value int) {
	i.setValue(value)
}

type computeCell struct {
	getValue    func() int
	addCallback func(func(int)) Canceler
}

func (c *computeCell) Value() int {
	return c.getValue()
}

func (c *computeCell) AddCallback(callback func(int)) Canceler {
	return c.addCallback(callback)
}

type canceler func()

func (c canceler) Cancel() {
	c()
}
