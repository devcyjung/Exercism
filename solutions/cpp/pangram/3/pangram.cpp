#include "pangram.h"
#include <bitset>

namespace pangram {

bool is_pangram(const std::string& input)
{
    std::bitset<26> pangram_map {};
    for (auto ch: input)
    {
        if ('a' <= ch && ch <= 'z')
        {
            pangram_map.set(ch - 'a');
        }
        if ('A' <= ch && ch <= 'Z')
        {
            pangram_map.set(ch - 'A');
        }
    }
    return pangram_map.count() == pangram_map.size();
}

}  // namespace pangram