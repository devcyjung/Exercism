package strain

import "slices"

func Keep[S ~[]E, E any](collection S, predicate func(E) bool) (r S) {
    slices.Grow(r, len(collection))
    for _, v := range collection {
        if predicate(v) {
            r = append(r, v)
        }
    }
    slices.Clip(r)
    return
}

func Discard[S ~[]E, E any](collection S, predicate func(E) bool) (r S) {
    slices.Grow(r, len(collection))
    for _, v := range collection {
        if !predicate(v) {
            r = append(r, v)
        }
    }
    slices.Clip(r)
    return
}