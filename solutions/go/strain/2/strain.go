package strain

import "slices"

func Keep[S ~[]T, T any](collection S, predicate func(T) bool) (r S) {
    slices.Grow(r, len(collection))
    for _, v := range collection {
        if predicate(v) {
            r = append(r, v)
        }
    }
    slices.Clip(r)
    return
}

func Discard[S ~[]T, T any](collection S, predicate func(T) bool) (r S) {
    slices.Grow(r, len(collection))
    for _, v := range collection {
        if !predicate(v) {
            r = append(r, v)
        }
    }
    slices.Clip(r)
    return
}