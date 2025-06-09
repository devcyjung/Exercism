package lasagna

func PreparationTime(layers []string, averageTime int) int {
    if averageTime == 0 {
        averageTime = 2
    }
    return len(layers) * averageTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range(layers) {
        switch layer {
        case "noodles":
            noodles += 50
        case "sauce":
            sauce += 0.2
        }
    }
    return
}

func AddSecretIngredient(yours []string, mine []string) {
    if len(mine) > 1 && len(yours) > 1 {
        mine[len(mine) - 1] = yours[len(yours) - 1]
    }
}

func ScaleRecipe(quantities []float64, portions int) (ret []float64) {
    ret = make([]float64, len(quantities))
    for i := range(ret) {
        ret[i] = quantities[i] * float64(portions) / 2
    }
    return 
}