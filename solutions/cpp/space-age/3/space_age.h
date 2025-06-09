#pragma once

namespace space_age {

class space_age
{
private:
    const long long int seconds_;
    const double earth_age_;
    inline static const long long int earth_year_seconds = 365.25 * 24 * 60 * 60;
    inline static constexpr double mercury_year = 0.2408467;
    inline static constexpr double venus_year = 0.61519726;
    inline static constexpr double mars_year = 1.8808158;
    inline static constexpr double jupiter_year = 11.862615;
    inline static constexpr double saturn_year = 29.447498;
    inline static constexpr double uranus_year = 84.016846;
    inline static constexpr double neptune_year = 164.79132;
    
public:
    space_age(long long int seconds);
    inline constexpr long long int seconds(void) const noexcept {
        return seconds_;
    }
    inline constexpr double on_mercury(void) const noexcept {
        return earth_age_ / mercury_year;
    }
    inline constexpr double on_venus(void) const noexcept {
        return earth_age_ / venus_year;
    }
    inline constexpr double on_earth(void) const noexcept {
        return earth_age_;
    }
    inline constexpr double on_mars(void) const noexcept {
        return earth_age_ / mars_year;
    }
    inline constexpr double on_jupiter(void) const noexcept {
        return earth_age_ / jupiter_year;
    }
    inline constexpr double on_saturn(void) const noexcept {
        return earth_age_ / saturn_year;
    }
    inline constexpr double on_uranus(void) const noexcept {
        return earth_age_ / uranus_year;
    }
    inline constexpr double on_neptune(void) const noexcept {
        return earth_age_ / neptune_year;
    }
};

}  // namespace space_age