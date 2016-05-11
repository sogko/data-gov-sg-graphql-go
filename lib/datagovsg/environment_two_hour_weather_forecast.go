package datagovsg

type TwoHourWeatherForecastOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
	Date     string `json:"date,omitempy" url:"date,omitempy"`
}

type TwoHourWeatherForecast struct {
	Area     string `json:"area,omitempty"`
	Forecast string `json:"forecast,omitempty"`
}

type TwoHourWeatherForecastResult struct {
	UpdateTimestamp string                   `json:"update_timestamp,omitempty"`
	Timestamp       string                   `json:"timestamp,omitempty"`
	ValidPeriod     DatetimeRange            `json:"valid_period,omitempty"`
	Forecasts       []TwoHourWeatherForecast `json:"forecasts,omitempty"`
}

type TwoHourWeatherForecastResponse struct {
	APIInfo      APIInfo                        `json:"api_info,omitempty"`
	AreaMetadata []Area                         `json:"area_metadata,omitempty"`
	Items        []TwoHourWeatherForecastResult `json:"items,omitempty"`
}

func (resp *TwoHourWeatherForecastResponse) AreaByName(name string) Area {
	for _, area := range resp.AreaMetadata {
		if area.Name == name {
			return area
		}
	}
	return Area{}

}
func (resp *TwoHourWeatherForecastResponse) ToGraphQL() interface{} {

	items := []TwoHourWeatherForecastResultItemGraphQL{}
	for _, item := range resp.Items {
		forecasts := []TwoHourWeatherForecastGraphQL{}
		for _, forecast := range item.Forecasts {

			f := TwoHourWeatherForecastGraphQL{
				Forecast: forecast.Forecast,
				Area:     resp.AreaByName(forecast.Area),
			}
			forecasts = append(forecasts, f)
		}

		i := TwoHourWeatherForecastResultItemGraphQL{
			UpdateTimestamp: item.UpdateTimestamp,
			Timestamp:       item.Timestamp,
			ValidPeriod:     item.ValidPeriod,
			Forecasts:       forecasts,
		}
		items = append(items, i)
	}

	return TwoHourWeatherForecastResultGraphQL{
		APIInfo: resp.APIInfo,
		Items:   items,
	}
}

type TwoHourWeatherForecastResultGraphQL struct {
	APIInfo APIInfo                                   `json:"api_info,omitempty"`
	Items   []TwoHourWeatherForecastResultItemGraphQL `json:"items,omitempty"`
}

type TwoHourWeatherForecastResultItemGraphQL struct {
	UpdateTimestamp string                          `json:"update_timestamp,omitempty"`
	Timestamp       string                          `json:"timestamp,omitempty"`
	ValidPeriod     DatetimeRange                   `json:"valid_period,omitempty"`
	Forecasts       []TwoHourWeatherForecastGraphQL `json:"forecasts,omitempty"`
}

type TwoHourWeatherForecastGraphQL struct {
	Area     Area   `json:"area,omitempty"`
	Forecast string `json:"forecast,omitempty"`
}
