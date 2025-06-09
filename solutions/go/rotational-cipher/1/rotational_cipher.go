package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	r := []rune(plain)
    for i, v := range r {
        if 'A' <= v && v <= 'Z' {
            r[i] += rune(shiftKey)
            if r[i] > 'Z' {
                r[i] -= 'Z'+1-'A'
            }
        } else if 'a' <= v && v <= 'z' {
            r[i] += rune(shiftKey)
            if r[i] > 'z' {
            	r[i] -= 'z'+1-'a'
            }
        }
    }
    return string(r)
}
