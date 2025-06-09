#include "isogram.h"
#include <stddef.h>

bool is_isogram(const char phrase[]) {
    if (!phrase) {
        return false;
    }
    bool contains[26] = {false};
    size_t i = 0;
    char cur;
    while ((cur = phrase[i++])) {
        if (cur == ' ' || cur == '-') {
            continue;
        }
        if ('a' <= cur && cur <= 'z') {
            if (contains[cur - 'a']) {
                return false;
            }
            contains[cur - 'a'] = true;
        } else if ('A' <= cur && cur <= 'Z') {
            if (contains[cur - 'A']) {
                return false;
            }
            contains[cur - 'A'] = true;
        } else {
            return false;
        }
    }
    return true;
}