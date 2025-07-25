#ifndef TRIANGLE_H_
#define TRIANGLE_H_

#include <algorithm>
#include <array>
#include <stdexcept>

namespace triangle {

enum class flavor: uint8_t
{
    equilateral, isosceles, scalene,
};

template <typename A, typename B, typename C>
inline constexpr flavor kind(A a, B b, C c)
{
    using T = std::common_type_t<A, B, C>;
    static_assert(std::is_arithmetic<T>::value, "Not an arithmetic type");
    std::array<T, 3> sides = {static_cast<T>(a), static_cast<T>(b), static_cast<T>(c)};
    std::sort(sides.begin(), sides.end());
    if (sides[0] <= 0 || sides[0] + sides[1] <= sides[2])
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

#endif  // TRIANGLE_H_