#include "raindrops.h"

namespace raindrops {
    std::string convert(int drops) {
        std::string res{};
        if (drops % 3 == 0) {
            res += "Pling";
        }
        if (drops % 5 == 0) {
            res += "Plang";
        }
        if (drops % 7 == 0) {
            res += "Plong";
        }
        if (res.size() == 0) {
            res += std::to_string(drops);
        }
        return res;
    }
}
