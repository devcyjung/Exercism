#ifndef SUM_OF_MULTIPLES_H
#define SUM_OF_MULTIPLES_H

#include <initializer_list>
#include <stdexcept>
#include <type_traits>
#include <unordered_set>

namespace sum_of_multiples {

template<typename T>
T to(const std::initializer_list<T>& divisors, const T& limit) {
    if constexpr (!std::is_integral_v<T>) {
        throw std::domain_error("Only integral types are supported");
    }
    std::unordered_set<T> set {};
    for (const auto& divisor : divisors) {
        for (T multiple { divisor }; multiple < limit; multiple += divisor) {
            set.insert(multiple);
        }
    }
    T sum {};
    for (const auto& value: set) {
        sum += value;
    }
    return sum;
}

}  // namespace sum_of_multiples

#endif    // SUM_OF_MULTIPLES_H