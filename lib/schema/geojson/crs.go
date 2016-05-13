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

func isTypeOfCRS(ttype string, value interface{}, info graphql.ResolveInfo) bool {
	valueMap, _ := value.(map[string]interface{})
	switch ttype {
	case "link":
		_, hasHref := valueMap["href"]
		return hasHref
	case "name":
		_, hasName := valueMap["name"]
		return hasName
	}
	return false
}

var NamedCRSPropertiesObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONNamedCRSProperties",
	Description: "GeoJSON Named Coordinate Reference System (CRS) Properties",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfCRS("name", value, info)
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
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfCRS("link", value, info)
	},
})
var CRSPropertiesUnion = graphql.NewUnion(graphql.UnionConfig{
	Name:        "GeoJSONCRSProperties",
	Description: "GeoJSON Coordinate Reference System (CRS) Properties",
	Types: []*graphql.Object{
		NamedCRSPropertiesObject,
		LinkedCRSPropertiesObject,
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
