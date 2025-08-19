package react

type (
    reactor struct {
        inputCell   func(int) InputCell
		computeCell func(func() int) ComputeCell
    }
	inputCell struct {
    	getValue func() int
    	setValue func(int)
    }
    computeCell struct {
    	getValue    func() int
    	addCallback func(func(int)) Canceler
    }
    canceler func()
)

func New() Reactor {
	updateFns := make([]func(), 0)
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
						for i := index; i < len(updateFns); i++ {
                			updateFns[i]()
                		}
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

func (i *inputCell) Value() int {
	return i.getValue()
}

func (i *inputCell) SetValue(value int) {
	i.setValue(value)
}

func (c *computeCell) Value() int {
	return c.getValue()
}

func (c *computeCell) AddCallback(callback func(int)) Canceler {
	return c.addCallback(callback)
}

func (c canceler) Cancel() {
	c()
}
