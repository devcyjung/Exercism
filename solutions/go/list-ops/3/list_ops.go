package listops

import "slices"

type IntList = List[int, int]

type List[E, R any] []E

func (s List[E, R]) Foldl(foldFn func(acc R, elem E) R, initial R) R {
    acc := initial
	for _, elem := range s {
        acc = foldFn(acc, elem)
    }
    return acc
}

func (s List[E, R]) Foldr(foldFn func(elem E, acc R) R, initial R) R {
    acc := initial
	for i := len(s) - 1; i >= 0 ; i-- {
        acc = foldFn(s[i], acc)
    }
    return acc
}

func (s List[E, R]) Filter(filterFn func(elem E) bool) List[E, R] {
    result := slices.Clone(s)
    result = slices.DeleteFunc(result, func(elem E) bool {
        return !filterFn(elem)
    })
    return result
}

func (s List[E, _]) Length() int {
	return len(s)
}

func (s List[E, R]) Map(mapFn func(elem E) R) List[R, R] {
    result := make(List[R, R], 0, len(s))
    for _, elem := range s {
        result = append(result, mapFn(elem))
    }
    return result
}

func (s List[E, R]) Reverse() List[E, R] {
    result := slices.Clone(s)
    slices.Reverse(result)
    return result
}

func (s List[E, R]) Append(lst ...List[E, R]) List[E, R] {
    result := slices.Concat(s, slices.Concat(lst...))
    if result == nil {
        result = make(List[E, R], 0)
    }
    return result
}

func (s List[E, R]) Concat(lists []List[E, R]) List[E, R] {
    result := slices.Concat(s, slices.Concat(lists...))
    if result == nil {
        result = make(List[E, R], 0)
    }
    return result
}