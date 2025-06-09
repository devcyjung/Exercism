package strain

func Keep[T any](collection []T, predicate func(T) bool) (r []T) {
    for _, v := range collection {
        if predicate(v) {
            r = append(r, v)
        }
    }
    return
}

func Discard[T any](collection []T, predicate func(T) bool) (r []T) {
    for _, v := range collection {
        if !predicate(v) {
            r = append(r, v)
        }
    }
    return
}