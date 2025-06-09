#include "space_age.h"

namespace space_age {

space_age::space_age(long long int seconds)
    : seconds_(seconds),
      earth_age_(static_cast<double>(seconds_) / earth_year_seconds) {}

}  // namespace space_age