#pragma once

#include <cctype>

namespace etl {

template <typename K, typename V, template<typename...> typename Map, template<typename...> typename List>
inline constexpr Map<V, K> transform(const Map<K, List<V>>& old_map) noexcept
{
    Map<V, K> result{};
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
