package flatten

import "slices"

func Flatten(nested any) []any {
	switch value := nested.(type) {
    case []any:
        flatten := make([][]any, len(value))
        for i, elem := range value {
            flatten[i] = Flatten(elem)
        }
        return append([]any{}, slices.Concat(flatten...)...)
    case any:
        return []any{value}
    default:
        return nil
    }
}