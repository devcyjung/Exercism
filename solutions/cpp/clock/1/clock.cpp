#include "clock.h"
#include <iomanip>
#include <sstream>

namespace date_independent {

clock::clock(const int hour, const int minute) : _hm(std::div(((hour * 60 + minute) % (MINUTES_IN_DAY) + MINUTES_IN_DAY) % (MINUTES_IN_DAY), 60))
{}

clock clock::at(const int hour, const int minute)
{
    return clock(hour, minute);
}

clock clock::plus(const int add) const
{
    return clock(_hm.quot, _hm.rem + add);
}

clock::operator std::string() const
{
    std::stringstream stream;
    stream
        << std::setw(2) << std::setfill('0') << _hm.quot
        << ":"
        << std::setw(2) << std::setfill('0') << _hm.rem;
    return stream.str();
}

bool clock::operator==(const clock& other) const
{
    return _hm.quot == other._hm.quot && _hm.rem == other._hm.rem;
}

bool clock::operator!=(const clock& other) const
{
    return !(*this == other);
}

}  // namespace date_independent
