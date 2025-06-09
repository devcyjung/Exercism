package flatten

import "slices"

func Flatten(nested any) []any {
	switch value := nested.(type) {
    case []any:
        flatten := make([][]any, len(value))
        for i, elem := range value {
            flatten[i] = Flatten(elem)
        }
        flattened := slices.Concat(flatten...)
        if flattened == nil {
            return []any{}
        }
        return flattened
    case any:
        return []any{value}
    default:
        return nil
    }
}