package letter

/*
#cgo CFLAGS: -march=native -O3 -msse4.2
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <emmintrin.h>

// Decode a single UTF-8 rune and return codepoint + bytes read
int decode_utf8(const char* s, int* rune, int* size) {
	unsigned char c = (unsigned char)s[0];

	if (c < 0x80) {
		*rune = c;
		*size = 1;
		return 0;
	} else if ((c & 0xE0) == 0xC0) {
		*rune = ((int)(c & 0x1F) << 6) | ((int)(s[1] & 0x3F));
		*size = 2;
		return 0;
	} else if ((c & 0xF0) == 0xE0) {
		*rune = ((int)(c & 0x0F) << 12) | ((int)(s[1] & 0x3F) << 6) | ((int)(s[2] & 0x3F));
		*size = 3;
		return 0;
	} else if ((c & 0xF8) == 0xF0) {
		*rune = ((int)(c & 0x07) << 18) | ((int)(s[1] & 0x3F) << 12) | ((int)(s[2] & 0x3F) << 6) | ((int)(s[3] & 0x3F));
		*size = 4;
		return 0;
	}
	return -1; // Invalid UTF-8
}

// Fills `output_keys` and `output_counts` with runes and their counts.
// Returns number of unique runes.
int count_runes(const char* s, int len, int* output_keys, int* output_counts, int max_entries) {
	int size = 0;
	int ascii_hist[128] = {0};

	for (int i = 0; i < len;) {
		if (len - i >= 16) {
			__m128i block = _mm_loadu_si128((const __m128i*)(s + i));
			__m128i high_bits = _mm_and_si128(block, _mm_set1_epi8(0x80));
			int mask = _mm_movemask_epi8(high_bits);

			if (mask == 0) {
				for (int j = 0; j < 16; ++j) {
					unsigned char ch = (unsigned char)s[i + j];
					++ascii_hist[ch];
				}
				i += 16;
				continue;
			}
		}

		int cp, step;
		if (decode_utf8(s + i, &cp, &step) != 0)
			break;
		i += step;

		int found = 0;
		for (int j = 0; j < size; ++j) {
			if (output_keys[j] == cp) {
				++output_counts[j];
				found = 1;
				break;
			}
		}
		if (!found && size < max_entries) {
			output_keys[size] = cp;
			output_counts[size] = 1;
			++size;
		}
	}

	for (int ch = 0; ch < 128; ++ch) {
		if (ascii_hist[ch] == 0) continue;
		int found = 0;
		for (int j = 0; j < size; ++j) {
			if (output_keys[j] == ch) {
				output_counts[j] += ascii_hist[ch];
				found = 1;
				break;
			}
		}
		if (!found && size < max_entries) {
			output_keys[size] = ch;
			output_counts[size] = ascii_hist[ch];
			++size;
		}
	}
	return size;
}
*/
import "C"
import (
	"sync"
	"unsafe"
)

type FreqMap map[rune]int

func countRunesC(s string) FreqMap {
	const maxEntries = 2048

	ptr := (*C.char)(unsafe.Pointer(&[]byte(s)[0]))
	length := C.int(len(s))

	keys := (*C.int)(C.malloc(C.size_t(maxEntries) * C.size_t(C.sizeof_int)))
	counts := (*C.int)(C.malloc(C.size_t(maxEntries) * C.size_t(C.sizeof_int)))
	defer C.free(unsafe.Pointer(keys))
	defer C.free(unsafe.Pointer(counts))

	n := C.count_runes(ptr, length, keys, counts, maxEntries)

	m := make(FreqMap)
	sliceKeys := (*[1 << 20]C.int)(unsafe.Pointer(keys))[:n:n]
	sliceCounts := (*[1 << 20]C.int)(unsafe.Pointer(counts))[:n:n]

	for i := 0; i < int(n); i++ {
		m[rune(sliceKeys[i])] = int(sliceCounts[i])
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	ch := make(chan FreqMap, len(texts))

	for _, text := range texts {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			ch <- countRunesC(t)
		}(text)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := make(FreqMap)
	for m := range ch {
		for r, c := range m {
			result[r] += c
		}
	}
	return result
}

func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}