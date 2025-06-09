package flatten

func Flatten(nested any) []any {
    result := make([]any, 0, 1)
	switch value := nested.(type) {
    case []any:
        for _, elem := range value {
            result = append(result, Flatten(elem)...)
        }
    case any:
        result = append(result, value)
    default:
        return nil
    }
    return result
}
