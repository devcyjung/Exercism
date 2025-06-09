package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
        name, age, address,
    }
}

func (r *Resident) HasRequiredInfo() bool {
	return r != nil && r.Name != "" && r.Address["street"] != ""
}

func (r *Resident) Delete() {
    if r == nil {
        return
    }
	r.Name = ""
    r.Age = 0
    r.Address = nil
}

func Count(residents []*Resident) int {
	cnt := 0
    for _, r := range residents {
        if r != nil && r.HasRequiredInfo() {
            cnt++
        }
    }
    return cnt
}