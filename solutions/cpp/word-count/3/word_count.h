#pragma once
#include <iomanip>
#include <map>
#include <string>
#include <string_view>

template <typename K, typename V>
std::ostream& operator<<(std::ostream& ostr, const std::map<K, V>& map)
{
    ostr << "{ ";
    auto comma {map.size()};
    for (const auto& p : map)
        ostr << '"' << p.first << "\":" << p.second << (--comma ? ", " : " ");
    return ostr << "}\n";
}

namespace word_count {

[[nodiscard]] auto words(std::string_view text) noexcept -> std::map<std::string, int>;

}  // namespace word_count
