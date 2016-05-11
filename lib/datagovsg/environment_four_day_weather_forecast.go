package datagovsg

type FourDayWeatherForecastOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
	Date     string `json:"date,omitempy" url:"date,omitempy"`
}

type FourDayWeatherForecast struct {
	Timestamp        string           `json:"timestamp,omitempty"`
	Wind             Wind             `json:"wind,omitempty"`
	Forecast         string           `json:"forecast,omitempty"`
	RelativeHumidity RelativeHumidity `json:"relative_humidity,omitempty"`
	Date             string           `json:"date,omitempty"`
	Temperature      Temperature      `json:"temperature,omitempty"`
}

type FourDayWeatherForecastResult struct {
	UpdateTimestamp string                   `json:"update_timestamp,omitempty"`
	Timestamp       string                   `json:"timestamp,omitempty"`
	Forecasts       []FourDayWeatherForecast `json:"forecasts,omitempty"`
}

type FourDayWeatherForecastResponse struct {
	APIInfo APIInfo                        `json:"api_info,omitempty"`
	Items   []FourDayWeatherForecastResult `json:"items,omitempty"`
}

func (resp *FourDayWeatherForecastResponse) ToGraphQL() interface{} {
	return resp
}
