#include "collatz_conjecture.h"

int steps(int start) {
    if (start < 1) {
        return ERROR_VALUE;
    }
    int count = 0;
    while (start > 1) {
        ++count;
        if ((start & 1) == 0) {
            start >>= 1;
        } else {
            start += (start << 1) + 1;
        }
    }
    return count;
}