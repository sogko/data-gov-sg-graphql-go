package environment

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var psiReadingObject = graphql.NewObject(graphql.ObjectConfig{
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

var psiReadingRegionsObject = graphql.NewObject(graphql.ObjectConfig{
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
var psiReadingIntervalsObject = graphql.NewObject(graphql.ObjectConfig{
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

var psiReadingsResultItemObject = graphql.NewObject(graphql.ObjectConfig{
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

var psiReadingsResultObject = graphql.NewObject(graphql.ObjectConfig{
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
