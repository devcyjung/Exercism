#include "eliuds_eggs.h"

namespace chicken_coop {

    unsigned int positions_to_quantity(unsigned int n) {
        unsigned int r{0};
        while (n != 0) {
            r += n & 1;
            n >>= 1;
        }
        return r;
    };

}  // namespace chicken_coop
