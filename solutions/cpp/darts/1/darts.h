#if !defined(DARTS_H)
#define DARTS_H

#include <cmath>
namespace darts {
    constexpr unsigned int score(double x, double y) noexcept {
        if (auto radius{std::hypot(x, y)}; radius > 10) {
            return 0;
        } else if (radius > 5) {
            return 1;
        } else if (radius > 1) {
            return 5;
        } else {
            return 10;
        }
    }
}

#endif