package environment

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/datagovsg"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

// Environment
var EnvironmentObject *graphql.Object

// Two Hour Weather Forecast
var twoHourWeatherForecastResultObject *graphql.Object
var twoHourWeatherForecastResultItemObject *graphql.Object
var twoHourWeatherForecastObject *graphql.Object

// Twenty Four Hour Weather Forecast
var regionWeatherForecastObject *graphql.Object
var twentyFourHourWeatherForecastResultObject *graphql.Object
var twentyFourHourWeatherForecastResultItemObject *graphql.Object
var twentyFourHourWeatherForecastObject *graphql.Object
var generalTwentyFourHourWeatherForecastObject *graphql.Object

// Four Day Weather Forecast
var fourDayWeatherForecastResultObject *graphql.Object
var fourDayWeatherForecastResultItemObject *graphql.Object
var fourDayWeatherForecastObject *graphql.Object

// PM25 Readings
var pm25ReadingsResultObject *graphql.Object
var pm25ReadingsResultItemObject *graphql.Object
var pm25ReadingIntervalsObject *graphql.Object
var pm25ReadingRegionsObject *graphql.Object
var pm25ReadingObject *graphql.Object

// PSI Readings
var psiReadingsResultObject *graphql.Object
var psiReadingsResultItemObject *graphql.Object
var psiReadingIntervalsObject *graphql.Object
var psiReadingRegionsObject *graphql.Object
var psiReadingObject *graphql.Object

// UV-Index Readings
var uvIndexReadingsResultObject *graphql.Object
var uvIndexReadingsResultItemObject *graphql.Object
var uvIndexReadingObject *graphql.Object

func init() {

	// Two Hour Weather Forecast
	twoHourWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "TwoHourWeatherForecast",
		Fields: graphql.Fields{
			"area": &graphql.Field{
				Type: graphql.NewNonNull(common.AreaObject),
			},
			"forecast": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	twoHourWeatherForecastResultItemObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "TwoHourWeatherForecastResultItem",
		Fields: graphql.Fields{
			"update_timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"valid_period": &graphql.Field{
				Type: graphql.NewNonNull(common.DateTimeRangeObject),
			},
			"forecasts": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(twoHourWeatherForecastObject))),
			},
		},
	})
	twoHourWeatherForecastResultObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "TwoHourWeatherForecastResult",
		Fields: graphql.Fields{
			"api_info": &graphql.Field{
				Type: graphql.NewNonNull(common.APIInfoStatusObject),
			},
			"items": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(twoHourWeatherForecastResultItemObject))),
			},
		},
	})

	// Twenty Four Hour Weather Forecast
	regionWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "RegionWeatherForecast",
		Fields: graphql.Fields{
			"south": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"north": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"east": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"central": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"west": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	generalTwentyFourHourWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "GeneralTwentyFourHourWeatherForecast",
		Fields: graphql.Fields{
			"forecast": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"relative_humidity": &graphql.Field{
				Type: graphql.NewNonNull(common.RelativeHumidityObject),
			},
			"temperature": &graphql.Field{
				Type: graphql.NewNonNull(common.TemperatureObject),
			},
			"wind": &graphql.Field{
				Type: graphql.NewNonNull(common.WindObject),
			},
		},
	})
	twentyFourHourWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "TwentyFourHourWeatherForecast",
		Fields: graphql.Fields{
			"time": &graphql.Field{
				Type: graphql.NewNonNull(common.DateTimeRangeObject),
			},
			"regions": &graphql.Field{
				Type: graphql.NewNonNull(regionWeatherForecastObject),
			},
		},
	})
	twentyFourHourWeatherForecastResultItemObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "TwentyFourHourWeatherForecastResultItem",
		Fields: graphql.Fields{
			"update_timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"valid_period": &graphql.Field{
				Type: graphql.NewNonNull(common.DateTimeRangeObject),
			},
			"general": &graphql.Field{
				Type: graphql.NewNonNull(generalTwentyFourHourWeatherForecastObject),
			},
			"periods": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(twentyFourHourWeatherForecastObject))),
			},
		},
	})
	twentyFourHourWeatherForecastResultObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "TwentyFourHourWeatherForecastResult",
		Fields: graphql.Fields{
			"api_info": &graphql.Field{
				Type: graphql.NewNonNull(common.APIInfoStatusObject),
			},
			"items": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(twentyFourHourWeatherForecastResultItemObject))),
			},
		},
	})

	// Four Day Weather Forecast
	fourDayWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "FourDayWeatherForecast",
		Fields: graphql.Fields{
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"wind": &graphql.Field{
				Type: graphql.NewNonNull(common.WindObject),
			},
			"forecast": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"relative_humidity": &graphql.Field{
				Type: graphql.NewNonNull(common.RelativeHumidityObject),
			},
			"date": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"temperature": &graphql.Field{
				Type: graphql.NewNonNull(common.TemperatureObject),
			},
		},
	})
	fourDayWeatherForecastResultItemObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "FourDayWeatherForecastResultItem",
		Fields: graphql.Fields{
			"update_timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"forecasts": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(fourDayWeatherForecastObject))),
			},
		},
	})
	fourDayWeatherForecastResultObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "FourDayWeatherForecastResult",
		Fields: graphql.Fields{
			"api_info": &graphql.Field{
				Type: graphql.NewNonNull(common.APIInfoStatusObject),
			},
			"items": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(fourDayWeatherForecastResultItemObject))),
			},
		},
	})

	// PM25 Readings
	pm25ReadingObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PM25Reading",
		Fields: graphql.Fields{
			"value": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"area": &graphql.Field{
				Type: graphql.NewNonNull(common.AreaObject),
			},
		},
	})
	pm25ReadingRegionsObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PM25ReadingRegions",
		Fields: graphql.Fields{
			"south": &graphql.Field{
				Type: graphql.NewNonNull(pm25ReadingObject),
			},
			"north": &graphql.Field{
				Type: graphql.NewNonNull(pm25ReadingObject),
			},
			"east": &graphql.Field{
				Type: graphql.NewNonNull(pm25ReadingObject),
			},
			"central": &graphql.Field{
				Type: graphql.NewNonNull(pm25ReadingObject),
			},
			"west": &graphql.Field{
				Type: graphql.NewNonNull(pm25ReadingObject),
			},
		},
	})
	pm25ReadingIntervalsObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PM25ReadingIntervals",
		Fields: graphql.Fields{
			"pm25_one_hourly": &graphql.Field{
				Type: graphql.NewNonNull(pm25ReadingRegionsObject),
			},
		},
	})
	pm25ReadingsResultItemObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PM25ReadingsResultItem",
		Fields: graphql.Fields{
			"update_timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"readings": &graphql.Field{
				Type: graphql.NewNonNull(pm25ReadingIntervalsObject),
			},
		},
	})
	pm25ReadingsResultObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PM25ReadingsResult",
		Fields: graphql.Fields{
			"api_info": &graphql.Field{
				Type: graphql.NewNonNull(common.APIInfoStatusObject),
			},
			"items": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(pm25ReadingsResultItemObject))),
			},
		},
	})

	// PSI Readings
	psiReadingObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PSIReading",
		Fields: graphql.Fields{
			"value": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"area": &graphql.Field{
				Type: graphql.NewNonNull(common.AreaObject),
			},
		},
	})
	psiReadingRegionsObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PSIReadingRegions",
		Fields: graphql.Fields{
			"national": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingObject),
			},
			"south": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingObject),
			},
			"north": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingObject),
			},
			"east": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingObject),
			},
			"central": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingObject),
			},
			"west": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingObject),
			},
		},
	})
	psiReadingIntervalsObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PSIReadingIntervals",
		Fields: graphql.Fields{
			"psi_twenty_four_hourly": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"pm10_twenty_four_hourly": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"pm10_sub_index": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"pm25_twenty_four_hourly": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"psi_three_hourly": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"o2_twenty_four_hourly": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"o3_sub_index": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"no2_one_hour_max": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"so2_sub_index": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"pm25_sub_index": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"co_eight_hour_max": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"co_sub_index": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
			"o3_eight_hour_max": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingRegionsObject),
			},
		},
	})
	psiReadingsResultItemObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PSIReadingsResultItem",
		Fields: graphql.Fields{
			"update_timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"readings": &graphql.Field{
				Type: graphql.NewNonNull(psiReadingIntervalsObject),
			},
		},
	})
	psiReadingsResultObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "PSIReadingsResult",
		Fields: graphql.Fields{
			"api_info": &graphql.Field{
				Type: graphql.NewNonNull(common.APIInfoStatusObject),
			},
			"items": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(psiReadingsResultItemObject))),
			},
		},
	})

	// UV-Index Readings
	uvIndexReadingObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "UVIndexReading",
		Fields: graphql.Fields{
			"value": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	uvIndexReadingsResultItemObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "UVIndexReadingsResultItem",
		Fields: graphql.Fields{
			"update_timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"timestamp": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"index": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(uvIndexReadingObject))),
			},
		},
	})
	uvIndexReadingsResultObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "UVIndexReadingsResult",
		Fields: graphql.Fields{
			"api_info": &graphql.Field{
				Type: graphql.NewNonNull(common.APIInfoStatusObject),
			},
			"items": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(uvIndexReadingsResultItemObject))),
			},
		},
	})

	// Environment
	EnvironmentObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Environment",
		Description: "Environment-related APIs",
		Fields: graphql.Fields{
			"two_hour_weather_forecast": &graphql.Field{
				Name: "Two Hour Weather Forecast",
				Type: graphql.NewNonNull(twoHourWeatherForecastResultObject),
				Args: graphql.FieldConfigArgument{
					"date_time": &graphql.ArgumentConfig{
						Type: common.DateTimeStringScalar,
					},
					"date": &graphql.ArgumentConfig{
						Type: common.DateStringScalar,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					c := datagovsg.GetClientFromContext(p.Context)

					dateTime, _ := p.Args["date_time"].(string)
					date, _ := p.Args["date"].(string)

					v, _ := query.Values(datagovsg.TwoHourWeatherForecastOptions{
						DateTime: dateTime,
						Date:     date,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/environment/2-hour-weather-forecast?%v", v.Encode()),
						&datagovsg.TwoHourWeatherForecastResult{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.TwoHourWeatherForecastResult)
					return resp.ToGraphQL(), nil
				},
			},
			"twenty_four_hour_weather_forecast": &graphql.Field{
				Name: "Twenty-Four Hour Weather Forecast",
				Type: graphql.NewNonNull(twentyFourHourWeatherForecastResultObject),
				Args: graphql.FieldConfigArgument{
					"date_time": &graphql.ArgumentConfig{
						Type: common.DateTimeStringScalar,
					},
					"date": &graphql.ArgumentConfig{
						Type: common.DateStringScalar,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					c := datagovsg.GetClientFromContext(p.Context)

					dateTime, _ := p.Args["date_time"].(string)
					date, _ := p.Args["date"].(string)

					v, _ := query.Values(datagovsg.TwentyFourHourWeatherForecastOptions{
						DateTime: dateTime,
						Date:     date,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/environment/24-hour-weather-forecast?%v", v.Encode()),
						&datagovsg.TwentyFourHourWeatherForecastResult{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.TwentyFourHourWeatherForecastResult)
					return resp.ToGraphQL(), nil
				},
			},
			"four_day_weather_forecast": &graphql.Field{
				Name: "Four Day Weather Forecast",
				Type: graphql.NewNonNull(fourDayWeatherForecastResultObject),
				Args: graphql.FieldConfigArgument{
					"date_time": &graphql.ArgumentConfig{
						Type: common.DateTimeStringScalar,
					},
					"date": &graphql.ArgumentConfig{
						Type: common.DateStringScalar,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					c := datagovsg.GetClientFromContext(p.Context)

					dateTime, _ := p.Args["date_time"].(string)
					date, _ := p.Args["date"].(string)

					v, _ := query.Values(datagovsg.FourDayWeatherForecastOptions{
						DateTime: dateTime,
						Date:     date,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/environment/4-day-weather-forecast?%v", v.Encode()),
						&datagovsg.FourDayWeatherForecastResult{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.FourDayWeatherForecastResult)
					return resp.ToGraphQL(), nil
				},
			},
			"pm25": &graphql.Field{
				Name: "PM25 Readings",
				Type: graphql.NewNonNull(pm25ReadingsResultObject),
				Args: graphql.FieldConfigArgument{
					"date_time": &graphql.ArgumentConfig{
						Type: common.DateTimeStringScalar,
					},
					"date": &graphql.ArgumentConfig{
						Type: common.DateStringScalar,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					c := datagovsg.GetClientFromContext(p.Context)

					dateTime, _ := p.Args["date_time"].(string)
					date, _ := p.Args["date"].(string)

					v, _ := query.Values(datagovsg.PM25ReadingsOptions{
						DateTime: dateTime,
						Date:     date,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/environment/pm25?%v", v.Encode()),
						&datagovsg.PM25ReadingsResult{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.PM25ReadingsResult)
					return resp.ToGraphQL(), nil
				},
			},
			"psi": &graphql.Field{
				Name: "PSI Readings",
				Type: graphql.NewNonNull(psiReadingsResultObject),
				Args: graphql.FieldConfigArgument{
					"date_time": &graphql.ArgumentConfig{
						Type: common.DateTimeStringScalar,
					},
					"date": &graphql.ArgumentConfig{
						Type: common.DateStringScalar,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					c := datagovsg.GetClientFromContext(p.Context)

					dateTime, _ := p.Args["date_time"].(string)
					date, _ := p.Args["date"].(string)

					v, _ := query.Values(datagovsg.PSIReadingsOptions{
						DateTime: dateTime,
						Date:     date,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/environment/psi?%v", v.Encode()),
						&datagovsg.PSIReadingsResult{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.PSIReadingsResult)
					return resp.ToGraphQL(), nil
				},
			},
			"uv_index": &graphql.Field{
				Name: "PSI Readings",
				Type: graphql.NewNonNull(uvIndexReadingsResultObject),
				Args: graphql.FieldConfigArgument{
					"date_time": &graphql.ArgumentConfig{
						Type: common.DateTimeStringScalar,
					},
					"date": &graphql.ArgumentConfig{
						Type: common.DateStringScalar,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					c := datagovsg.GetClientFromContext(p.Context)

					dateTime, _ := p.Args["date_time"].(string)
					date, _ := p.Args["date"].(string)

					v, _ := query.Values(datagovsg.UVIndexOptions{
						DateTime: dateTime,
						Date:     date,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/environment/uv-index?%v", v.Encode()),
						&datagovsg.UVIndexReadingsResult{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.UVIndexReadingsResult)
					return resp.ToGraphQL(), nil
				},
			},
		},
	})
}
