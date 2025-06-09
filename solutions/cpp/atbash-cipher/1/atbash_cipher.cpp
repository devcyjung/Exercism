#include "atbash_cipher.h"

namespace atbash_cipher {

    std::string encode(const std::string& input) {
        size_t result_len{0};
        std::string result{};

        for (auto&& ch: input) {
            if ('0' <= ch && ch <= '9') {
                if (result_len > 0 && result_len % 5 == 0) {
                    result.push_back(' ');
                }
                ++result_len;
                result.push_back(ch);
                continue;
            }
            if ('a' <= ch && ch <= 'z') {
                if (result_len > 0 && result_len % 5 == 0) {
                    result.push_back(' ');
                }
                ++result_len;
                result.push_back('z'-ch+'a');
                continue;
            }
            if ('A' <= ch && ch <= 'Z') {
                if (result_len > 0 && result_len % 5 == 0) {
                    result.push_back(' ');
                }
                ++result_len;
                result.push_back('Z'-ch+'a');
                continue;
            }
        }
        return result;
    }

    std::string decode(const std::string& input) {
        std::string result{};

        for (auto&& ch: input) {
            if ('0' <= ch && ch <= '9') {
                result.push_back(ch);
                continue;
            }
            if ('a' <= ch && ch <= 'z') {
                result.push_back('z'-ch+'a');
                continue;
            }
            if ('A' <= ch && ch <= 'Z') {
                result.push_back('Z'-ch+'a');
                continue;
            }
        }
        return result;
    }

}  // namespace atbash_cipher
