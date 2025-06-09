package atbash

func Atbash(s string) string {
	var runes []rune
    var size int
	for _, v := range s {
        if !('a' <= v && v <= 'z') && !('A' <= v && v <= 'Z') && !('0' <= v && v<= '9') {
            continue
        }
        if size > 0 && size % 5 == 0 {
            runes = append(runes, ' ')
        }
        if 'a' <= v && v <= 'z' {
            runes = append(runes, 'z' - v + 'a')
            size++
        } else if 'A' <= v && v <= 'Z' {
            runes = append(runes, 'Z' - v + 'a')
            size++
        } else {
            runes = append(runes, v)
            size++
        }
    }
    return string(runes)
}
