#include "clock.h"

#include <iomanip>
#include <tuple>

namespace {

constexpr int euclid_rem(int dividend, int divisor) noexcept
{
    return (dividend % divisor + divisor) % divisor;
}

}  // namespace

namespace date_independent {

clock::clock(const int hour, const int minute) noexcept
    : _hm(std::div(euclid_rem(hour * 60 + minute, MINUTES_IN_DAY), 60)) {}

auto clock::tied() const noexcept { return std::tie(_hm.quot, _hm.rem); }

int clock::hour() const noexcept { return _hm.quot; }

int clock::minute() const noexcept { return _hm.rem; }

clock clock::at(const int hour, const int minute) noexcept { return clock(hour, minute); }

clock clock::plus(const int add) const noexcept { return clock(hour(), minute() + add); }

clock::operator std::string() const noexcept
{
    std::stringstream stream {};
    stream
        << std::setfill('0')
        << std::setw(2) << hour()
        << ':'
        << std::setw(2) << minute();
    return stream.str();
}

bool clock::operator==(const clock& other) const noexcept { return tied() == other.tied(); }

bool clock::operator!=(const clock& other) const noexcept { return !(*this == other); }

}  // namespace date_independent