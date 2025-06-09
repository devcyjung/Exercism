package flatten

func Flatten(nested any) (r []any) {
    r = []any{}
	switch t := nested.(type) {
        case []any:
        	for _, v := range t {
                r = append(r, Flatten(v)...)
            }
        case any:
            r = append(r, t)
    }
    return
}
