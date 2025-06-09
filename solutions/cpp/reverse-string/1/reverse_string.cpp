#include "reverse_string.h"

namespace reverse_string { 
    std::string reverse_string(std::string input) {
        std::string res(input.size(), '\0');
        for (int i = 0, j = input.size()-1; i <= j ; ++i, --j) {
            res[i] = input[j];
            res[j] = input[i];
        }
        return res;
    }
}
