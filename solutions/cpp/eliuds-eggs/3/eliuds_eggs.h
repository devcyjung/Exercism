#pragma once

namespace chicken_coop {

    [[nodiscard]]
    inline constexpr
    unsigned int positions_to_quantity(const unsigned int n) noexcept {
        return __builtin_popcount(n);
    }

}  // namespace chicken_coop