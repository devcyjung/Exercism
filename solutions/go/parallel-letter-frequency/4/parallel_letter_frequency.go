package letter

import "sync"

type FreqMap map[rune]uint64

func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) FreqMap {
	frequencies := FreqMap{}
    var mu sync.Mutex
    var wg sync.WaitGroup
    channels := make([]chan FreqMap, len(texts))
    for i, t := range texts {
        channels[i] = make(chan FreqMap)
        wg.Add(1)
        go reader(frequencies, &mu, &wg , channels[i])
        go writer(t, channels[i])
    }
    wg.Wait()
    return frequencies
}

func reader(writeTo FreqMap, mu *sync.Mutex, wg *sync.WaitGroup , ch <-chan FreqMap) {
    readFrom := <-ch
    mu.Lock()
    defer func() {
        mu.Unlock()
        wg.Done()
    }()
    for k, v := range readFrom {
        writeTo[k] += v
    }
}

func writer(text string, ch chan<- FreqMap) {
    ch <- Frequency(text)
    close(ch)
}
