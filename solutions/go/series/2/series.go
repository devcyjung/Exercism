package series

func All(n int, s string) []string {
	r := []string{}
    if n > len(s) {
        return r
    }
    for i := 0; i+n <= len(s); i++ {
        r = append(r, s[i:i+n])
    }
    return r
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func SafeFirst(n int, s string) (r string, b bool) {
    if n > len(s) {
        return
    }
    r, b = s[:n], true
    return
}
