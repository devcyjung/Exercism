#include "pangram.h"
#include <array>
#include <cctype>

namespace pangram {

bool is_pangram(const std::string& input)
{
    std::array<bool, 26> pangram_map = { 0 };
    for (auto it: input)
    {
        if (std::isupper(it))
        {
            pangram_map[it - 'A'] = true;
        }
        if (std::islower(it))
        {
            pangram_map[it - 'a'] = true;
        }
    }
    for (auto it: pangram_map)
    {
        if (!it)
        {
            return false;
        }
    }
    return true;
}

}  // namespace pangram