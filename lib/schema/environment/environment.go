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
var regionWeatherForecastObject *graphql.Object

// Two Hour Weather Forecast
var twoHourWeatherForecastResultObject *graphql.Object
var twoHourWeatherForecastResultItemObject *graphql.Object
var twoHourWeatherForecastObject *graphql.Object

func getHTTPClient(p graphql.ResolveParams) *datagovsg.Client {
	if c, ok := p.Context.Value("client").(*datagovsg.Client); ok {
		return c
	}
	return datagovsg.NewClient("")
}

func init() {

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

					c := getHTTPClient(p)

					dateTime, _ := p.Args["date_time"].(string)
					date, _ := p.Args["date"].(string)

					v, _ := query.Values(datagovsg.TwoHourWeatherForecastOptions{
						DateTime: dateTime,
						Date:     date,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/environment/2-hour-weather-forecast?%v", v.Encode()),
						&datagovsg.TwoHourWeatherForecastResponse{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.TwoHourWeatherForecastResponse)
					return resp.ToGraphQL(), nil
				},
			},
		},
	})
}
