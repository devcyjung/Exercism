package etl

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
    for k, v := range in {
        for _, s := range v {
            out[string([]rune(s)[0] - 'A' + 'a')] = k
        }
    }
    return out
}
