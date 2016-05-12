package transport

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/datagovsg"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

// Transport
var TransportObject *graphql.Object

// Traffic Images
var trafficImagesResultObject *graphql.Object
var trafficImagesResultItemObject *graphql.Object
var trafficImageCameraObject *graphql.Object
var trafficImageMetadataObject *graphql.Object

func init() {

	trafficImageMetadataObject = graphql.NewObject(graphql.ObjectConfig{
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
	trafficImageCameraObject = graphql.NewObject(graphql.ObjectConfig{
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
	trafficImagesResultItemObject = graphql.NewObject(graphql.ObjectConfig{
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
	trafficImagesResultObject = graphql.NewObject(graphql.ObjectConfig{
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

	TransportObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Transport",
		Description: "Transport-related APIs",
		Fields: graphql.Fields{
			"traffic_images": &graphql.Field{
				Name: "Traffic Images",
				Type: graphql.NewNonNull(trafficImagesResultObject),
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

					v, _ := query.Values(datagovsg.TrafficImagesOptions{
						DateTime: dateTime,
					})

					ch := c.Get(
						fmt.Sprintf("https://api.data.gov.sg/v1/transport/traffic-images?%v", v.Encode()),
						&datagovsg.TrafficImagesResult{},
					)
					res := <-ch
					if res.Err != nil {
						return nil, res.Err
					}
					resp, _ := res.Body.(*datagovsg.TrafficImagesResult)
					return resp.ToGraphQL(), nil
				},
			},
		},
	})
}
