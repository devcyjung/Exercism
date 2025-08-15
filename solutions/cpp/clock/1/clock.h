#pragma once
#include <cstdlib>
#include <string>

namespace date_independent {

class clock
{
    static constexpr int MINUTES_IN_DAY = 24 * 60;
    std::div_t _hm;
    explicit clock(int hour, int minute);
public:
    static clock at(int hour, int minute);
    [[nodiscard]] clock plus(int add) const;
    explicit operator std::string() const;
    bool operator==(const clock& other) const;
    bool operator!=(const clock& other) const;
};

}  // namespace date_independent
