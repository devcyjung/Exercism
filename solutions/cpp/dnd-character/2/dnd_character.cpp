#include "dnd_character.h"

#include <random>

namespace dnd_character
{
    
int roll(void) noexcept
{
    std::random_device rd{};
    std::uniform_int_distribution<int> dist(1, 6);
    return dist(rd);
}

int ability(void) noexcept
{
    int minimum = roll();
    int dice = 0;
    int sum = 0;
    for (auto i = 0; i < 3; ++i)
    {
        dice = roll();
        if (dice < minimum)
        {
            sum += minimum;
            minimum = dice;
        }
        else
        {
            sum += dice;
        }
    }
    return sum;
}
    
} // namespace dnd_character