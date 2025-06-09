package listops

type IntList = List[int, int]

type List[T, R any] []T

func (s List[T, R]) Foldl(fn func(R, T) R, initial R) R {
	for _, v := range s {
        initial = fn(initial, v)
    }
    return initial
}

func (s List[T, R]) Foldr(fn func(T, R) R, initial R) R {
	for i := len(s)-1; i >= 0 ; i-- {
        initial = fn(s[i], initial)
    }
    return initial
}

func (s List[T, R]) Filter(fn func(T) bool) (r List[T, R]) {
    if s == nil || len(s) == 0 {
        return s
    }
	for _, v := range s {
        if fn(v) {
            r = append(r, v)
        }
    }
    return
}

func (s List[T, _]) Length() int {
	return len(s)
}

func (s List[T, R]) Map(fn func(T) T) (r List[T, R]) {
    if s == nil || len(s) == 0 {
        return s
    }
    for _, v := range s {
        r = append(r, fn(v))
    }
    return
}

func (s List[T, R]) Reverse() (r List[T, R]) {
    if s == nil || len(s) == 0 {
        return s
    }
	for i := len(s)-1; i >= 0 ; i-- {
        r = append(r, s[i])
    }
    return
}

func (s List[T, R]) Append(lst List[T, R]) (r List[T, R]) {
    if s == nil && lst == nil || len(s) == 0 && len(lst) == 0 {
        return s
    }
	for _, v := range s {
        r = append(r, v)
    }
    for _, v := range lst {
        r = append(r, v)
    }
    return r
}

func (s List[T, R]) Concat(lists []List[T, R]) (r List[T, R]) {
    if s == nil && lists == nil || len(s) == 0 && len(lists) == 0 {
        return s
    }
    for _, v := range s {
        r = append(r, v)
    }
	for _, v := range lists {
        r = r.Append(v)
    }
    return
}
