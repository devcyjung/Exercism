package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = iota
	Fahrenheit
)

func (t TemperatureUnit) String() string {
    switch t {
    case Celsius:
        return "°C"
    case Fahrenheit:
        return "°F"
    default:
        return "Invalid TemperatureUnit"
    }
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
    return fmt.Sprint(t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (s SpeedUnit) String() string {
    switch s {
    case KmPerHour:
        return "km/h"
    case MilesPerHour:
        return "mph"
    default:
        return "Invalid SpeedUnit"
    }
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
    return fmt.Sprint(s.magnitude, s.unit)
}

type Humidity int

func (h Humidity) String() string {
    return fmt.Sprint(int(h), "% Humidity")
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      Humidity
}

func (m MeteorologyData) String() string {
    return fmt.Sprint(
        m.location, ": " ,
        m.temperature, ", ",
        "Wind ", m.windDirection, " at ",
        m.windSpeed, ", ",
        m.humidity,
    )
}