package transport

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/geojson"
)

var taxiAvailabiltyResultObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxiAvailabilityResult",
	Fields: graphql.Fields{
		"taxi_count": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"api_info": &graphql.Field{
			Type: graphql.NewNonNull(common.APIInfoStatusObject),
		},
		"timestamp": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"result": &graphql.Field{
			Type: graphql.NewNonNull(geojson.GeoJSONInterface),
		},
	},
})
