#ifndef HAMMING_H_
#define HAMMING_H_

#include <optional>
#include <string_view>

namespace hamming {

std::optional<int> compute_opt(std::string_view strand1, std::string_view strand2) noexcept;
int compute(std::string_view strand1, std::string_view strand2);

}  // namespace hamming

#endif     // HAMMING_H_