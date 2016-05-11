package datagovsg

type TwentyFourHourWeatherForecastOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
	Date     string `json:"date,omitempy" url:"date,omitempy"`
}

type TwentyFourHourWeatherForecast struct {
	Time    DatetimeRange         `json:"time,omitempty"`
	Regions RegionWeatherForecast `json:"regions,omitempty"`
}

type GeneralTwentyFourHourWeatherForecast struct {
	Forecast         string           `json:"forecast,omitempty"`
	RelativeHumidity RelativeHumidity `json:"relative_humidity,omitempty"`
	Temperature      Temperature      `json:"temperature,omitempty"`
	Wind             Wind             `json:"wind,omitempty"`
}

type TwentyFourHourWeatherForecastResult struct {
	UpdateTimestamp string                               `json:"update_timestamp,omitempty"`
	Timestamp       string                               `json:"timestamp,omitempty"`
	ValidPeriod     DatetimeRange                        `json:"valid_period,omitempty"`
	General         GeneralTwentyFourHourWeatherForecast `json:"general,omitempty"`
	Periods         []TwentyFourHourWeatherForecast      `json:"periods,omitempty"`
}

type TwentyFourHourWeatherForecastResponse struct {
	APIInfo APIInfo                               `json:"api_info,omitempty"`
	Items   []TwentyFourHourWeatherForecastResult `json:"items,omitempty"`
}

func (resp *TwentyFourHourWeatherForecastResponse) ToGraphQL() interface{} {
	return resp
}
