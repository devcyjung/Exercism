#pragma once
#include <iomanip>
#include <iostream>
#include <map>
#include <string>
#include <string_view>

template <typename K, typename V, typename Compare = std::less<K>, typename Allocator = std::allocator<std::pair<const K, V>>>
std::ostream& operator<<(std::ostream& ostr, const std::map<K, V, Compare, Allocator>& map)
{
    ostr << "{ ";
    auto comma {map.size()};
    for (const auto& p : map)
        ostr << '"' << p.first << "\":" << p.second << (--comma ? ", " : " ");
    return ostr << "}\n";
}

namespace word_count {

auto words(std::string_view text) -> std::map<std::string, int>;

}  // namespace word_count
