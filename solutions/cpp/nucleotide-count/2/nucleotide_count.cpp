#include "nucleotide_count.h"

#include <stdexcept>

namespace nucleotide_count {

auto count_opt(const std::string& strand) noexcept -> std::optional<std::map<char, int>> {
    std::map<char, int> counter{{'A', 0}, {'C', 0}, {'G', 0}, {'T', 0}};
    const auto& invalid = counter.end();
    for (const auto& ch: strand) {
        if (const auto& it = counter.find(ch); it != invalid) {
            ++(it->second);
        }
        else {
            return {};
        }
    }
    return counter;
}

auto count(const std::string& strand) -> std::map<char, int> {
    if (auto result = count_opt(strand); result) {
        return std::move(*result);
    } else {
        throw std::invalid_argument("invalid characters");
    }
}

}  // namespace nucleotide_count