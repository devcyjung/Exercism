#ifndef HAMMING_H_
#define HAMMING_H_

#include <optional>
#include <stdexcept>
#include <string_view>

namespace hamming {

inline constexpr bool is_valid_char(char ch) noexcept {
    return ch == 'A' || ch == 'C' || ch == 'G' || ch == 'T';
}
    
inline constexpr std::optional<int> safe_compute(
    std::string_view strand1, std::string_view strand2
) noexcept {
    if (strand1.size() != strand2.size()) {
        return {};
    }
    int distance = 0;
    for (size_t i = 0; i < strand1.size(); ++i) {
        char ch1 = static_cast<char>(std::toupper(static_cast<unsigned char>(strand1[i])));
        char ch2 = static_cast<char>(std::toupper(static_cast<unsigned char>(strand2[i])));
        if (!is_valid_char(ch1) || !is_valid_char(ch2)) {
            return {};
        }
        if (ch1 != ch2) {
            ++distance;
        }
    }
    return distance;
}
    
inline constexpr int compute(std::string_view strand1, std::string_view strand2) {
    if (auto result = safe_compute(strand1, strand2); result) {
        return *result;
    }
    throw std::domain_error("Invalid input");
}

}  // namespace hamming

#endif     // HAMMING_H_