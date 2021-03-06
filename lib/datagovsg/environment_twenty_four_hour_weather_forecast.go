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

type TwentyFourHourWeatherForecastResultItem struct {
	UpdateTimestamp string                               `json:"update_timestamp,omitempty"`
	Timestamp       string                               `json:"timestamp,omitempty"`
	ValidPeriod     DatetimeRange                        `json:"valid_period,omitempty"`
	General         GeneralTwentyFourHourWeatherForecast `json:"general,omitempty"`
	Periods         []TwentyFourHourWeatherForecast      `json:"periods,omitempty"`
}

type TwentyFourHourWeatherForecastResult struct {
	APIInfo APIInfo                                   `json:"api_info,omitempty"`
	Items   []TwentyFourHourWeatherForecastResultItem `json:"items,omitempty"`
}

func (resp *TwentyFourHourWeatherForecastResult) ToGraphQL() interface{} {
	return resp
}
