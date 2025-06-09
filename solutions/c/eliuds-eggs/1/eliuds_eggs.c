#include "eliuds_eggs.h"

#include <stdbool.h>
#include <stddef.h>

unsigned int egg_count(const unsigned long long number) {
    static bool init = false;
    static unsigned int popcount[256]; 
    if (!init) {
        popcount[0] = 0;
        for (size_t i = 1; i < 256; ++i) {
            popcount[i] = popcount[i >> 1] + (i & 1);
        }
        init = true;
    }
    return popcount[(number >> 0) & 255]
        + popcount[(number >> 1 * 8) & 255]
        + popcount[(number >> 2 * 8) & 255]
        + popcount[(number >> 3 * 8) & 255]
        + popcount[(number >> 4 * 8) & 255]
        + popcount[(number >> 5 * 8) & 255]
        + popcount[(number >> 6 * 8) & 255]
        + popcount[(number >> 7 * 8) & 255];
}