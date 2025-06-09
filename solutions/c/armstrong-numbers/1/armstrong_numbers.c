#include "armstrong_numbers.h"

#include <stddef.h>
#include <math.h>

bool is_armstrong_number(const int candidate) {
    int acc = candidate;
    int digits[64];
    size_t len = 0;
    while (acc > 0) {
        digits[len++] = acc % 10;
        acc /= 10;
    }
    for (size_t i = 0; i < len; ++i) {
        acc += pow(digits[i], len);
    }
    return candidate == acc;
}