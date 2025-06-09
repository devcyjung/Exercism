#pragma once

#include <string>

namespace pangram {

constexpr unsigned int ALL_FOUND = (1u << 26) - 1;

bool is_pangram(const std::string& input) noexcept;

}  // namespace pangram