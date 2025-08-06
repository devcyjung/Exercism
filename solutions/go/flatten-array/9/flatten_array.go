package flatten

func Flatten(nested any) []any {
	switch value := nested.(type) {
    case []any:
        flatten := make([]any, 0)
        for _, elem := range value {
            flatten = append(flatten, Flatten(elem)...)
        }
        return flatten
    case any:
        return []any{value}
    default:
        return []any{}
    }
}