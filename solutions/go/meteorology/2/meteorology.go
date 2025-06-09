package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = iota
	Fahrenheit
)

func (t TemperatureUnit) String() string {
    if t == 0 {
        return "°C"
    }
    return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
    return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (s SpeedUnit) String() string {
    if s == 0 {
        return "km/h"
    }
    return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
    return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
    return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}