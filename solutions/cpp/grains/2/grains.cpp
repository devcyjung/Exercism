#include "grains.h"

namespace grains {

    uint64_t square(uint8_t grid) {
        return 1ULL << (grid - 1);
    }
    
    uint64_t total() {
        return ~0ULL;
    }

}  // namespace grains
