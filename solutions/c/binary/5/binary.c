#include "binary.h"
#include <stddef.h>

int convert(const char *const input) {
    if (!input) return INVALID;
    int acc = 0;
    size_t index = 0;
    char cur;
    while ((cur = input[index++])) {
        switch (cur) {
        case '0':
            acc <<= 1;
            break;
        case '1':
            acc <<= 1;
            acc |= 1;
            break;
        default:
            return INVALID;
        }
    }
    return acc;
}