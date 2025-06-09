package letter

import (
    "runtime"
    "strings"
    "unicode/utf8"
)

type FreqMap map[rune]int

var workerCount = runtime.NumCPU()

func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) FreqMap {

    totalRunes := 0
    for _, text := range texts {
        totalRunes += utf8.RuneCountInString(text)
    }

    chunkSizes := make([]int, workerCount)
    baseSize := totalRunes / workerCount
    for i := range chunkSizes {
        chunkSizes[i] = baseSize
    }
    for i := 0; i < totalRunes%workerCount; i++ {
        chunkSizes[i]++
    }
    chunks := make([]string, workerCount)
    currentChunk := 0
    runesInCurrent := 0
    var builder strings.Builder

    for _, text := range texts {
        for _, r := range text {
            if runesInCurrent == chunkSizes[currentChunk] {
                chunks[currentChunk] = builder.String()
                builder.Reset()
                currentChunk++
                runesInCurrent = 0
            }
            builder.WriteRune(r)
            runesInCurrent++
        }
    }

    if currentChunk < workerCount {
        chunks[currentChunk] = builder.String()
    }

    ch := make(chan FreqMap, workerCount)
    for _, chunk := range chunks {
        go func(c string) {
            ch <- Frequency(c)
        }(chunk)
    }

    total := FreqMap{}
    for i := 0; i < workerCount; i++ {
        freq := <-ch
        for k, v := range freq {
            total[k] += v
        }
    }

    return total
}
