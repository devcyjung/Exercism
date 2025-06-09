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
	chunk := strings.Join(texts, "")
    partitionIndices := make([]int, workerCount - 1)
    for i := range partitionIndices {
        partIdx := len(chunk) * i / workerCount
        for !utf8.RuneStart(chunk[partIdx]) {
            partIdx++
        }
        partitionIndices[i] = partIdx
    }
    parts := make([]string, workerCount)
    channel := make(chan FreqMap, workerCount)
    for i := range parts {
        switch i {
        case 0:
            parts[i] = chunk[:partitionIndices[0]]
        case len(parts) - 1:
            parts[i] = chunk[partitionIndices[i-1]:]
        default:
            parts[i] = chunk[partitionIndices[i-1]:partitionIndices[i]]
        }
        go func(part string) {
            channel <- Frequency(part)
        }(parts[i])
    }
    totalFrequency := make(FreqMap)
    for i := 0; i < workerCount; i++ {
        freq := <- channel
        for k, v := range freq {
            totalFrequency[k] += v
        }
    }
    return totalFrequency
}
