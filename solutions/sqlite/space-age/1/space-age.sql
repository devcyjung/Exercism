with coefficient(planet, coeff) as (
    values
        ("Mercury", 31557600 * 0.2408467),
        ("Venus", 31557600 * 0.61519726),
        ("Earth", 31557600 * 1.0),
        ("Mars", 31557600 * 1.8808158),
        ("Jupiter", 31557600 * 11.862615),
        ("Saturn", 31557600 * 29.447498),
        ("Uranus", 31557600 * 84.016846),
        ("Neptune", 31557600 * 164.79132)
)
update "space-age"
set result = round(seconds / coeff, 2)
from coefficient
where "space-age".planet = coefficient.planet;