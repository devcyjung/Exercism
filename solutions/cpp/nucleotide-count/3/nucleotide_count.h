#ifndef NUCLEOTIDE_COUNT_H_
#define NUCLEOTIDE_COUNT_H_

#include <map>
#include <optional>
#include <string_view>

namespace nucleotide_count {

auto count_opt(std::string_view) noexcept -> std::optional<std::map<char, int>>;

auto count(std::string_view) -> std::map<char, int>;

}  // namespace nucleotide_count

#endif    // NUCLEOTIDE_COUNT_H_