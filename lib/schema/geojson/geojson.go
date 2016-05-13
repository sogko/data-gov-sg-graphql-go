package geojson

import (
	"github.com/graphql-go/graphql"
)

var TypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "GeoJSONType",
	Description: "GeoJSON Types Enum",
	Values: graphql.EnumValueConfigMap{
		"Point": &graphql.EnumValueConfig{
			Value: "Point",
		},
		"MultiPoint": &graphql.EnumValueConfig{
			Value: "MultiPoint",
		},
		"LineString": &graphql.EnumValueConfig{
			Value: "LineString",
		},
		"MultiLineString": &graphql.EnumValueConfig{
			Value: "MultiLineString",
		},
		"Polygon": &graphql.EnumValueConfig{
			Value: "Polygon",
		},
		"MultiPolygon": &graphql.EnumValueConfig{
			Value: "MultiPolygon",
		},
		"GeometryCollection": &graphql.EnumValueConfig{
			Value: "GeometryCollection",
		},
		"Feature": &graphql.EnumValueConfig{
			Value: "Feature",
		},
		"FeatureCollection": &graphql.EnumValueConfig{
			Value: "FeatureCollection",
		},
	},
})

var GeoJSONInterface = graphql.NewInterface(graphql.InterfaceConfig{
	Name:        "GeoJSONInterface",
	Description: "GeoJSON Interface",
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
	},
})
var GeometryInterface = graphql.NewInterface(graphql.InterfaceConfig{
	Name:        "GeoJSONGeometryInterface",
	Description: "GeoJSON Geometry Interface",
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
	},
})

func isTypeOfGeoJSON(geoJSONType string, value interface{}, info graphql.ResolveInfo) bool {

	if value, ok := value.(map[string]interface{}); ok {
		ttype, _ := value["type"]
		return ttype == geoJSONType
	}
	return false
}

var PointObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONPoint",
	Description: "GeoJSON Point Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
		GeometryInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("Point", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"coordinates": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
	},
})

var MultiPointObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONMultiPoint",
	Description: "GeoJSON MultiPoint Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
		GeometryInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("MultiPoint", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"coordinates": &graphql.Field{
			Type: graphql.NewList(
				graphql.NewList(graphql.Float),
			),
		},
	},
})

var LineStringObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONLineString",
	Description: "GeoJSON LineString Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
		GeometryInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("LineString", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"coordinates": &graphql.Field{
			Type: graphql.NewList(
				graphql.NewList(graphql.Float),
			),
		},
	},
})

var MultiLineStringObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONMultiLineString",
	Description: "GeoJSON MultiLineString Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
		GeometryInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("MultiLineString", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"coordinates": &graphql.Field{
			Type: graphql.NewList(
				graphql.NewList(
					graphql.NewList(graphql.Float),
				),
			),
		},
	},
})

var PolygonObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONPolygon",
	Description: "GeoJSON Polygon Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
		GeometryInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("Polygon", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"coordinates": &graphql.Field{
			Type: graphql.NewList(
				graphql.NewList(
					graphql.NewList(graphql.Float),
				),
			),
		},
	},
})

var MultiPolygonObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONMultiPolygon",
	Description: "GeoJSON MultiPolygon Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
		GeometryInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("MultiPolygon", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"coordinates": &graphql.Field{
			Type: graphql.NewList(
				graphql.NewList(
					graphql.NewList(
						graphql.NewList(graphql.Float),
					),
				),
			),
		},
	},
})

var GeometryCollectionObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONGeometryCollection",
	Description: "GeoJSON GeometryCollection Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("GeometryCollection", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"geometries": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(
				graphql.NewNonNull(GeometryInterface),
			)),
		},
	},
})

var FeatureObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONFeature",
	Description: "GeoJSON Feature Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("Feature", value, info)
	},
	Fields: (graphql.FieldsThunk)(func() graphql.Fields {
		return graphql.Fields{
			"type": &graphql.Field{
				Type: graphql.NewNonNull(TypeEnum),
			},
			"crs": &graphql.Field{
				Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
			},
			"bbox": &graphql.Field{
				Type: graphql.NewList(graphql.Float),
			},
			"geometry": &graphql.Field{
				Type: GeoJSONInterface,
			},
			// Interim definition
			"properties": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.String,
			},
		}
	}),
})

var FeatureCollectionObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoJSONFeatureCollection",
	Description: "GeoJSON FeatureCollection Object",
	Interfaces: []*graphql.Interface{
		GeoJSONInterface,
	},
	IsTypeOf: func(value interface{}, info graphql.ResolveInfo) bool {
		return isTypeOfGeoJSON("FeatureCollection", value, info)
	},
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.NewNonNull(TypeEnum),
		},
		"crs": &graphql.Field{
			Type: graphql.NewNonNull(CoordinateReferenceSystemObject),
		},
		"bbox": &graphql.Field{
			Type: graphql.NewList(graphql.Float),
		},
		"features": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(
				graphql.NewNonNull(FeatureObject),
			)),
		},
	},
})
