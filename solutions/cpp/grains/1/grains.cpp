#include "grains.h"

namespace grains {

    uint64_t square(uint8_t grid) {
        return (uint64_t)1 << (grid - 1);
    }
    
    uint64_t total() {
        return ~(uint64_t)0;
    }

}  // namespace grains
