package flatten

func Flatten(nested interface{}) (r []interface{}) {
    r = make([]interface{}, 0)
	switch t := nested.(type) {
        case []interface{}:
        	for _, v := range t {
                r = append(r, Flatten(v)...)
            }
        default:
        	if t != nil {
            	r = append(r, t)
            }
    }
    return
}
