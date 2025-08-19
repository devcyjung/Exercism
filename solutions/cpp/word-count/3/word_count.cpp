#include "word_count.h"

#include <algorithm>

namespace {

[[nodiscard]]
constexpr bool iswordchar(const char ch) noexcept
{
    return std::isalnum(static_cast<unsigned char>(ch)) || ch == '\'';
}

[[nodiscard]]
std::string trim_matching(const std::string_view text, char trim_char) noexcept
{
    const auto first {text.find_first_not_of(trim_char)};
    if (first == std::string_view::npos)
    {
        return std::string {};
    }
    const auto last {text.find_last_not_of(trim_char)};
    const auto length {last - first + 1};
    return std::string {text.substr(first, length)};
}

}  // namespace

namespace word_count {

auto words(const std::string_view text) noexcept -> std::map<std::string, int>
{
    std::string buffer {};
    std::map<std::string, int> word_count_map {};
    for (auto it {std::begin(text)}; it <= std::end(text); it = std::next(it))
    {
        if (it == std::end(text) || !iswordchar(*it))
        {
            if (auto word {trim_matching(buffer, '\'')}; !word.empty())
            {
                std::transform(word.begin(), word.end(), word.begin(), [](const unsigned char ch) { return std::tolower(ch); });
                ++word_count_map[word];
            }
            buffer.clear();
        }
        else
        {
            buffer += *it;
        }
    }
    return word_count_map;
}

}  // namespace word_count
