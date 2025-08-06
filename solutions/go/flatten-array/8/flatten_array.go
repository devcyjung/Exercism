package flatten

import "slices"

func Flatten(nested any) []any {
    if f := flat(nested); f == nil {
        return []any{}
    } else {
        return f
    }
}

func flat(nested any) []any {
	switch value := nested.(type) {
    case []any:
        flatten := make([][]any, len(value))
        for i, elem := range value {
            flatten[i] = flat(elem)
        }
        return slices.Concat(flatten...)
    case any:
        return []any{value}
    default:
        return nil
    }
}