package geojson

import (
	"github.com/graphql-go/graphql"
)

var CRSTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "GeoJSONCRSType",
	Description: "GeoJSON Coordinate Reference System (CRS) Types Enum",
	Values: graphql.EnumValueConfigMap{
		"name": &graphql.EnumValueConfig{
			Value: "name",
		},
		"link": &graphql.EnumValueConfig{
			Value: "link",
		},
	},
})

var NamedCRSPropertiesObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONNamedCRSProperties",
	Description: "GeoJSON Named Coordinate Reference System (CRS) Properties",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
var LinkedCRSPropertiesObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONLinkedCRSProperties",
	Description: "GeoJSON Linked Coordinate Reference System (CRS) Properties",
	Fields: graphql.Fields{
		"href": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var CRSPropertiesUnion = graphql.NewUnion(graphql.UnionConfig{
	Name:        "GeoJSONCRSProperties",
	Description: "GeoJSON Coordinate Reference System (CRS) Properties",
	Types: []*graphql.Object{
		NamedCRSPropertiesObject,
		LinkedCRSPropertiesObject,
	},
	ResolveType: func(value interface{}, info graphql.ResolveInfo) *graphql.Object {
		if valueMap, ok := value.(map[string]interface{}); ok {
			if _, hasHref := valueMap["href"]; hasHref {
				return LinkedCRSPropertiesObject
			}
			if _, hasName := valueMap["name"]; hasName {
				return NamedCRSPropertiesObject
			}
		}
		return nil
	},
})

var CoordinateReferenceSystemObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONCoordinateReferenceSystem",
	Description: "GeoJSON Coordinate Reference System (CRS)",
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(CRSTypeEnum),
		},
		"properties": &graphql.Field{
			Type: graphql.NewNonNull(CRSPropertiesUnion),
		},
	},
})
