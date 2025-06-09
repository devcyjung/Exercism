#include "eliuds_eggs.h"

namespace chicken_coop {

    unsigned int positions_to_quantity(unsigned int n) noexcept {
        return __builtin_popcount(n);
    }

}  // namespace chicken_coop
