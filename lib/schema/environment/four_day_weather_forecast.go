package environment

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var fourDayWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
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

var fourDayWeatherForecastResultItemObject = graphql.NewObject(graphql.ObjectConfig{
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

var fourDayWeatherForecastResultObject = graphql.NewObject(graphql.ObjectConfig{
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
