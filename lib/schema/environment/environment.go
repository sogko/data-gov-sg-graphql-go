package environment

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/datagovsg"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var environmentObject *graphql.Object

func RootObject() *graphql.Object {
	if environmentObject != nil {
		return environmentObject
	}

	// Environment
	environmentObject = graphql.NewObject(graphql.ObjectConfig{
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
	return environmentObject
}
