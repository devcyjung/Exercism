#include "pangram.h"
#include <stddef.h>

bool is_pangram(const char *const sentence) {
    if (!sentence) {
        return false;
    }
    size_t index = 0;
    char cur;
    bool map[26];
    while ((cur = sentence[index++])) {
        if ('a' <= cur && cur <= 'z') {
            map[cur - 'a'] = true;
        }
        if ('A' <= cur && cur <= 'Z') {
            map[cur - 'A'] = true;
        }
    }
    for (size_t i = 0; i < 26; ++i) {
        if (!map[i]) {
            return false;
        }
    }
    return true;
}