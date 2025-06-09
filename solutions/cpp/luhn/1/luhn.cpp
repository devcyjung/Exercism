#include "luhn.h"

namespace luhn {

    bool valid(std::string s) {
        int count{0};
        char ch{};
        int sum{0};
        int t{};
        for (auto it = s.rbegin(); it != s.rend(); ++it) {
            ch = *it;
            if ('0' <= ch && ch <= '9') {
                if (count % 2 == 1) {
                    t = 2 * (ch - '0');
                    sum += t;
                    if (t > 9) {
                        sum -= 9;
                    }
                } else {
                    sum += (ch - '0');
                }
                ++count;
            } else if (ch != ' ' && ch != '\n' && ch != 'r' && ch != '\t') {
                return false;
            }
        }
        if (count <= 1) {
            return false;
        }
        return sum % 10 == 0;
    }

}  // namespace luhn
