package common

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"regexp"
)

// Scalars
var DateTimeStringScalar *graphql.Scalar
var DateStringScalar *graphql.Scalar

// Common
var APIInfoStatusObject *graphql.Object
var DateTimeRangeObject *graphql.Object
var LocationObject *graphql.Object
var AreaObject *graphql.Object
var SpeedObject *graphql.Object
var RelativeHumidityObject *graphql.Object
var TemperatureObject *graphql.Object
var WindObject *graphql.Object

func init() {

	dateRegexp, err := regexp.Compile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	if err != nil {
		panic(err)
	}
	dateTimeRegexp, err := regexp.Compile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]")
	if err != nil {
		panic(err)
	}

	DateTimeStringScalar = graphql.NewScalar(graphql.ScalarConfig{
		Name:        "DatetimeString",
		Description: "Format: YYYY-MM-DD[T]HH:mm:ss",
		Serialize: func(value interface{}) interface{} {
			val := fmt.Sprintf("%v", value)
			if dateTimeRegexp.MatchString(val) {
				return val
			}
			return ""
		},
		ParseValue: func(value interface{}) interface{} {
			val := fmt.Sprintf("%v", value)
			if dateTimeRegexp.MatchString(val) {
				return val
			}
			return nil
		},
		ParseLiteral: func(valueAST ast.Value) interface{} {
			switch valueAST := valueAST.(type) {
			case *ast.StringValue:
				if dateTimeRegexp.MatchString(valueAST.Value) {
					return valueAST.Value
				}
			}
			return nil
		},
	})

	DateStringScalar = graphql.NewScalar(graphql.ScalarConfig{
		Name:        "DateString",
		Description: "Format: YYYY-MM-DD",
		Serialize: func(value interface{}) interface{} {
			val := fmt.Sprintf("%v", value)
			if dateRegexp.MatchString(val) {
				return val
			}
			return ""
		},
		ParseValue: func(value interface{}) interface{} {
			val := fmt.Sprintf("%v", value)
			if dateRegexp.MatchString(val) {
				return val
			}
			return nil
		},
		ParseLiteral: func(valueAST ast.Value) interface{} {
			switch valueAST := valueAST.(type) {
			case *ast.StringValue:
				if dateRegexp.MatchString(valueAST.Value) {
					return valueAST.Value
				}
			}
			return nil
		},
	})

	// Common
	APIInfoStatusObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "APIInfoStatus",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	DateTimeRangeObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "DateTimeRange",
		Fields: graphql.Fields{
			"start": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"end": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	LocationObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"longitude": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"latitude": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	})
	AreaObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Area",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"label_location": &graphql.Field{
				Type: graphql.NewNonNull(LocationObject),
			},
		},
	})
	SpeedObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Speed",
		Fields: graphql.Fields{
			"high": &graphql.Field{
				Type: graphql.Int,
			},
			"low": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
	RelativeHumidityObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "RelativeHumidity",
		Fields: graphql.Fields{
			"high": &graphql.Field{
				Type: graphql.Int,
			},
			"low": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
	TemperatureObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Temperature",
		Fields: graphql.Fields{
			"high": &graphql.Field{
				Type: graphql.Int,
			},
			"low": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
	WindObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Wind",
		Fields: graphql.Fields{
			"speed": &graphql.Field{
				Type: graphql.NewNonNull(SpeedObject),
			},
			"direction": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})

}
