package transport

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var trafficImageMetadataObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "TrafficImageMetadata",
	Fields: graphql.Fields{
		"height": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"width": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"md5": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var trafficImageCameraObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "TrafficImageCamera",
	Fields: graphql.Fields{
		"timestamp": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"image": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"location": &graphql.Field{
			Type: graphql.NewNonNull(common.LocationObject),
		},
		"camera_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"image_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"image_metadata": &graphql.Field{
			Type: graphql.NewNonNull(trafficImageMetadataObject),
		},
	},
})

var trafficImagesResultItemObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "TrafficImagesResultItem",
	Fields: graphql.Fields{
		"timestamp": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"cameras": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(trafficImageCameraObject))),
		},
	},
})

var trafficImagesResultObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "TrafficImagesResult",
	Fields: graphql.Fields{
		"api_info": &graphql.Field{
			Type: graphql.NewNonNull(common.APIInfoStatusObject),
		},
		"items": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(trafficImagesResultItemObject))),
		},
	},
})
