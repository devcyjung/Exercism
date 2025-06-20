package purchase

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseVehicle(option1, option2 string) string {
    suffix := " is clearly the better choice."
	if option1 < option2 {
        return option1 + suffix
    }
    return option2 + suffix
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	var discount float64
    if age < 3 {
        discount = 0.8
    } else if age < 10 {
        discount = 0.7
    } else {
        discount = 0.5
    }
    return discount * originalPrice
}
