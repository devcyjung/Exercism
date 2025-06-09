package series

func All(n int, s string) (result []string) {
	length := n
    if length > len(s) {
        return
    }
    result = make([]string, 0, len(s) - length + 1)
    for i := 0; i + length <= len(s); i++ {
        result = append(result, s[i:i + length])
    }
    return
}

func UnsafeFirst(n int, s string) (result string) {
	length := n
    if length > len(s) {
        return
    }
    result = s[:length]
    return
}

func First(n int, s string) (string, bool) {
    ok := n <= len(s)
    return UnsafeFirst(n, s), ok
}