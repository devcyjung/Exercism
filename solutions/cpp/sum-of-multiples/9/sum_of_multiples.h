#ifndef SUM_OF_MULTIPLES_H
#define SUM_OF_MULTIPLES_H

#include <cmath>
#include <initializer_list>
#include <numeric>
#include <unordered_set>

namespace sum_of_multiples {

template<typename T>
T to(const std::initializer_list<T>& divisors, const T& limit) {
    static_assert(std::is_integral_v<T>, "T must be integral");
    std::unordered_set<T> set {};
    const bool negativeness = limit < 0;
    for (const auto divisor : divisors) {
        if (divisor == 0 || negativeness != (divisor < 0)) continue;
        for (
            T multiple {divisor};
            negativeness ? (multiple > limit) : (multiple < limit);
            multiple += divisor
        ) {
            set.insert(multiple);
        }
    }
    return std::reduce(set.begin(), set.end(), T{});
}

}  // namespace sum_of_multiples

#endif    // SUM_OF_MULTIPLES_H