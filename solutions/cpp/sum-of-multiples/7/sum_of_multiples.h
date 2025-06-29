#ifndef SUM_OF_MULTIPLES_H
#define SUM_OF_MULTIPLES_H

#include <initializer_list>
#include <unordered_set>

namespace sum_of_multiples {

template<typename T>
T to(const std::initializer_list<T>& divisors, const T& limit) {
    static_assert(std::is_integral_v<T>, "T must be integral");
    std::unordered_set<T> set {};
    for (const auto& divisor : divisors) {
        for (T multiple { divisor }; divisor > 0 && multiple < limit; multiple += divisor) {
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