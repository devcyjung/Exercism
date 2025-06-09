#include "hexadecimal.h"

namespace hexadecimal {

    int convert(std::string s) {
        char c{};
        int sum{0};
        for (auto it = s.begin(); it < s.end(); ++it) {
            sum *= 16;
            c = *it;
            if (c >= 'a' && c <= 'f') {
                sum += c - 'a' + 10;
            } else if (c >= 'A' && c <= 'F') {
                sum += c - 'A' + 10;
            } else if (c >= '0' && c <= '9') {
                sum += c - '0';
            } else {
                return 0;
            }
        }
        return sum;
    }

}  // namespace hexadecimal
