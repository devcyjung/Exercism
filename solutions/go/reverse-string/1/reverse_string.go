package reverse

// import "golang.org/x/text/unicode/norm"

func Reverse(input string) string {
	// input = norm.NFC.String(input)
    runes := []rune(input)
    for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
