package environment

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var regionWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
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

var generalTwentyFourHourWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
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

var twentyFourHourWeatherForecastObject = graphql.NewObject(graphql.ObjectConfig{
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

var twentyFourHourWeatherForecastResultItemObject = graphql.NewObject(graphql.ObjectConfig{
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

var twentyFourHourWeatherForecastResultObject = graphql.NewObject(graphql.ObjectConfig{
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
