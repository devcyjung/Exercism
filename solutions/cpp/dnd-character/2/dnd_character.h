#ifndef DND_CHARACTER_H_
#define DND_CHARACTER_H_

namespace dnd_character
{

[[nodiscard]]
int roll(void) noexcept;

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
int ability(void) noexcept;

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