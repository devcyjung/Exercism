#include "sieve.h"

namespace sieve {

    const std::vector<int> primes(const int n) {
        std::vector<int> result(0);
        if (n < 2) {
            return result;
        }
        result.push_back(2);
        bool is_prime{};
        for (int cur = 3; cur <= n; ++cur) {
            is_prime = true;
            for (auto prime = result.begin(); prime != result.end(); ++prime) {
                if (cur % *prime == 0) {
                    is_prime = false;
                    break;
                }
            }
            if (is_prime) {
                result.push_back(cur);
            }
        }
        return result;
    }

}  // namespace sieve
