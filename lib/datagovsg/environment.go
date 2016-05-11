package datagovsg

type APIInfo struct {
	Status string `json:"status,omitempty"`
}
type Location struct {
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

type DatetimeRange struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Area struct {
	Name          string   `json:"name,omitempty"`
	LabelLocation Location `json:"label_location,omitempty"`
}

type Speed struct {
	High int `json:"high,omitempty"`
	Low  int `json:"low,omitempty"`
}

type RelativeHumidity struct {
	High int `json:"high,omitempty"`
	Low  int `json:"low,omitempty"`
}

type Temperature struct {
	High int `json:"high,omitempty"`
	Low  int `json:"low,omitempty"`
}

type Wind struct {
	Speed     Speed  `json:"speed,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type RegionWeatherForecast struct {
	South   string `json:"south,omitempty"`
	North   string `json:"north,omitempty"`
	East    string `json:"east,omitempty"`
	Central string `json:"central,omitempty"`
	West    string `json:"west,omitempty"`
}
