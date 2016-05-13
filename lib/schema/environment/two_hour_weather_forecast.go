package environment

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var twoHourWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
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

var twoHourWeatherForecastResultItemObject = graphql.NewObject(graphql.ObjectConfig{
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

var twoHourWeatherForecastResultObject = graphql.NewObject(graphql.ObjectConfig{
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
