#ifndef DND_CHARACTER_H_
#define DND_CHARACTER_H_

#include <random>

namespace dnd_character {

[[nodiscard]]
inline int roll(void) noexcept
{
    std::random_device rd{};
    std::uniform_int_distribution<int> dist(1, 6);
    return dist(rd);
}

[[nodiscard]]
inline constexpr int modifier(int ability) noexcept
{
    int base = ability - 10;
    if (base % 2 >= 0)
    {
        return base / 2;
    }
    else
    {
        return base / 2 - 1;
    }
}

[[nodiscard]]
inline int ability(void) noexcept
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

struct Character
{
    const int strength, dexterity, constitution, intelligence,
        wisdom, charisma, hitpoints;

    [[nodiscard]]
    inline Character(void)
        : strength(ability()), dexterity(ability()), constitution(ability()),
          intelligence(ability()), wisdom(ability()), charisma(ability()),
          hitpoints(modifier(constitution) + 10)
    {
    }
};

}  // namespace dnd_character

#endif // DND_CHARACTER_H_