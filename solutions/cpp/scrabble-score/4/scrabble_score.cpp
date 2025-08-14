#include "scrabble_score.h"

#include <cctype>
#include <functional>
#include <numeric>

namespace scrabble_score
{
    unsigned int to_score(const char letter)
    {
        switch (const char lower = static_cast<char>(std::tolower(static_cast<unsigned char>(letter))))
        {
            case 'd':
            case 'g':
                return 2;
            case 'b':
            case 'c':
            case 'm':
            case 'p':
                return 3;
            case 'f':
            case 'h':
            case 'v':
            case 'w':
            case 'y':
                return 4;
            case 'k':
                return 5;
            case 'j':
            case 'x':
                return 8;
            case 'q':
            case 'z':
                return 10;
            default:
                if (lower >= 'a' && lower <= 'z')
                {
                    return 1;
                }
                return 0;
        }
    }

    unsigned int score(const std::string_view str)
    {
        return std::accumulate(std::begin(str), std::end(str), 0,
                               [](const unsigned int score, const char letter) { return score + to_score(letter); });
    }
} // namespace scrabble_score
