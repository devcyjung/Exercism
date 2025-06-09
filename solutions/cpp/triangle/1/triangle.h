#pragma once

#include <algorithm>
#include <array>
#include <cstdint>
#include <numeric>
#include <stdexcept>
#include <type_traits>

namespace triangle {

enum class flavor: uint8_t
{
    equilateral, isosceles, scalene,
};

template <typename A, typename B, typename C>
inline flavor kind(const A&& a, const B&& b, const C&& c)
{
    using T = std::common_type_t<A, B, C>;
    static_assert(std::is_arithmetic<T>::value, "Not an arithmetic type");
    T x = static_cast<T>(a);
    T y = static_cast<T>(b);
    T z = static_cast<T>(c);
    std::array<T, 3> sides = {x, y, z};
    std::sort(sides.begin(), sides.end());
    if (sides[0] <= 0 || sides[2] * 2 >= std::accumulate(sides.begin(), sides.end(), static_cast<T>(0)))
    {
        throw std::domain_error("Not a triangle");
    }
    if (sides[0] == sides[2])
    {
        return flavor::equilateral;
    }
    if (sides[1] == sides[0] || sides[1] == sides[2])
    {
        return flavor::isosceles; 
    }
    return flavor::scalene;
}

}  // namespace triangle