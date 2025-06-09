#include "pangram.h"

namespace pangram {

bool is_pangram(const std::string& input)
{
    unsigned int pangram_map {};
    for (auto ch: input)
    {
        if ('a' <= ch && ch <= 'z')
        {
            pangram_map |= (1u << (ch - 'a')); 
        }
        if ('A' <= ch && ch <= 'Z')
        {
            pangram_map |= (1u << (ch - 'A'));
        }
    }
    return __builtin_popcount(pangram_map) == 26;
}

}  // namespace pangram