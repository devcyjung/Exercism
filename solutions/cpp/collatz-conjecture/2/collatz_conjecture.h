#ifndef COLLATZ_CONJECTURE_H_
#define COLLATZ_CONJECTURE_H_

#include <stdexcept>

namespace collatz_conjecture {

inline constexpr int64_t steps(int64_t number) {
    if (number < 1) {
        throw std::domain_error("Input is less than 1");
    }
    long long step_count = 0;
    while (number > 1) {
        ++step_count;
        switch (number & 1) {
        case 0:
            number >>= 1;
            break;
        default:
            number += (number << 1) + 1;
        }
    }
    return step_count;
}

}  // namespace collatz_conjecture

#endif   // COLLATZ_CONJECTURE_H_