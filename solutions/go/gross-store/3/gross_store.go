package gross

func Units() map[string]int {
	return map[string]int {
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

func NewBill() map[string]int {
	return make(map[string]int)
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	v, ok := units[unit]
    if !ok {
        return false
    }
    bill[item] += v
    return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billCount, billHas := bill[item]
    unitSize, unitHas := units[unit]
    if !billHas || !unitHas || billCount < unitSize {
        return false
    }
    if billCount == unitSize {
        delete(bill, item)
        return true
    }
    bill[item] -= unitSize
    return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	billCount, billHas := bill[item]
    if !billHas {
        return 0, false
    }
    return billCount, true
}
