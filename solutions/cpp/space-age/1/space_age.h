#pragma once

namespace space_age {

class space_age
{
private:
    const long long int seconds_;
    const double earth_age = static_cast<double>(seconds_) / earth_year_seconds;
    static const long long int earth_year_seconds = 365.25 * 24 * 60 * 60;
    static constexpr double mercury_year = 0.2408467;
    static constexpr double venus_year = 0.61519726;
    static constexpr double mars_year = 1.8808158;
    static constexpr double jupiter_year = 11.862615;
    static constexpr double saturn_year = 29.447498;
    static constexpr double uranus_year = 84.016846;
    static constexpr double neptune_year = 164.79132;
    
public:
    space_age(long long int seconds): seconds_(seconds) {}
    constexpr long long int seconds(void) const {
        return seconds_;
    }
    constexpr double on_mercury(void) const {
        return earth_age / mercury_year;
    }
    constexpr double on_venus(void) const {
        return earth_age / venus_year;
    }
    constexpr double on_earth(void) const {
        return earth_age;
    }
    constexpr double on_mars(void) const {
        return earth_age / mars_year;
    }
    constexpr double on_jupiter(void) const {
        return earth_age / jupiter_year;
    }
    constexpr double on_saturn(void) const {
        return earth_age / saturn_year;
    }
    constexpr double on_uranus(void) const {
        return earth_age / uranus_year;
    }
    constexpr double on_neptune(void) const {
        return earth_age / neptune_year;
    }
};

}  // namespace space_age
