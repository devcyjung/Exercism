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
	v1, ok1 := bill[item]
    v2, ok2 := units[unit]
    if !ok1 || !ok2 || v1 < v2 {
        return false
    }
    if v1 == v2 {
        delete(bill, item)
        return true
    }
    bill[item] -= v2
    return true
}

func GetItem(bill map[string]int, item string) (v int, ok bool) {
	v, ok = bill[item]
    return
}
