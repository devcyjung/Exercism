#include "trinary.h"

namespace trinary {

    int to_decimal(std::string s) {
        int sum {0};
        for (auto it = s.begin(); it != s.end(); ++it) {
            sum *= 3;
            switch (*it - '0') {
                case 0:
                    break;
                case 1:
                    ++sum;
                    break;
                case 2:
                    sum += 2;
                    break;
                default:
                    return 0;
            }
        }
        return sum;
    }

}  // namespace trinary
