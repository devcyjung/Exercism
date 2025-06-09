#include "pangram.h"

namespace pangram {
    
bool is_pangram(const std::string& input) noexcept
{
    if (input.size() < 26)
    {
        return false;
    }
    unsigned int pangram_map {};
    for (auto ch: input)
    {
        if ('a' <= ch && ch <= 'z')
        {
            pangram_map |= (1u << (ch - 'a')); 
        }
        else if ('A' <= ch && ch <= 'Z')
        {
            pangram_map |= (1u << (ch - 'A'));
        }
        if (pangram_map == ALL_FOUND)
        {
            return true;
        }
    }
    return pangram_map == ALL_FOUND;
}

}  // namespace pangram