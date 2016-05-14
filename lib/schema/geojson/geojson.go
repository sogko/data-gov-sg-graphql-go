package geojson

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/kr/pretty"
	"strconv"
)

var GeoJSONInterface *graphql.Interface
var GeometryInterface *graphql.Interface

var TypeEnum *graphql.Enum
var CoordinatesScalar *graphql.Scalar

var PointObject *graphql.Object
var MultiPointObject *graphql.Object
var LineStringObject *graphql.Object
var MultiLineStringObject *graphql.Object
var PolygonObject *graphql.Object
var MultiPolygonObject *graphql.Object
var GeometryCollectionObject *graphql.Object
var FeatureObject *graphql.Object
var FeatureCollectionObject *graphql.Object

func init() {

	TypeEnum = graphql.NewEnum(graphql.EnumConfig{
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

	CoordinatesScalar = graphql.NewScalar(graphql.ScalarConfig{
		Name: "GeoJSONCoordinates",
		Description: "GeoJSON Coordinates.\nThe \"coordinates\" member of a geometry object is composed of one position " +
			"(in the case of a Point geometry), an array of positions (LineString or MultiPoint geometries), an array of " +
			" arrays of positions (Polygons, MultiLineStrings), or a multidimensional array of positions (MultiPolygon)\n\n" +
			"A position is represented by an array of numbers. There must be at least two elements, and may be more. The " +
			"order of elements must follow x, y, z order (easting, northing, altitude for coordinates in a projected coordinate " +
			"reference system, or longitude, latitude, altitude for coordinates in a geographic coordinate reference system). " +
			"Any number of additional elements are allowed -- interpretation and meaning of additional elements is beyond the " +
			"scope of this specification.",
		Serialize:    coerceFloat32,
		ParseValue:   coerceFloat32,
		ParseLiteral: parseCoordinates,
	})

	GeoJSONInterface = graphql.NewInterface(graphql.InterfaceConfig{
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
		ResolveType: func(value interface{}, info graphql.ResolveInfo) *graphql.Object {
			valueMap := map[string]interface{}{}
			valueMap, ok := value.(map[string]interface{})
			if !ok {
				// Fallback to force-cast it to map[string]interface{}
				// Quite expensive to do this every time GraphQL needs to do ResolveType() check
				// To avoid this, marshal-unmarshal your struct to map[string]interface{}
				b, _ := json.Marshal(value)
				json.Unmarshal(b, &valueMap)
			}
			ttype, _ := valueMap["type"]
			switch ttype {
			case "Point":
				return PointObject
			case "MultiPoint":
				return MultiPointObject
			case "LineString":
				return LineStringObject
			case "MultiLineString":
				return MultiLineStringObject
			case "Polygon":
				return PolygonObject
			case "MultiPolygon":
				return MultiPolygonObject
			case "GeometryCollection":
				return GeometryCollectionObject
			case "Feature":
				return FeatureObject
			case "FeatureCollection":
				return FeatureCollectionObject
			}
			return nil
		},
	})
	GeometryInterface = graphql.NewInterface(graphql.InterfaceConfig{
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
			"coordinates": &graphql.Field{
				Type: CoordinatesScalar,
			},
		},
		ResolveType: func(value interface{}, info graphql.ResolveInfo) *graphql.Object {
			valueMap := map[string]interface{}{}
			valueMap, ok := value.(map[string]interface{})
			if !ok {
				// Fallback to force-cast it to map[string]interface{}
				// Quite expensive to do this every time GraphQL needs to do ResolveType() check
				// To avoid this, marshal-unmarshal your struct to map[string]interface{}
				b, _ := json.Marshal(value)
				json.Unmarshal(b, &valueMap)
			}
			ttype, _ := valueMap["type"]
			switch ttype {
			case "Point":
				return PointObject
			case "MultiPoint":
				return MultiPointObject
			case "LineString":
				return LineStringObject
			case "MultiLineString":
				return MultiLineStringObject
			case "Polygon":
				return PolygonObject
			case "MultiPolygon":
				return MultiPolygonObject
			}
			return nil
		},
	})

	PointObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONPoint",
		Description: "GeoJSON Point Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
				GeometryInterface,
			}
		}),
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
				Type: CoordinatesScalar,
			},
		},
	})

	MultiPointObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONMultiPoint",
		Description: "GeoJSON MultiPoint Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
				GeometryInterface,
			}
		}),
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
				Type: CoordinatesScalar,
			},
		},
	})

	LineStringObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONLineString",
		Description: "GeoJSON LineString Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
				GeometryInterface,
			}
		}),
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
				Type: CoordinatesScalar,
			},
		},
	})

	MultiLineStringObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONMultiLineString",
		Description: "GeoJSON MultiLineString Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
				GeometryInterface,
			}
		}),
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
				Type: CoordinatesScalar,
			},
		},
	})

	PolygonObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONPolygon",
		Description: "GeoJSON Polygon Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
				GeometryInterface,
			}
		}),
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
				Type: CoordinatesScalar,
			},
		},
	})

	MultiPolygonObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONMultiPolygon",
		Description: "GeoJSON MultiPolygon Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
				GeometryInterface,
			}
		}),
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
				Type: CoordinatesScalar,
			},
		},
	})

	GeometryCollectionObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONGeometryCollection",
		Description: "GeoJSON GeometryCollection Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
			}
		}),
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

	var mapObject *graphql.Object
	mapObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "GeoJSONFeatureProperties",
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			return graphql.Fields{
				"value": &graphql.Field{
					Type: graphql.String,
					Args: map[string]*graphql.ArgumentConfig{
						"key": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						pretty.Println(p.Source)
						key, _ := p.Args["key"].(string)
						source, _ := p.Source.(map[string]interface{})
						res, _ := source[key]
						return res, nil
					},
				},
				"nestedValue": &graphql.Field{
					Type: mapObject,
					Args: map[string]*graphql.ArgumentConfig{
						"key": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						pretty.Println(p.Source)
						key, _ := p.Args["key"].(string)
						source, _ := p.Source.(map[string]interface{})
						res, _ := source[key]
						return res, nil
					},
				},
			}
		}),
	})

	jsonScalar := graphql.NewScalar(graphql.ScalarConfig{
		Name:         "JSONScalar",
		Serialize:    coerceObject,
		ParseValue:   coerceObject,
		ParseLiteral: parseObject,
	})

	FeatureObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONFeature",
		Description: "GeoJSON Feature Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
			}
		}),
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
					Type: GeometryInterface,
				},
				// Interim definition
				"properties": &graphql.Field{
					Type: jsonScalar,
				},
				"id": &graphql.Field{
					Type: graphql.String,
				},
			}
		}),
	})

	FeatureCollectionObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "GeoJSONFeatureCollection",
		Description: "GeoJSON FeatureCollection Object",
		Interfaces: (graphql.InterfacesThunk)(func() []*graphql.Interface {
			return []*graphql.Interface{
				GeoJSONInterface,
			}
		}),
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

}

func parseCoordinates(valueAST ast.Value) interface{} {
	switch valueAST := valueAST.(type) {
	case *ast.ListValue:
		values := []interface{}{}
		for _, value := range valueAST.Values {
			v := parseCoordinates(value)
			if v != nil {
				values = append(values, v)
			}
		}
		return values
	case *ast.FloatValue:
		if floatValue, err := strconv.ParseFloat(valueAST.Value, 32); err == nil {
			return floatValue
		}
	case *ast.IntValue:
		if floatValue, err := strconv.ParseFloat(valueAST.Value, 32); err == nil {
			return floatValue
		}
	}
	return nil
}

func coerceFloat32(value interface{}) interface{} {
	switch value := value.(type) {
	case []float64:
		values := []interface{}{}
		for _, v := range value {
			val := coerceFloat32(v)
			values = append(values, val)
		}
		return values
	case []float32:
		return value
	case []interface{}:
		values := []interface{}{}
		for _, v := range value {
			val := coerceFloat32(v)
			values = append(values, val)
		}
		return values
	case bool:
		if value == true {
			return float32(1)
		}
		return float32(0)
	case int:
		return float32(value)
	case float32:
		return value
	case float64:
		return float32(value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceFloat32(val)
	}
	return float32(0)
}

func parseObject(valueAST ast.Value) interface{} {
	switch valueAST := valueAST.(type) {
	case *ast.ListValue:
		values := []interface{}{}
		for _, value := range valueAST.Values {
			v := parseObject(value)
			if v != nil {
				values = append(values, v)
			}
		}
		return values
	case *ast.FloatValue:
		if floatValue, err := strconv.ParseFloat(valueAST.Value, 32); err == nil {
			return floatValue
		}
	case *ast.IntValue:
		if floatValue, err := strconv.ParseFloat(valueAST.Value, 32); err == nil {
			return floatValue
		}
	case *ast.ObjectValue:
		fields := map[string]interface{}{}
		for _, field := range valueAST.Fields {
			name := ""
			if field.Name != nil {
				name = field.Name.Value
			}
			if name == "" {
				continue
			}
			fields[name] = parseObject(field.Value)
		}
		return fields
	case *ast.EnumValue:
		return valueAST.Value
	case *ast.BooleanValue:
		return valueAST.Value
	case *ast.StringValue:
		return valueAST.Value
	}
	return nil
}

func coerceObject(value interface{}) interface{} {
	valueMap := map[string]interface{}{}
	valueMap, ok := value.(map[string]interface{})
	if !ok {
		b, _ := json.Marshal(value)
		json.Unmarshal(b, &valueMap)
	}
	return valueMap
}
