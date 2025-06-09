#ifndef SPACE_AGE_H_
#define SPACE_AGE_H_

namespace space_age {

class space_age
{
private:
    const long long int seconds_;
    const double earth_age_;
    inline static constexpr long long int earth_year_seconds = 365.25 * 24 * 60 * 60;
    inline static constexpr double mercury_year = 0.2408467;
    inline static constexpr double venus_year = 0.61519726;
    inline static constexpr double mars_year = 1.8808158;
    inline static constexpr double jupiter_year = 11.862615;
    inline static constexpr double saturn_year = 29.447498;
    inline static constexpr double uranus_year = 84.016846;
    inline static constexpr double neptune_year = 164.79132;
    
public:
    [[nodiscard]]
    explicit inline constexpr space_age(long long int seconds) noexcept
        : seconds_(seconds),
          earth_age_(static_cast<double>(seconds_) / earth_year_seconds)
    {
    }

    [[nodiscard]]
    inline constexpr long long int seconds(void) const noexcept
    {
        return seconds_;
    }

    [[nodiscard]]
    inline constexpr double on_mercury(void) const noexcept
    {
        return earth_age_ / mercury_year;
    }

    [[nodiscard]]
    inline constexpr double on_venus(void) const noexcept
    {
        return earth_age_ / venus_year;
    }

    [[nodiscard]]
    inline constexpr double on_earth(void) const noexcept
    {
        return earth_age_;
    }

    [[nodiscard]]
    inline constexpr double on_mars(void) const noexcept
    {
        return earth_age_ / mars_year;
    }

    [[nodiscard]]
    inline constexpr double on_jupiter(void) const noexcept
    {
        return earth_age_ / jupiter_year;
    }

    [[nodiscard]]
    inline constexpr double on_saturn(void) const noexcept
    {
        return earth_age_ / saturn_year;
    }

    [[nodiscard]]
    inline constexpr double on_uranus(void) const noexcept
    {
        return earth_age_ / uranus_year;
    }

    [[nodiscard]]
    inline constexpr double on_neptune(void) const noexcept
    {
        return earth_age_ / neptune_year;
    }
};

}  // namespace space_age

#endif // SPACE_AGE_H_