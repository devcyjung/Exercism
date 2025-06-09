#include "hamming.h"
#include <string.h>

int compute(const char *lhs, const char *rhs) {
    if (!lhs || !rhs) {
        return ERROR_NULL_PTR;
    }
    size_t len = strlen(lhs);
    if (len != strlen(rhs)) {
        return ERROR_UNEQUAL_LENGTH;
    }
    int distance = 0;
    for (size_t i = 0; i < len; ++i) {
        if (lhs[i] != rhs[i]) {
            ++distance;
        }
    }
    return distance;
}