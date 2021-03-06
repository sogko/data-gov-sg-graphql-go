package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/environment"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/transport"
)

var Root graphql.Schema

func init() {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:        "RootQuery",
		Description: "Root queries for Data.gov.sg real-time APIs",
		Fields: graphql.Fields{
			"environment": &graphql.Field{
				Description: "Environment-related APIs",
				Type:        environment.RootObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
			"transport": &graphql.Field{
				Description: "Transport-related APIs",
				Type:        transport.RootObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{}, nil
				},
			},
		},
	})
	var err error
	Root, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		panic(err)
	}
}
