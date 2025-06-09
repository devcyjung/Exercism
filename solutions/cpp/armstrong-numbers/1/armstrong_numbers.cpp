#include "armstrong_numbers.h"

#include <cmath>
namespace armstrong_numbers {

    bool is_armstrong_number(unsigned int n) {
        auto t = n;
        unsigned int power{0};
        do {
            t /= 10;
            ++power;
        } while (t != 0);
        auto rem = n;
        unsigned int sum{0};
        while (rem != 0) {
            sum += std::pow(rem % 10, power);
            rem /= 10;
        }
        return sum == n;
    }

}  // namespace armstrong_numbers
