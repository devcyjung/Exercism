#pragma once
#include <string_view>

namespace binary {

constexpr unsigned int convert(std::string_view str) noexcept
{
    unsigned int acc {};
    for (const auto ch : str)
    {
        switch (ch)
        {
            case '0':
                acc <<= 1;
                break;
            case '1':
                acc <<= 1;
                ++acc;
                break;
            default:
                return 0;
        }
    }
    return acc;
}

}  // namespace binary
