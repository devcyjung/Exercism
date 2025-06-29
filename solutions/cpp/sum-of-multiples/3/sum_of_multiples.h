#ifndef SUM_OF_MULTIPLES_H
#define SUM_OF_MULTIPLES_H

#include <initializer_list>
#include <stdexcept>
#include <type_traits>
#include <unordered_set>

namespace sum_of_multiples {

template<typename T>
T to(const std::initializer_list<T>& divisors, const T& limit, std::true_type) noexcept {
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

template<typename T>
T to(const std::initializer_list<T>&, const T&, std::false_type) {
    throw std::domain_error("Input is not integer type");
}

template<typename T>
T to(const std::initializer_list<T>& divisors, const T& limit) {
    return to(divisors, limit, std::is_integral<T>());
}

}  // namespace sum_of_multiples

#endif    // SUM_OF_MULTIPLES_H