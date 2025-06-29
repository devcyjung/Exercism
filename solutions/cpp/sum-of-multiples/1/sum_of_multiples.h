#pragma once

#include <initializer_list>
#include <stdexcept>
#include <type_traits>
#include <unordered_set>

namespace sum_of_multiples {

template<typename T>
constexpr T to(const std::initializer_list<T>& divisors, const T& limit, std::true_type) {
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
constexpr auto to(const std::initializer_list<T>&, const T&, std::false_type) {
    throw std::domain_error("Input is not integer type");
}

template<typename T>
constexpr auto to(const std::initializer_list<T>& divisors, const T& limit) {
    return to(divisors, limit, std::is_integral<T>());
}

}  // namespace sum_of_multiples
