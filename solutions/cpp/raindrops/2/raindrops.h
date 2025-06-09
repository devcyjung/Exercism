#ifndef RAINDROPS_H_
#define RAINDROPS_H_

#include <array>
#include <string>
#include <string_view>
#include <utility>

namespace raindrops {

inline constexpr std::array<std::pair<int, std::string_view>, 3> sound_map = {{
    {3, "Pling"},
    {5, "Plang"},
    {7, "Plong"},
}};
    
[[nodiscard]]    
inline std::string convert(int drops) noexcept
{
    std::string sound_str{};
    sound_str.reserve(15);
    for (auto [divisor, sound]: sound_map)
    {
        if (drops % divisor == 0)
        {
            sound_str += sound;
        }
    }
    if (sound_str.empty())
    {
        sound_str += std::to_string(drops);
    }
    sound_str.shrink_to_fit();
    return sound_str;
}
    
} // namespace raindrops

#endif // RAINDROPS_H_