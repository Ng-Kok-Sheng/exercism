package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	unit := t.unit.String()
	return fmt.Sprintf("%d %s", t.degree, unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
	unit := []string{"km/h", "mph"}
	return unit[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	unit := s.unit.String()
	return fmt.Sprintf("%d %s", s.magnitude, unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (md MeteorologyData) String() string {
	return fmt.Sprintf(
		"%s: %s, Wind %s at %s, %d%% Humidity",
		md.location,
		md.temperature.String(),
		md.windDirection,
		md.windSpeed.String(),
		md.humidity,
	)
}
