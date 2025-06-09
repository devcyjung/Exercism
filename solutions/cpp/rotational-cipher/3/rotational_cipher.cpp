#include "rotational_cipher.h"

#include <cctype>
namespace rotational_cipher {

    std::string rotate(const std::string& input, int n) {
        n %= 26;
        if (n < 0) {
            n += 26;
        }
        std::string result{};
        for (auto&& ch: input) {
            if (std::isupper(ch)) {
                result.push_back(std::isupper(ch + n) ? ch + n : ch + n - 'Z' - 1 + 'A');
            } else if (std::islower(ch)) {
                result.push_back(std::islower(ch + n) ? ch + n : ch + n - 'z' - 1 + 'a');
            } else {
                result.push_back(ch);
            }
        }
        return result;
    }

}  // namespace rotational_cipher
