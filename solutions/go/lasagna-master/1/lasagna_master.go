package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTime int) int {
    if averageTime == 0 {
        averageTime = 2
    }
    return len(layers) * averageTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, v := range(layers) {
        switch v {
            case "noodles":
            	noodles += 50
            case "sauce":
            	sauce += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(yours []string, mine []string) {
    mine[len(mine) - 1] = yours[len(yours) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (ret []float64) {
    ret = make([]float64, len(quantities))
    for i := range(ret) {
        ret[i] = quantities[i] * float64(portions) / float64(2)
    }
    return 
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
