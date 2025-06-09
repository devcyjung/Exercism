package letter

import "reflect"

type FreqMap map[rune]int

func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) FreqMap {
	frequencies := FreqMap{}
    size := len(texts)
    channels := make([]chan FreqMap, size)
    for i, t := range texts {
        channels[i] = make(chan FreqMap)
        go worker(t, channels[i])
    }
    cases := make([]reflect.SelectCase, size)
    for i, ch := range channels {
        cases[i] = reflect.SelectCase{
            Dir:	reflect.SelectRecv,
            Chan:	reflect.ValueOf(ch),
        }
    }
    for size > 0 {
        index, val, ok := reflect.Select(cases)
        if !ok {
            cases[index].Chan = reflect.ValueOf(nil)
            size--
            continue
        }
        it := val.MapRange()
        for it.Next() {
            frequencies[rune(it.Key().Int())] += int(it.Value().Int())
        }
    }
    return frequencies
}

func worker(text string, ch chan<- FreqMap) {
    ch <- Frequency(text)
    close(ch)
}
