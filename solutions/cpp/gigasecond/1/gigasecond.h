#pragma once
#include "boost/date_time/posix_time/posix_time_types.hpp"

namespace gigasecond {

[[nodiscard]]
inline constexpr boost::posix_time::ptime advance(boost::posix_time::ptime now) noexcept
{
    return now + boost::posix_time::seconds(1'000'000'000);
}

}  // namespace gigasecond
