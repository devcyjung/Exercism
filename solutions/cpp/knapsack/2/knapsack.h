#ifndef KNAPSACK_H_
#define KNAPSACK_H_

#include <iterator>
#include <type_traits>
#include <vector>

namespace knapsack {

template <typename W, typename V>
struct KnapsackItem {
    static_assert(std::is_arithmetic<W>::value, "Weight must be numeric");
    static_assert(std::is_arithmetic<V>::value, "Value must be numeric");
    using weight_type = W;
    using value_type = V;
    W weight;
    V value;
};

using Item = KnapsackItem<int, int>;

template <typename Iterable>
using item_t = typename std::decay<decltype(
    *std::begin(std::declval<Iterable>())
)>::type;

template <typename T, typename = void>
struct is_iterable: std::false_type {};

template <typename T>
struct is_iterable<
    T,
    std::void_t<
        decltype(std::begin(std::declval<T&>())),
        decltype(std::end(std::declval<T&>()))
    >
>: std::true_type {};

template <typename T>
struct is_knapsack_item : std::false_type {};

template <typename W, typename V>
struct is_knapsack_item<KnapsackItem<W, V>> : std::true_type {};

template <typename Iterable,
          typename Item = item_t<Iterable>,
          typename = std::enable_if_t<
              is_iterable<Iterable>::value &&
              is_knapsack_item<Item>::value
          >>
auto maximum_value(
    typename Item::weight_type capacity, const Iterable& items
) -> typename Item::value_type
{
    using Value = typename Item::value_type;
    const auto first = std::begin(items);
    const auto last = std::end(items);
    if (first == last)
    {
        return static_cast<Value>(0);
    }
    const auto item_count = static_cast<size_t>(std::distance(first, last));
    const auto max_weight = static_cast<size_t>(capacity);
    std::vector<std::vector<Value>> memo(
        item_count,
        std::vector<Value>(max_weight+1, static_cast<Value>(0))
    );
    for (auto it = first; it != last; ++it)
    {
        const size_t item_weight = static_cast<size_t>((*it).weight);
        const Value item_value = (*it).value;
        const size_t i = std::distance(first, it);
        for (size_t j = 1; j < max_weight+1; ++j)
        {
            if (i == 0 && item_weight > j)
            {
                memo[i][j] = 0;
            }
            else if (i == 0)
            {
                memo[i][j] = item_value;
            }
            else if (item_weight > j)
            {
                memo[i][j] = memo[i-1][j];
            }
            else
            {
                memo[i][j] = std::max(memo[i-1][j-item_weight]+item_value, memo[i-1][j]);
            }
        }
    }
    return memo[item_count-1][max_weight];
}

}  // namespace knapsack

#endif  // KNAPSACK_H_