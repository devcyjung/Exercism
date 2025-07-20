#include "binary.h"

int convert(const char *input) {
    if (!input) return INVALID;
    int acc = 0;
    int cur;
    for (const char *it = input; *it; ++it) {
        cur = *it - '0';
        if (cur >> 1) {
            return INVALID;
        }
        acc <<= 1;
        acc += cur;
    }
    return acc;
}