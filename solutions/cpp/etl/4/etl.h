#pragma once
#include <type_traits>
#include <cctype>

namespace etl {

#if __cplusplus >= 202002L
    // ===== C++20 version with concepts =====
    template <typename T>
    concept CharLike = std::is_convertible_v<T, unsigned char>;

    template <typename Map, typename Key, typename Value>
    concept MapLike = requires(Map m, Key k, Value v) {
        typename Map::key_type;
        typename Map::mapped_type;
        m[k] = v;
        { m.begin() } -> std::same_as<typename Map::iterator>;
    };

    template <typename List, typename Value>
    concept ListLike = requires(List lst, Value v) {
        typename List::value_type;
        { lst.begin() } -> std::same_as<typename List::iterator>;
    };

    template <
        typename K,
        typename V,
        template<typename...> typename Map,
        template<typename...> typename List
    >
    requires CharLike<V> &&
             MapLike<Map<K, List<V>>, K, List<V>> &&
             ListLike<List<V>, V>
    Map<V, K> transform(const Map<K, List<V>>& old_map)
    {
        Map<V, K> result{};
        for (const auto& [key, list] : old_map) {
            for (const auto& value : list) {
                result[std::tolower(static_cast<unsigned char>(value))] = key;
            }
        }
        return result;
    }

#else
    // ===== C++17 fallback with static_assert =====
    template <
        typename K,
        typename V,
        template<typename...> typename Map,
        template<typename...> typename List
    >
    auto transform(const Map<K, List<V>>& old_map)
        -> typename std::enable_if<
            std::is_convertible<V, unsigned char>::value &&
            std::is_class<Map<K, List<V>>>::value &&
            std::is_class<List<V>>::value,
            Map<V, K>
        >::type
    {
        static_assert(std::is_convertible<V, unsigned char>::value,
                      "V must be convertible to unsigned char for std::tolower");
        static_assert(std::is_class<Map<K, List<V>>>::value,
                      "Map<K, List<V>> must be a class type");
        static_assert(std::is_class<List<V>>::value,
                      "List<V> must be a class type");

        Map<V, K> result{};
        for (const auto& [key, list] : old_map) {
            for (const auto& value : list) {
                result[std::tolower(static_cast<unsigned char>(value))] = key;
            }
        }
        return result;
    }
#endif

} // namespace etl
