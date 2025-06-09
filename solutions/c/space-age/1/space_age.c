#include "space_age.h"

static const int64_t EARTH_YEAR_IN_SECONDS = 365.25 * 24 * 60 * 60;
static const float PLANET_YEARS[8] = {
    0.2408467, 0.61519726, 1.0, 1.8808158, 11.862615, 29.447498, 84.016846, 164.79132
};

float age(planet_t planet, int64_t seconds) {
    if (planet < 0 || planet > 7) {
        return INVALID_INPUT;
    }
    return (float) seconds / (EARTH_YEAR_IN_SECONDS * PLANET_YEARS[planet]);
}