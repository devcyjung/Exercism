#include "prime_factors.h"

#include <stdexcept>

namespace prime_factors {
    
    std::vector<long long> of(long long n) {
        std::vector<long long> result(0);
        auto rem = n;
        if (rem < 1) {
            return result;
        }
        long long div = 2;
        while (rem != 1 && div <= rem) {
            while (rem % div == 0) {
                rem /= div;
                result.push_back(div);
            }
            ++div;
        }
        if (rem != 1) {
            throw std::runtime_error("n not divided evenly");
        }
        return result;
    }

}  // namespace prime_factors
