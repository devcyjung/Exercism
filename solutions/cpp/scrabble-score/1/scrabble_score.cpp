#include "scrabble_score.h"

#include <cctype>

#include "test/catch.hpp"

namespace scrabble_score {

unsigned int score(std::string_view str)
{
    unsigned int score = 0;
    for (const char c : str)
    {
        switch (const char lower = static_cast<char>(std::tolower(static_cast<unsigned char>(c))))
        {
            case 'd':
            case 'g':
                score += 2;
                break;
            case 'b':
            case 'c':
            case 'm':
            case 'p':
                score += 3;
                break;
            case 'f':
            case 'h':
            case 'v':
            case 'w':
            case 'y':
                score += 4;
                break;
            case 'k':
                score += 5;
                break;
            case 'j':
            case 'x':
                score += 8;
                break;
            case 'q':
            case 'z':
                score += 10;
                break;
            default:
                if (lower >= 'a' && lower <= 'z') {
                    score += 1;
                }
        }
    }
    return score;
}

}  // namespace scrabble_score
