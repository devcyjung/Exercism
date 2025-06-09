#ifndef NUCLEOTIDE_COUNT_H_
#define NUCLEOTIDE_COUNT_H_

#include <map>
#include <optional>
#include <string>

namespace nucleotide_count {

auto count_opt(const std::string&) noexcept -> std::optional<std::map<char, int>>;

auto count(const std::string&) -> std::map<char, int>;

}  // namespace nucleotide_count

#endif    // NUCLEOTIDE_COUNT_H_