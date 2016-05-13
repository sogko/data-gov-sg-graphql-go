package transport

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/datagovsg"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/common"
)

var transportObject *graphql.Object

func RootObject() *graphql.Object {
	if transportObject != nil {
		return transportObject
	}
	transportObject = graphql.NewObject(graphql.ObjectConfig{
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
	return transportObject
}
