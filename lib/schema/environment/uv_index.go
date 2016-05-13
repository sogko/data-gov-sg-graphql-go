package environment

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var uvIndexReadingsResultObject = graphql.NewObject(graphql.ObjectConfig{
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

var uvIndexReadingsResultItemObject = graphql.NewObject(graphql.ObjectConfig{
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

var uvIndexReadingObject = graphql.NewObject(graphql.ObjectConfig{
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
