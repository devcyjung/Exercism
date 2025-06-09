#pragma once

#include <cstdint>

namespace chicken_coop {

    [[nodiscard]]
    inline constexpr
    uint64_t positions_to_quantity(const uint64_t n) noexcept {
        return __builtin_popcount(n);
    }

}  // namespace chicken_coop