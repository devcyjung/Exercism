#pragma once
#include <string>

namespace date_independent {

class clock
{
public:
    static clock at(int hour, int minute) noexcept;
    [[nodiscard]] int hour() const noexcept;
    [[nodiscard]] int minute() const noexcept;
    [[nodiscard]] clock plus(int add) const noexcept;
    explicit operator std::string() const noexcept;
    bool operator==(const clock& other) const noexcept;
    bool operator!=(const clock& other) const noexcept;
private:
    static constexpr int MINUTES_IN_DAY = 24 * 60;
    std::div_t _hm;
    explicit clock(int hour, int minute) noexcept;
    [[nodiscard]] auto tied() const noexcept;
};

}  // namespace date_independent
