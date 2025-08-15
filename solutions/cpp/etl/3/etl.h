#pragma once

#include <cctype>
#include <map>

namespace etl {

template <typename Map>
inline constexpr auto transform(const Map& old_map) noexcept
{
    using Key = typename Map::key_type;
    using List = typename Map::mapped_type;
    using Value = typename List::value_type;
    std::map<Value, Key> result{};
    for (const auto& [key, list] : old_map)
    {
        for (const auto& value : list)
        {
            result[std::tolower(static_cast<unsigned char>(value))] = key;
        }
    }
    return result;
}

}  // namespace etl
