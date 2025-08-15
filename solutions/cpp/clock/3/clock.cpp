#include "clock.h"
#include <iomanip>
#include <sstream>

namespace date_independent {

clock::clock(const int hour, const int minute) noexcept
    : _hm(std::div(((hour * 60 + minute) % MINUTES_IN_DAY + MINUTES_IN_DAY) % MINUTES_IN_DAY, 60))
{
}

clock clock::at(const int hour, const int minute) noexcept
{
    return clock(hour, minute);
}

clock clock::plus(const int add) const noexcept
{
    return clock(_hm.quot, _hm.rem + add);
}

clock::operator std::string() const noexcept
{
    std::stringstream stream;
    stream
        << std::setfill('0')
        << std::setw(2) << _hm.quot
        << ":"
        << std::setw(2) << _hm.rem;
    return stream.str();
}

bool clock::operator==(const clock& other) const noexcept
{
    return _hm.quot == other._hm.quot && _hm.rem == other._hm.rem;
}

bool clock::operator!=(const clock& other) const noexcept
{
    return !(*this == other);
}

}  // namespace date_independent
