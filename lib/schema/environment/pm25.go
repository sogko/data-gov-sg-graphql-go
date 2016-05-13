package environment

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var pm25ReadingObject = graphql.NewObject(graphql.ObjectConfig{
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

var pm25ReadingRegionsObject = graphql.NewObject(graphql.ObjectConfig{
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

var pm25ReadingIntervalsObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "PM25ReadingIntervals",
	Fields: graphql.Fields{
		"pm25_one_hourly": &graphql.Field{
			Type: graphql.NewNonNull(pm25ReadingRegionsObject),
		},
	},
})

var pm25ReadingsResultItemObject = graphql.NewObject(graphql.ObjectConfig{
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

var pm25ReadingsResultObject = graphql.NewObject(graphql.ObjectConfig{
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
