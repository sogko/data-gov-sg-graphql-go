package geojson_test

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/kr/pretty"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema/geojson"
)

var TestGeoJSONSchema graphql.Schema

type testDataItem struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
type testData struct {
	Items []testDataItem `json:"items"`
}

var testDataAll testData
var testDataMap map[string]testDataItem

func init() {
	var err error

	// load geojson test data
	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &testDataAll)

	// organize data into a map
	testDataMap = map[string]testDataItem{}
	for _, item := range testDataAll.Items {
		testDataMap[item.Name] = item
	}

	// define test schema
	testItemObject := graphql.NewObject(graphql.ObjectConfig{
		Name: "TestItem",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: geojson.GeoJSONInterface,
			},
		},
	})
	TestGeoJSONSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"all": &graphql.Field{
					Type:    graphql.NewList(testItemObject),
					Resolve: resolveGeoJSONTestData("all"),
				},
				"point": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_point"),
				},
				"multipoint": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_multipoint"),
				},
				"linestring": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_linestring"),
				},
				"multilinestring": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_multilinestring"),
				},
				"polygon": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_polygon"),
				},
				"multipolygon": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_multipolygon"),
				},
				"feature": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_feature"),
				},
				"feature_collection": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_feature_collection"),
				},
				"geometry_collection": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_geometry_collection"),
				},
				"crs_named_crs": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_crs_named_crs"),
				},
				"crs_linked_crs": &graphql.Field{
					Type:    testItemObject,
					Resolve: resolveGeoJSONTestData("geojson_crs_linked_crs"),
				},
			},
		}),
	})
	if err != nil {
		panic(err)
	}
}

func resolveGeoJSONTestData(name string) func(p graphql.ResolveParams) (interface{}, error) {
	if name == "all" {
		return func(p graphql.ResolveParams) (interface{}, error) {
			return testDataAll.Items, nil
		}
	}
	res, _ := testDataMap[name]
	return func(p graphql.ResolveParams) (interface{}, error) {
		return res, nil
	}
}
func doGraphQL(schema graphql.Schema, query string) *graphql.Result {
	p := graphql.Params{
		RequestString: query,
		Schema:        schema,
	}
	return graphql.Do(p)
}

func assertGraphQLResultEqual(t *testing.T, expected *graphql.Result, result *graphql.Result) {
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected result to be equal, got diff %v", pretty.Diff(expected, result))
	}
}
func assertGraphQLResultNotEqual(t *testing.T, expected *graphql.Result, result *graphql.Result) {
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected result to be not equal, got equal")
	}
}

func TestAll(t *testing.T) {
	query := `
	{
		all {
			name
			description
			data {
			   type
			   ... Geometry
			   ... Feature
			   ... FeatureCollection
			   ... GeometryCollection
			}
		}
	}
	fragment Geometry on GeoJSONGeometryInterface {
		type
		coordinates
	}
	fragment Feature on GeoJSONFeature {
		geometry {
			... Geometry
		}
	}
	fragment FeatureCollection on GeoJSONFeatureCollection {
		features {
			... Feature
		}
	}
	fragment GeometryCollection on GeoJSONGeometryCollection {
		geometries {
			... Geometry
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"all": []interface{}{
				map[string]interface{}{
					"name":        "geojson_point",
					"description": "Returns a GeoJSON Point object",
					"data": map[string]interface{}{
						"coordinates": []interface{}{
							float32(-105.0162124633789),
							float32(39.57421875),
						},
						"type": "Point",
					},
				},
				map[string]interface{}{
					"description": "Returns a GeoJSON MultiPoint object",
					"name":        "geojson_multipoint",
					"data": map[string]interface{}{
						"coordinates": []interface{}{
							[]interface{}{
								float32(-105.0162124633789),
								float32(39.57421875),
							},
							[]interface{}{
								float32(-80.66651153564453),
								float32(35.053993225097656),
							},
						},
						"type": "MultiPoint",
					},
				},
				map[string]interface{}{
					"data": map[string]interface{}{
						"type": "LineString",
						"coordinates": []interface{}{
							[]interface{}{
								float32(-101.744384765625),
								float32(39.32154846191406),
							},
							[]interface{}{
								float32(-101.5521240234375),
								float32(39.330047607421875),
							},
							[]interface{}{
								float32(-101.40380859375),
								float32(39.330047607421875),
							},
							[]interface{}{
								float32(-101.3323974609375),
								float32(39.36403274536133),
							},
							[]interface{}{
								float32(-101.041259765625),
								float32(39.36827850341797),
							},
							[]interface{}{
								float32(-100.975341796875),
								float32(39.30455017089844),
							},
							[]interface{}{
								float32(-100.9149169921875),
								float32(39.245018005371094),
							},
							[]interface{}{
								float32(-100.843505859375),
								float32(39.16414260864258),
							},
							[]interface{}{
								float32(-100.8050537109375),
								float32(39.104488372802734),
							},
							[]interface{}{
								float32(-100.491943359375),
								float32(39.10022735595703),
							},
							[]interface{}{
								float32(-100.43701171875),
								float32(39.09596252441406),
							},
							[]interface{}{
								float32(-100.338134765625),
								float32(39.09596252441406),
							},
							[]interface{}{
								float32(-100.1953125),
								float32(39.02771759033203),
							},
							[]interface{}{
								float32(-100.008544921875),
								float32(39.01064682006836),
							},
							[]interface{}{
								float32(-99.86572265625),
								float32(39.00210952758789),
							},
							[]interface{}{
								float32(-99.6844482421875),
								float32(38.97222137451172),
							},
							[]interface{}{
								float32(-99.51416015625),
								float32(38.929500579833984),
							},
							[]interface{}{
								float32(-99.38232421875),
								float32(38.920955657958984),
							},
							[]interface{}{
								float32(-99.3218994140625),
								float32(38.89530944824219),
							},
							[]interface{}{
								float32(-99.1131591796875),
								float32(38.869651794433594),
							},
							[]interface{}{
								float32(-99.0802001953125),
								float32(38.85681915283203),
							},
							[]interface{}{
								float32(-98.822021484375),
								float32(38.85681915283203),
							},
							[]interface{}{
								float32(-98.448486328125),
								float32(38.848262786865234),
							},
							[]interface{}{
								float32(-98.206787109375),
								float32(38.848262786865234),
							},
							[]interface{}{
								float32(-98.02001953125),
								float32(38.878204345703125),
							},
							[]interface{}{
								float32(-97.635498046875),
								float32(38.87392807006836),
							},
						},
					},
					"name":        "geojson_linestring",
					"description": "Returns a GeoJSON LineString object",
				},
				map[string]interface{}{
					"name":        "geojson_multilinestring",
					"description": "Returns a GeoJSON MultiLineString object",
					"data": map[string]interface{}{
						"coordinates": []interface{}{
							[]interface{}{
								[]interface{}{
									float32(-105.02144622802734),
									float32(39.57805633544922),
								},
								[]interface{}{
									float32(-105.0215072631836),
									float32(39.57780838012695),
								},
								[]interface{}{
									float32(-105.02157592773438),
									float32(39.57749557495117),
								},
								[]interface{}{
									float32(-105.02157592773438),
									float32(39.57716369628906),
								},
								[]interface{}{
									float32(-105.02157592773438),
									float32(39.57703399658203),
								},
								[]interface{}{
									float32(-105.02153015136719),
									float32(39.5767822265625),
								},
							},
							[]interface{}{
								[]interface{}{
									float32(-105.0198974609375),
									float32(39.57499694824219),
								},
								[]interface{}{
									float32(-105.01959991455078),
									float32(39.57489776611328),
								},
								[]interface{}{
									float32(-105.01905822753906),
									float32(39.57478332519531),
								},
							},
							[]interface{}{
								[]interface{}{
									float32(-105.01717376708984),
									float32(39.57440185546875),
								},
								[]interface{}{
									float32(-105.01698303222656),
									float32(39.57438659667969),
								},
								[]interface{}{
									float32(-105.01663970947266),
									float32(39.57438659667969),
								},
								[]interface{}{
									float32(-105.01651000976562),
									float32(39.57440185546875),
								},
								[]interface{}{
									float32(-105.01595306396484),
									float32(39.57426834106445),
								},
							},
							[]interface{}{
								[]interface{}{
									float32(-105.01427459716797),
									float32(39.573970794677734),
								},
								[]interface{}{
									float32(-105.01412963867188),
									float32(39.574039459228516),
								},
								[]interface{}{
									float32(-105.01382446289062),
									float32(39.57416915893555),
								},
								[]interface{}{
									float32(-105.01331329345703),
									float32(39.5744514465332),
								},
							},
						},
						"type": "MultiLineString",
					},
				},
				map[string]interface{}{
					"name":        "geojson_polygon",
					"description": "Returns a GeoJSON Polygon object",
					"data": map[string]interface{}{
						"coordinates": []interface{}{
							[]interface{}{
								[]interface{}{
									float32(-84.32281494140625),
									float32(34.989501953125),
								},
								[]interface{}{
									float32(-84.29122924804688),
									float32(35.219818115234375),
								},
								[]interface{}{
									float32(-84.24041748046875),
									float32(35.25458908081055),
								},
								[]interface{}{
									float32(-84.22531127929688),
									float32(35.26692581176758),
								},
								[]interface{}{
									float32(-84.20745849609375),
									float32(35.265804290771484),
								},
								[]interface{}{
									float32(-84.19921875),
									float32(35.246742248535156),
								},
								[]interface{}{
									float32(-84.16213989257812),
									float32(35.24113464355469),
								},
								[]interface{}{
									float32(-84.12368774414062),
									float32(35.248985290527344),
								},
								[]interface{}{
									float32(-84.09072875976562),
									float32(35.248985290527344),
								},
								[]interface{}{
									float32(-84.08798217773438),
									float32(35.26468276977539),
								},
								[]interface{}{
									float32(-84.04266357421875),
									float32(35.277015686035156),
								},
								[]interface{}{
									float32(-84.03030395507812),
									float32(35.291587829589844),
								},
								[]interface{}{
									float32(-84.0234375),
									float32(35.30615997314453),
								},
								[]interface{}{
									float32(-84.03305053710938),
									float32(35.327449798583984),
								},
								[]interface{}{
									float32(-84.03579711914062),
									float32(35.343135833740234),
								},
								[]interface{}{
									float32(-84.03579711914062),
									float32(35.34873580932617),
								},
								[]interface{}{
									float32(-84.01657104492188),
									float32(35.3554573059082),
								},
								[]interface{}{
									float32(-84.01107788085938),
									float32(35.373374938964844),
								},
								[]interface{}{
									float32(-84.00970458984375),
									float32(35.39128875732422),
								},
								[]interface{}{
									float32(-84.01931762695312),
									float32(35.414794921875),
								},
								[]interface{}{
									float32(-84.00283813476562),
									float32(35.429344177246094),
								},
								[]interface{}{
									float32(-83.93692016601562),
									float32(35.474090576171875),
								},
								[]interface{}{
									float32(-83.91220092773438),
									float32(35.4763298034668),
								},
								[]interface{}{
									float32(-83.88885498046875),
									float32(35.5042839050293),
								},
								[]interface{}{
									float32(-83.88473510742188),
									float32(35.516578674316406),
								},
								[]interface{}{
									float32(-83.8751220703125),
									float32(35.52104949951172),
								},
								[]interface{}{
									float32(-83.8531494140625),
									float32(35.52104949951172),
								},
								[]interface{}{
									float32(-83.82843017578125),
									float32(35.52104949951172),
								},
								[]interface{}{
									float32(-83.8092041015625),
									float32(35.534461975097656),
								},
								[]interface{}{
									float32(-83.80233764648438),
									float32(35.54116439819336),
								},
								[]interface{}{
									float32(-83.76800537109375),
									float32(35.56239318847656),
								},
								[]interface{}{
									float32(-83.7432861328125),
									float32(35.56239318847656),
								},
								[]interface{}{
									float32(-83.71994018554688),
									float32(35.56239318847656),
								},
								[]interface{}{
									float32(-83.67050170898438),
									float32(35.56909942626953),
								},
								[]interface{}{
									float32(-83.6334228515625),
									float32(35.570213317871094),
								},
								[]interface{}{
									float32(-83.61007690429688),
									float32(35.5769157409668),
								},
								[]interface{}{
									float32(-83.59634399414062),
									float32(35.574684143066406),
								},
								[]interface{}{
									float32(-83.5894775390625),
									float32(35.559043884277344),
								},
								[]interface{}{
									float32(-83.55239868164062),
									float32(35.56574630737305),
								},
								[]interface{}{
									float32(-83.49746704101562),
									float32(35.56351089477539),
								},
								[]interface{}{
									float32(-83.47000122070312),
									float32(35.58696746826172),
								},
								[]interface{}{
									float32(-83.4466552734375),
									float32(35.608184814453125),
								},
								[]interface{}{
									float32(-83.37936401367188),
									float32(35.63609313964844),
								},
								[]interface{}{
									float32(-83.35739135742188),
									float32(35.65618133544922),
								},
								[]interface{}{
									float32(-83.32305908203125),
									float32(35.666221618652344),
								},
								[]interface{}{
									float32(-83.3148193359375),
									float32(35.65394973754883),
								},
								[]interface{}{
									float32(-83.29971313476562),
									float32(35.66064453125),
								},
								[]interface{}{
									float32(-83.28598022460938),
									float32(35.67180252075195),
								},
								[]interface{}{
									float32(-83.26126098632812),
									float32(35.690765380859375),
								},
								[]interface{}{
									float32(-83.25714111328125),
									float32(35.69968795776367),
								},
								[]interface{}{
									float32(-83.25576782226562),
									float32(35.71529769897461),
								},
								[]interface{}{
									float32(-83.23516845703125),
									float32(35.72310256958008),
								},
								[]interface{}{
									float32(-83.19808959960938),
									float32(35.727561950683594),
								},
								[]interface{}{
									float32(-83.16238403320312),
									float32(35.75320053100586),
								},
								[]interface{}{
									float32(-83.15826416015625),
									float32(35.76322937011719),
								},
								[]interface{}{
									float32(-83.10333251953125),
									float32(35.76991653442383),
								},
								[]interface{}{
									float32(-83.08685302734375),
									float32(35.78439712524414),
								},
								[]interface{}{
									float32(-83.0511474609375),
									float32(35.787742614746094),
								},
								[]interface{}{
									float32(-83.01681518554688),
									float32(35.78328323364258),
								},
								[]interface{}{
									float32(-83.001708984375),
									float32(35.77882766723633),
								},
								[]interface{}{
									float32(-82.96737670898438),
									float32(35.793312072753906),
								},
								[]interface{}{
									float32(-82.94540405273438),
									float32(35.82004165649414),
								},
								[]interface{}{
									float32(-82.9193115234375),
									float32(35.85121154785156),
								},
								[]interface{}{
									float32(-82.9083251953125),
									float32(35.869022369384766),
								},
								[]interface{}{
									float32(-82.90557861328125),
									float32(35.87792205810547),
								},
								[]interface{}{
									float32(-82.91244506835938),
									float32(35.92353057861328),
								},
								[]interface{}{
									float32(-82.88360595703125),
									float32(35.94688415527344),
								},
								[]interface{}{
									float32(-82.85614013671875),
									float32(35.95132827758789),
								},
								[]interface{}{
									float32(-82.8424072265625),
									float32(35.94243621826172),
								},
								[]interface{}{
									float32(-82.825927734375),
									float32(35.924644470214844),
								},
								[]interface{}{
									float32(-82.80670166015625),
									float32(35.927982330322266),
								},
								[]interface{}{
									float32(-82.80532836914062),
									float32(35.94243621826172),
								},
								[]interface{}{
									float32(-82.77923583984375),
									float32(35.97356033325195),
								},
								[]interface{}{
									float32(-82.78060913085938),
									float32(35.99245071411133),
								},
								[]interface{}{
									float32(-82.76138305664062),
									float32(36.003562927246094),
								},
								[]interface{}{
									float32(-82.69546508789062),
									float32(36.04465866088867),
								},
								[]interface{}{
									float32(-82.6446533203125),
									float32(36.06019973754883),
								},
								[]interface{}{
									float32(-82.61306762695312),
									float32(36.06019973754883),
								},
								[]interface{}{
									float32(-82.606201171875),
									float32(36.03355407714844),
								},
								[]interface{}{
									float32(-82.606201171875),
									float32(35.99134063720703),
								},
								[]interface{}{
									float32(-82.606201171875),
									float32(35.97911834716797),
								},
								[]interface{}{
									float32(-82.5787353515625),
									float32(35.961334228515625),
								},
								[]interface{}{
									float32(-82.5677490234375),
									float32(35.95132827758789),
								},
								[]interface{}{
									float32(-82.53067016601562),
									float32(35.972450256347656),
								},
								[]interface{}{
									float32(-82.46475219726562),
									float32(36.00689697265625),
								},
								[]interface{}{
									float32(-82.41668701171875),
									float32(36.0701904296875),
								},
								[]interface{}{
									float32(-82.37960815429688),
									float32(36.10126876831055),
								},
								[]interface{}{
									float32(-82.35488891601562),
									float32(36.1179084777832),
								},
								[]interface{}{
									float32(-82.34115600585938),
									float32(36.11347198486328),
								},
								[]interface{}{
									float32(-82.29583740234375),
									float32(36.13343811035156),
								},
								[]interface{}{
									float32(-82.26287841796875),
									float32(36.135658264160156),
								},
								[]interface{}{
									float32(-82.23403930664062),
									float32(36.135658264160156),
								},
								[]interface{}{
									float32(-82.2216796875),
									float32(36.154510498046875),
								},
								[]interface{}{
									float32(-82.20382690429688),
									float32(36.155616760253906),
								},
								[]interface{}{
									float32(-82.19009399414062),
									float32(36.144527435302734),
								},
								[]interface{}{
									float32(-82.15438842773438),
									float32(36.15007400512695),
								},
								[]interface{}{
									float32(-82.14065551757812),
									float32(36.13454818725586),
								},
								[]interface{}{
									float32(-82.1337890625),
									float32(36.116798400878906),
								},
								[]interface{}{
									float32(-82.12142944335938),
									float32(36.10570526123047),
								},
								[]interface{}{
									float32(-82.08984375),
									float32(36.10792541503906),
								},
								[]interface{}{
									float32(-82.05276489257812),
									float32(36.12678146362305),
								},
								[]interface{}{
									float32(-82.03628540039062),
									float32(36.12900161743164),
								},
								[]interface{}{
									float32(-81.91268920898438),
									float32(36.294097900390625),
								},
								[]interface{}{
									float32(-81.89071655273438),
									float32(36.309593200683594),
								},
								[]interface{}{
									float32(-81.86325073242188),
									float32(36.33504104614258),
								},
								[]interface{}{
									float32(-81.83029174804688),
									float32(36.34499740600586),
								},
								[]interface{}{
									float32(-81.80145263671875),
									float32(36.356056213378906),
								},
								[]interface{}{
									float32(-81.77947998046875),
									float32(36.34610366821289),
								},
								[]interface{}{
									float32(-81.76162719726562),
									float32(36.33835983276367),
								},
								[]interface{}{
									float32(-81.73690795898438),
									float32(36.33835983276367),
								},
								[]interface{}{
									float32(-81.71905517578125),
									float32(36.33835983276367),
								},
								[]interface{}{
									float32(-81.70669555664062),
									float32(36.33504104614258),
								},
								[]interface{}{
									float32(-81.70669555664062),
									float32(36.3427848815918),
								},
								[]interface{}{
									float32(-81.72317504882812),
									float32(36.35716247558594),
								},
								[]interface{}{
									float32(-81.7327880859375),
									float32(36.37928009033203),
								},
								[]interface{}{
									float32(-81.73690795898438),
									float32(36.40028381347656),
								},
								[]interface{}{
									float32(-81.73690795898438),
									float32(36.41354751586914),
								},
								[]interface{}{
									float32(-81.72454833984375),
									float32(36.423492431640625),
								},
								[]interface{}{
									float32(-81.71768188476562),
									float32(36.44559097290039),
								},
								[]interface{}{
									float32(-81.69845581054688),
									float32(36.47541046142578),
								},
								[]interface{}{
									float32(-81.69845581054688),
									float32(36.510738372802734),
								},
								[]interface{}{
									float32(-81.705322265625),
									float32(36.53060531616211),
								},
								[]interface{}{
									float32(-81.69158935546875),
									float32(36.55929183959961),
								},
								[]interface{}{
									float32(-81.68060302734375),
									float32(36.5648078918457),
								},
								[]interface{}{
									float32(-81.68197631835938),
									float32(36.58686447143555),
								},
								[]interface{}{
									float32(-81.04202270507812),
									float32(36.56370162963867),
								},
								[]interface{}{
									float32(-80.74264526367188),
									float32(36.56149673461914),
								},
								[]interface{}{
									float32(-79.89120483398438),
									float32(36.54053497314453),
								},
								[]interface{}{
									float32(-78.68408203125),
									float32(36.539432525634766),
								},
								[]interface{}{
									float32(-77.88345336914062),
									float32(36.54053497314453),
								},
								[]interface{}{
									float32(-76.91665649414062),
									float32(36.54164123535156),
								},
								[]interface{}{
									float32(-76.91665649414062),
									float32(36.55046463012695),
								},
								[]interface{}{
									float32(-76.31103515625),
									float32(36.55156707763672),
								},
								[]interface{}{
									float32(-75.79605102539062),
									float32(36.54936218261719),
								},
								[]interface{}{
									float32(-75.6298828125),
									float32(36.075740814208984),
								},
								[]interface{}{
									float32(-75.4925537109375),
									float32(35.822265625),
								},
								[]interface{}{
									float32(-75.3936767578125),
									float32(35.639442443847656),
								},
								[]interface{}{
									float32(-75.41015625),
									float32(35.438297271728516),
								},
								[]interface{}{
									float32(-75.43212890625),
									float32(35.2635612487793),
								},
								[]interface{}{
									float32(-75.487060546875),
									float32(35.187278747558594),
								},
								[]interface{}{
									float32(-75.5914306640625),
									float32(35.17380905151367),
								},
								[]interface{}{
									float32(-75.9210205078125),
									float32(35.0479850769043),
								},
								[]interface{}{
									float32(-76.17919921875),
									float32(34.86790466308594),
								},
								[]interface{}{
									float32(-76.4154052734375),
									float32(34.62868881225586),
								},
								[]interface{}{
									float32(-76.4593505859375),
									float32(34.57442855834961),
								},
								[]interface{}{
									float32(-76.53076171875),
									float32(34.533714294433594),
								},
								[]interface{}{
									float32(-76.5911865234375),
									float32(34.55181121826172),
								},
								[]interface{}{
									float32(-76.651611328125),
									float32(34.61512756347656),
								},
								[]interface{}{
									float32(-76.761474609375),
									float32(34.633209228515625),
								},
								[]interface{}{
									float32(-77.069091796875),
									float32(34.597042083740234),
								},
								[]interface{}{
									float32(-77.376708984375),
									float32(34.456748962402344),
								},
								[]interface{}{
									float32(-77.5909423828125),
									float32(34.32075500488281),
								},
								[]interface{}{
									float32(-77.8326416015625),
									float32(33.97980880737305),
								},
								[]interface{}{
									float32(-77.9150390625),
									float32(33.80197525024414),
								},
								[]interface{}{
									float32(-77.9754638671875),
									float32(33.73804473876953),
								},
								[]interface{}{
									float32(-78.11279296875),
									float32(33.852169036865234),
								},
								[]interface{}{
									float32(-78.2830810546875),
									float32(33.852169036865234),
								},
								[]interface{}{
									float32(-78.4808349609375),
									float32(33.81566619873047),
								},
								[]interface{}{
									float32(-79.6728515625),
									float32(34.80478286743164),
								},
								[]interface{}{
									float32(-80.782470703125),
									float32(34.83634948730469),
								},
								[]interface{}{
									float32(-80.782470703125),
									float32(34.91746520996094),
								},
								[]interface{}{
									float32(-80.9307861328125),
									float32(35.09294509887695),
								},
								[]interface{}{
									float32(-81.0516357421875),
									float32(35.02999496459961),
								},
								[]interface{}{
									float32(-81.0516357421875),
									float32(35.05248260498047),
								},
								[]interface{}{
									float32(-81.0516357421875),
									float32(35.13787841796875),
								},
								[]interface{}{
									float32(-82.3150634765625),
									float32(35.19625473022461),
								},
								[]interface{}{
									float32(-82.3590087890625),
									float32(35.19625473022461),
								},
								[]interface{}{
									float32(-82.4029541015625),
									float32(35.22318649291992),
								},
								[]interface{}{
									float32(-82.4688720703125),
									float32(35.16931915283203),
								},
								[]interface{}{
									float32(-82.6885986328125),
									float32(35.115413665771484),
								},
								[]interface{}{
									float32(-82.781982421875),
									float32(35.06147766113281),
								},
								[]interface{}{
									float32(-83.1060791015625),
									float32(35.00300216674805),
								},
								[]interface{}{
									float32(-83.616943359375),
									float32(34.998504638671875),
								},
								[]interface{}{
									float32(-84.056396484375),
									float32(34.98500442504883),
								},
								[]interface{}{
									float32(-84.22119140625),
									float32(34.98500442504883),
								},
								[]interface{}{
									float32(-84.32281494140625),
									float32(34.989501953125),
								},
							},
							[]interface{}{
								[]interface{}{
									float32(-75.6903076171875),
									float32(35.7420539855957),
								},
								[]interface{}{
									float32(-75.5914306640625),
									float32(35.7420539855957),
								},
								[]interface{}{
									float32(-75.5419921875),
									float32(35.58584976196289),
								},
								[]interface{}{
									float32(-75.56396484375),
									float32(35.326332092285156),
								},
								[]interface{}{
									float32(-75.6903076171875),
									float32(35.28598403930664),
								},
								[]interface{}{
									float32(-75.970458984375),
									float32(35.16482925415039),
								},
								[]interface{}{
									float32(-76.2066650390625),
									float32(34.99400329589844),
								},
								[]interface{}{
									float32(-76.300048828125),
									float32(35.02999496459961),
								},
								[]interface{}{
									float32(-76.409912109375),
									float32(35.07946014404297),
								},
								[]interface{}{
									float32(-76.5252685546875),
									float32(35.10642623901367),
								},
								[]interface{}{
									float32(-76.4208984375),
									float32(35.25907516479492),
								},
								[]interface{}{
									float32(-76.3385009765625),
									float32(35.294952392578125),
								},
								[]interface{}{
									float32(-76.0858154296875),
									float32(35.299434661865234),
								},
								[]interface{}{
									float32(-75.948486328125),
									float32(35.442771911621094),
								},
								[]interface{}{
									float32(-75.8660888671875),
									float32(35.53669738769531),
								},
								[]interface{}{
									float32(-75.772705078125),
									float32(35.5679817199707),
								},
								[]interface{}{
									float32(-75.706787109375),
									float32(35.63497543334961),
								},
								[]interface{}{
									float32(-75.706787109375),
									float32(35.7420539855957),
								},
								[]interface{}{
									float32(-75.6903076171875),
									float32(35.7420539855957),
								},
							},
						},
						"type": "Polygon",
					},
				},
				map[string]interface{}{
					"description": "Returns a GeoJSON MultiPolygon object",
					"data": map[string]interface{}{
						"type": "MultiPolygon",
						"coordinates": []interface{}{
							[]interface{}{
								[]interface{}{
									[]interface{}{
										float32(-84.32281494140625),
										float32(34.989501953125),
									},
									[]interface{}{
										float32(-84.29122924804688),
										float32(35.219818115234375),
									},
									[]interface{}{
										float32(-84.24041748046875),
										float32(35.25458908081055),
									},
									[]interface{}{
										float32(-84.22531127929688),
										float32(35.26692581176758),
									},
									[]interface{}{
										float32(-84.20745849609375),
										float32(35.265804290771484),
									},
									[]interface{}{
										float32(-84.19921875),
										float32(35.246742248535156),
									},
									[]interface{}{
										float32(-84.16213989257812),
										float32(35.24113464355469),
									},
									[]interface{}{
										float32(-84.12368774414062),
										float32(35.248985290527344),
									},
									[]interface{}{
										float32(-84.09072875976562),
										float32(35.248985290527344),
									},
									[]interface{}{
										float32(-84.08798217773438),
										float32(35.26468276977539),
									},
									[]interface{}{
										float32(-84.04266357421875),
										float32(35.277015686035156),
									},
									[]interface{}{
										float32(-84.03030395507812),
										float32(35.291587829589844),
									},
									[]interface{}{
										float32(-84.0234375),
										float32(35.30615997314453),
									},
									[]interface{}{
										float32(-84.03305053710938),
										float32(35.327449798583984),
									},
									[]interface{}{
										float32(-84.03579711914062),
										float32(35.343135833740234),
									},
									[]interface{}{
										float32(-84.03579711914062),
										float32(35.34873580932617),
									},
									[]interface{}{
										float32(-84.01657104492188),
										float32(35.3554573059082),
									},
									[]interface{}{
										float32(-84.01107788085938),
										float32(35.373374938964844),
									},
									[]interface{}{
										float32(-84.00970458984375),
										float32(35.39128875732422),
									},
									[]interface{}{
										float32(-84.01931762695312),
										float32(35.414794921875),
									},
									[]interface{}{
										float32(-84.00283813476562),
										float32(35.429344177246094),
									},
									[]interface{}{
										float32(-83.93692016601562),
										float32(35.474090576171875),
									},
									[]interface{}{
										float32(-83.91220092773438),
										float32(35.4763298034668),
									},
									[]interface{}{
										float32(-83.88885498046875),
										float32(35.5042839050293),
									},
									[]interface{}{
										float32(-83.88473510742188),
										float32(35.516578674316406),
									},
									[]interface{}{
										float32(-83.8751220703125),
										float32(35.52104949951172),
									},
									[]interface{}{
										float32(-83.8531494140625),
										float32(35.52104949951172),
									},
									[]interface{}{
										float32(-83.82843017578125),
										float32(35.52104949951172),
									},
									[]interface{}{
										float32(-83.8092041015625),
										float32(35.534461975097656),
									},
									[]interface{}{
										float32(-83.80233764648438),
										float32(35.54116439819336),
									},
									[]interface{}{
										float32(-83.76800537109375),
										float32(35.56239318847656),
									},
									[]interface{}{
										float32(-83.7432861328125),
										float32(35.56239318847656),
									},
									[]interface{}{
										float32(-83.71994018554688),
										float32(35.56239318847656),
									},
									[]interface{}{
										float32(-83.67050170898438),
										float32(35.56909942626953),
									},
									[]interface{}{
										float32(-83.6334228515625),
										float32(35.570213317871094),
									},
									[]interface{}{
										float32(-83.61007690429688),
										float32(35.5769157409668),
									},
									[]interface{}{
										float32(-83.59634399414062),
										float32(35.574684143066406),
									},
									[]interface{}{
										float32(-83.5894775390625),
										float32(35.559043884277344),
									},
									[]interface{}{
										float32(-83.55239868164062),
										float32(35.56574630737305),
									},
									[]interface{}{
										float32(-83.49746704101562),
										float32(35.56351089477539),
									},
									[]interface{}{
										float32(-83.47000122070312),
										float32(35.58696746826172),
									},
									[]interface{}{
										float32(-83.4466552734375),
										float32(35.608184814453125),
									},
									[]interface{}{
										float32(-83.37936401367188),
										float32(35.63609313964844),
									},
									[]interface{}{
										float32(-83.35739135742188),
										float32(35.65618133544922),
									},
									[]interface{}{
										float32(-83.32305908203125),
										float32(35.666221618652344),
									},
									[]interface{}{
										float32(-83.3148193359375),
										float32(35.65394973754883),
									},
									[]interface{}{
										float32(-83.29971313476562),
										float32(35.66064453125),
									},
									[]interface{}{
										float32(-83.28598022460938),
										float32(35.67180252075195),
									},
									[]interface{}{
										float32(-83.26126098632812),
										float32(35.690765380859375),
									},
									[]interface{}{
										float32(-83.25714111328125),
										float32(35.69968795776367),
									},
									[]interface{}{
										float32(-83.25576782226562),
										float32(35.71529769897461),
									},
									[]interface{}{
										float32(-83.23516845703125),
										float32(35.72310256958008),
									},
									[]interface{}{
										float32(-83.19808959960938),
										float32(35.727561950683594),
									},
									[]interface{}{
										float32(-83.16238403320312),
										float32(35.75320053100586),
									},
									[]interface{}{
										float32(-83.15826416015625),
										float32(35.76322937011719),
									},
									[]interface{}{
										float32(-83.10333251953125),
										float32(35.76991653442383),
									},
									[]interface{}{
										float32(-83.08685302734375),
										float32(35.78439712524414),
									},
									[]interface{}{
										float32(-83.0511474609375),
										float32(35.787742614746094),
									},
									[]interface{}{
										float32(-83.01681518554688),
										float32(35.78328323364258),
									},
									[]interface{}{
										float32(-83.001708984375),
										float32(35.77882766723633),
									},
									[]interface{}{
										float32(-82.96737670898438),
										float32(35.793312072753906),
									},
									[]interface{}{
										float32(-82.94540405273438),
										float32(35.82004165649414),
									},
									[]interface{}{
										float32(-82.9193115234375),
										float32(35.85121154785156),
									},
									[]interface{}{
										float32(-82.9083251953125),
										float32(35.869022369384766),
									},
									[]interface{}{
										float32(-82.90557861328125),
										float32(35.87792205810547),
									},
									[]interface{}{
										float32(-82.91244506835938),
										float32(35.92353057861328),
									},
									[]interface{}{
										float32(-82.88360595703125),
										float32(35.94688415527344),
									},
									[]interface{}{
										float32(-82.85614013671875),
										float32(35.95132827758789),
									},
									[]interface{}{
										float32(-82.8424072265625),
										float32(35.94243621826172),
									},
									[]interface{}{
										float32(-82.825927734375),
										float32(35.924644470214844),
									},
									[]interface{}{
										float32(-82.80670166015625),
										float32(35.927982330322266),
									},
									[]interface{}{
										float32(-82.80532836914062),
										float32(35.94243621826172),
									},
									[]interface{}{
										float32(-82.77923583984375),
										float32(35.97356033325195),
									},
									[]interface{}{
										float32(-82.78060913085938),
										float32(35.99245071411133),
									},
									[]interface{}{
										float32(-82.76138305664062),
										float32(36.003562927246094),
									},
									[]interface{}{
										float32(-82.69546508789062),
										float32(36.04465866088867),
									},
									[]interface{}{
										float32(-82.6446533203125),
										float32(36.06019973754883),
									},
									[]interface{}{
										float32(-82.61306762695312),
										float32(36.06019973754883),
									},
									[]interface{}{
										float32(-82.606201171875),
										float32(36.03355407714844),
									},
									[]interface{}{
										float32(-82.606201171875),
										float32(35.99134063720703),
									},
									[]interface{}{
										float32(-82.606201171875),
										float32(35.97911834716797),
									},
									[]interface{}{
										float32(-82.5787353515625),
										float32(35.961334228515625),
									},
									[]interface{}{
										float32(-82.5677490234375),
										float32(35.95132827758789),
									},
									[]interface{}{
										float32(-82.53067016601562),
										float32(35.972450256347656),
									},
									[]interface{}{
										float32(-82.46475219726562),
										float32(36.00689697265625),
									},
									[]interface{}{
										float32(-82.41668701171875),
										float32(36.0701904296875),
									},
									[]interface{}{
										float32(-82.37960815429688),
										float32(36.10126876831055),
									},
									[]interface{}{
										float32(-82.35488891601562),
										float32(36.1179084777832),
									},
									[]interface{}{
										float32(-82.34115600585938),
										float32(36.11347198486328),
									},
									[]interface{}{
										float32(-82.29583740234375),
										float32(36.13343811035156),
									},
									[]interface{}{
										float32(-82.26287841796875),
										float32(36.135658264160156),
									},
									[]interface{}{
										float32(-82.23403930664062),
										float32(36.135658264160156),
									},
									[]interface{}{
										float32(-82.2216796875),
										float32(36.154510498046875),
									},
									[]interface{}{
										float32(-82.20382690429688),
										float32(36.155616760253906),
									},
									[]interface{}{
										float32(-82.19009399414062),
										float32(36.144527435302734),
									},
									[]interface{}{
										float32(-82.15438842773438),
										float32(36.15007400512695),
									},
									[]interface{}{
										float32(-82.14065551757812),
										float32(36.13454818725586),
									},
									[]interface{}{
										float32(-82.1337890625),
										float32(36.116798400878906),
									},
									[]interface{}{
										float32(-82.12142944335938),
										float32(36.10570526123047),
									},
									[]interface{}{
										float32(-82.08984375),
										float32(36.10792541503906),
									},
									[]interface{}{
										float32(-82.05276489257812),
										float32(36.12678146362305),
									},
									[]interface{}{
										float32(-82.03628540039062),
										float32(36.12900161743164),
									},
									[]interface{}{
										float32(-81.91268920898438),
										float32(36.294097900390625),
									},
									[]interface{}{
										float32(-81.89071655273438),
										float32(36.309593200683594),
									},
									[]interface{}{
										float32(-81.86325073242188),
										float32(36.33504104614258),
									},
									[]interface{}{
										float32(-81.83029174804688),
										float32(36.34499740600586),
									},
									[]interface{}{
										float32(-81.80145263671875),
										float32(36.356056213378906),
									},
									[]interface{}{
										float32(-81.77947998046875),
										float32(36.34610366821289),
									},
									[]interface{}{
										float32(-81.76162719726562),
										float32(36.33835983276367),
									},
									[]interface{}{
										float32(-81.73690795898438),
										float32(36.33835983276367),
									},
									[]interface{}{
										float32(-81.71905517578125),
										float32(36.33835983276367),
									},
									[]interface{}{
										float32(-81.70669555664062),
										float32(36.33504104614258),
									},
									[]interface{}{
										float32(-81.70669555664062),
										float32(36.3427848815918),
									},
									[]interface{}{
										float32(-81.72317504882812),
										float32(36.35716247558594),
									},
									[]interface{}{
										float32(-81.7327880859375),
										float32(36.37928009033203),
									},
									[]interface{}{
										float32(-81.73690795898438),
										float32(36.40028381347656),
									},
									[]interface{}{
										float32(-81.73690795898438),
										float32(36.41354751586914),
									},
									[]interface{}{
										float32(-81.72454833984375),
										float32(36.423492431640625),
									},
									[]interface{}{
										float32(-81.71768188476562),
										float32(36.44559097290039),
									},
									[]interface{}{
										float32(-81.69845581054688),
										float32(36.47541046142578),
									},
									[]interface{}{
										float32(-81.69845581054688),
										float32(36.510738372802734),
									},
									[]interface{}{
										float32(-81.705322265625),
										float32(36.53060531616211),
									},
									[]interface{}{
										float32(-81.69158935546875),
										float32(36.55929183959961),
									},
									[]interface{}{
										float32(-81.68060302734375),
										float32(36.5648078918457),
									},
									[]interface{}{
										float32(-81.68197631835938),
										float32(36.58686447143555),
									},
									[]interface{}{
										float32(-81.04202270507812),
										float32(36.56370162963867),
									},
									[]interface{}{
										float32(-80.74264526367188),
										float32(36.56149673461914),
									},
									[]interface{}{
										float32(-79.89120483398438),
										float32(36.54053497314453),
									},
									[]interface{}{
										float32(-78.68408203125),
										float32(36.539432525634766),
									},
									[]interface{}{
										float32(-77.88345336914062),
										float32(36.54053497314453),
									},
									[]interface{}{
										float32(-76.91665649414062),
										float32(36.54164123535156),
									},
									[]interface{}{
										float32(-76.91665649414062),
										float32(36.55046463012695),
									},
									[]interface{}{
										float32(-76.31103515625),
										float32(36.55156707763672),
									},
									[]interface{}{
										float32(-75.79605102539062),
										float32(36.54936218261719),
									},
									[]interface{}{
										float32(-75.6298828125),
										float32(36.075740814208984),
									},
									[]interface{}{
										float32(-75.4925537109375),
										float32(35.822265625),
									},
									[]interface{}{
										float32(-75.3936767578125),
										float32(35.639442443847656),
									},
									[]interface{}{
										float32(-75.41015625),
										float32(35.438297271728516),
									},
									[]interface{}{
										float32(-75.43212890625),
										float32(35.2635612487793),
									},
									[]interface{}{
										float32(-75.487060546875),
										float32(35.187278747558594),
									},
									[]interface{}{
										float32(-75.5914306640625),
										float32(35.17380905151367),
									},
									[]interface{}{
										float32(-75.9210205078125),
										float32(35.0479850769043),
									},
									[]interface{}{
										float32(-76.17919921875),
										float32(34.86790466308594),
									},
									[]interface{}{
										float32(-76.4154052734375),
										float32(34.62868881225586),
									},
									[]interface{}{
										float32(-76.4593505859375),
										float32(34.57442855834961),
									},
									[]interface{}{
										float32(-76.53076171875),
										float32(34.533714294433594),
									},
									[]interface{}{
										float32(-76.5911865234375),
										float32(34.55181121826172),
									},
									[]interface{}{
										float32(-76.651611328125),
										float32(34.61512756347656),
									},
									[]interface{}{
										float32(-76.761474609375),
										float32(34.633209228515625),
									},
									[]interface{}{
										float32(-77.069091796875),
										float32(34.597042083740234),
									},
									[]interface{}{
										float32(-77.376708984375),
										float32(34.456748962402344),
									},
									[]interface{}{
										float32(-77.5909423828125),
										float32(34.32075500488281),
									},
									[]interface{}{
										float32(-77.8326416015625),
										float32(33.97980880737305),
									},
									[]interface{}{
										float32(-77.9150390625),
										float32(33.80197525024414),
									},
									[]interface{}{
										float32(-77.9754638671875),
										float32(33.73804473876953),
									},
									[]interface{}{
										float32(-78.11279296875),
										float32(33.852169036865234),
									},
									[]interface{}{
										float32(-78.2830810546875),
										float32(33.852169036865234),
									},
									[]interface{}{
										float32(-78.4808349609375),
										float32(33.81566619873047),
									},
									[]interface{}{
										float32(-79.6728515625),
										float32(34.80478286743164),
									},
									[]interface{}{
										float32(-80.782470703125),
										float32(34.83634948730469),
									},
									[]interface{}{
										float32(-80.782470703125),
										float32(34.91746520996094),
									},
									[]interface{}{
										float32(-80.9307861328125),
										float32(35.09294509887695),
									},
									[]interface{}{
										float32(-81.0516357421875),
										float32(35.02999496459961),
									},
									[]interface{}{
										float32(-81.0516357421875),
										float32(35.05248260498047),
									},
									[]interface{}{
										float32(-81.0516357421875),
										float32(35.13787841796875),
									},
									[]interface{}{
										float32(-82.3150634765625),
										float32(35.19625473022461),
									},
									[]interface{}{
										float32(-82.3590087890625),
										float32(35.19625473022461),
									},
									[]interface{}{
										float32(-82.4029541015625),
										float32(35.22318649291992),
									},
									[]interface{}{
										float32(-82.4688720703125),
										float32(35.16931915283203),
									},
									[]interface{}{
										float32(-82.6885986328125),
										float32(35.115413665771484),
									},
									[]interface{}{
										float32(-82.781982421875),
										float32(35.06147766113281),
									},
									[]interface{}{
										float32(-83.1060791015625),
										float32(35.00300216674805),
									},
									[]interface{}{
										float32(-83.616943359375),
										float32(34.998504638671875),
									},
									[]interface{}{
										float32(-84.056396484375),
										float32(34.98500442504883),
									},
									[]interface{}{
										float32(-84.22119140625),
										float32(34.98500442504883),
									},
									[]interface{}{
										float32(-84.32281494140625),
										float32(34.989501953125),
									},
								},
								[]interface{}{
									[]interface{}{
										float32(-75.6903076171875),
										float32(35.7420539855957),
									},
									[]interface{}{
										float32(-75.5914306640625),
										float32(35.7420539855957),
									},
									[]interface{}{
										float32(-75.5419921875),
										float32(35.58584976196289),
									},
									[]interface{}{
										float32(-75.56396484375),
										float32(35.326332092285156),
									},
									[]interface{}{
										float32(-75.6903076171875),
										float32(35.28598403930664),
									},
									[]interface{}{
										float32(-75.970458984375),
										float32(35.16482925415039),
									},
									[]interface{}{
										float32(-76.2066650390625),
										float32(34.99400329589844),
									},
									[]interface{}{
										float32(-76.300048828125),
										float32(35.02999496459961),
									},
									[]interface{}{
										float32(-76.409912109375),
										float32(35.07946014404297),
									},
									[]interface{}{
										float32(-76.5252685546875),
										float32(35.10642623901367),
									},
									[]interface{}{
										float32(-76.4208984375),
										float32(35.25907516479492),
									},
									[]interface{}{
										float32(-76.3385009765625),
										float32(35.294952392578125),
									},
									[]interface{}{
										float32(-76.0858154296875),
										float32(35.299434661865234),
									},
									[]interface{}{
										float32(-75.948486328125),
										float32(35.442771911621094),
									},
									[]interface{}{
										float32(-75.8660888671875),
										float32(35.53669738769531),
									},
									[]interface{}{
										float32(-75.772705078125),
										float32(35.5679817199707),
									},
									[]interface{}{
										float32(-75.706787109375),
										float32(35.63497543334961),
									},
									[]interface{}{
										float32(-75.706787109375),
										float32(35.7420539855957),
									},
									[]interface{}{
										float32(-75.6903076171875),
										float32(35.7420539855957),
									},
								},
							},
							[]interface{}{
								[]interface{}{
									[]interface{}{
										float32(-109.0283203125),
										float32(36.98500442504883),
									},
									[]interface{}{
										float32(-109.0283203125),
										float32(40.979896545410156),
									},
									[]interface{}{
										float32(-102.06298828125),
										float32(40.979896545410156),
									},
									[]interface{}{
										float32(-102.06298828125),
										float32(37.0025520324707),
									},
									[]interface{}{
										float32(-109.0283203125),
										float32(36.98500442504883),
									},
								},
							},
						},
					},
					"name": "geojson_multipolygon",
				},
				map[string]interface{}{
					"name":        "geojson_feature",
					"description": "Returns a GeoJSON Feature object",
					"data": map[string]interface{}{
						"type": "Feature",
						"geometry": map[string]interface{}{
							"type": "Polygon",
							"coordinates": []interface{}{
								[]interface{}{
									[]interface{}{
										float32(-80.7248764038086),
										float32(35.26545333862305),
									},
									[]interface{}{
										float32(-80.72135925292969),
										float32(35.267276763916016),
									},
									[]interface{}{
										float32(-80.71517944335938),
										float32(35.267696380615234),
									},
									[]interface{}{
										float32(-80.71251678466797),
										float32(35.27035903930664),
									},
									[]interface{}{
										float32(-80.70857238769531),
										float32(35.26825714111328),
									},
									[]interface{}{
										float32(-80.70479583740234),
										float32(35.26839828491211),
									},
									[]interface{}{
										float32(-80.7032470703125),
										float32(35.26503372192383),
									},
									[]interface{}{
										float32(-80.71089172363281),
										float32(35.25536346435547),
									},
									[]interface{}{
										float32(-80.71681213378906),
										float32(35.25536346435547),
									},
									[]interface{}{
										float32(-80.71509552001953),
										float32(35.26054763793945),
									},
									[]interface{}{
										float32(-80.71869659423828),
										float32(35.26026916503906),
									},
									[]interface{}{
										float32(-80.72032928466797),
										float32(35.2606201171875),
									},
									[]interface{}{
										float32(-80.72264862060547),
										float32(35.260337829589844),
									},
									[]interface{}{
										float32(-80.7248764038086),
										float32(35.26545333862305),
									},
								},
							},
						},
					},
				},
				map[string]interface{}{
					"data": map[string]interface{}{
						"type": "FeatureCollection",
						"features": []interface{}{
							map[string]interface{}{
								"geometry": map[string]interface{}{
									"coordinates": []interface{}{
										float32(-80.87088775634766),
										float32(35.215152740478516),
									},
									"type": "Point",
								},
							},
							map[string]interface{}{
								"geometry": map[string]interface{}{
									"coordinates": []interface{}{
										float32(-80.83775329589844),
										float32(35.24980163574219),
									},
									"type": "Point",
								},
							},
							map[string]interface{}{
								"geometry": map[string]interface{}{
									"coordinates": []interface{}{
										float32(-80.83827209472656),
										float32(35.25674819946289),
									},
									"type": "Point",
								},
							},
							map[string]interface{}{
								"geometry": map[string]interface{}{
									"coordinates": []interface{}{
										float32(-80.83697509765625),
										float32(35.25751876831055),
									},
									"type": "Point",
								},
							},
							map[string]interface{}{
								"geometry": map[string]interface{}{
									"coordinates": []interface{}{
										float32(-80.81647491455078),
										float32(35.401485443115234),
									},
									"type": "Point",
								},
							},
							map[string]interface{}{
								"geometry": map[string]interface{}{
									"coordinates": []interface{}{
										float32(-80.83556365966797),
										float32(35.399173736572266),
									},
									"type": "Point",
								},
							},
							map[string]interface{}{
								"geometry": map[string]interface{}{
									"type": "Polygon",
									"coordinates": []interface{}{
										[]interface{}{
											[]interface{}{
												float32(-80.72487831115721),
												float32(35.26545403190955),
											},
											[]interface{}{
												float32(-80.72135925292969),
												float32(35.26727607954368),
											},
											[]interface{}{
												float32(-80.71517944335938),
												float32(35.26769654625573),
											},
											[]interface{}{
												float32(-80.7125186920166),
												float32(35.27035945142482),
											},
											[]interface{}{
												float32(-80.70857048034668),
												float32(35.268257165144064),
											},
											[]interface{}{
												float32(-80.70479393005371),
												float32(35.268397319259996),
											},
											[]interface{}{
												float32(-80.70324897766113),
												float32(35.26503355355979),
											},
											[]interface{}{
												float32(-80.71088790893555),
												float32(35.2553619492954),
											},
											[]interface{}{
												float32(-80.71681022644043),
												float32(35.2553619492954),
											},
											[]interface{}{
												float32(-80.7150936126709),
												float32(35.26054831539319),
											},
											[]interface{}{
												float32(-80.71869850158691),
												float32(35.26026797976481),
											},
											[]interface{}{
												float32(-80.72032928466797),
												float32(35.26061839914875),
											},
											[]interface{}{
												float32(-80.72264671325684),
												float32(35.26033806376283),
											},
											[]interface{}{
												float32(-80.72487831115721),
												float32(35.26545403190955),
											},
										},
									},
								},
							},
						},
					},
					"name":        "geojson_feature_collection",
					"description": "Returns a GeoJSON FeatureCollection object",
				},
				map[string]interface{}{
					"name":        "geojson_geometry_collection",
					"description": "Returns a GeoJSON GeometryCollection object",
					"data": map[string]interface{}{
						"type": "GeometryCollection",
						"geometries": []interface{}{
							map[string]interface{}{
								"coordinates": []interface{}{
									float32(-80.66080474853516),
									float32(35.04939270019531),
								},
								"type": "Point",
							},
							map[string]interface{}{
								"type": "Polygon",
								"coordinates": []interface{}{
									[]interface{}{
										[]interface{}{
											float32(-80.66458129882812),
											float32(35.04496383666992),
										},
										[]interface{}{
											float32(-80.66344451904297),
											float32(35.04603576660156),
										},
										[]interface{}{
											float32(-80.66259002685547),
											float32(35.04558181762695),
										},
										[]interface{}{
											float32(-80.66387176513672),
											float32(35.044281005859375),
										},
										[]interface{}{
											float32(-80.66458129882812),
											float32(35.04496383666992),
										},
									},
								},
							},
							map[string]interface{}{
								"type": "LineString",
								"coordinates": []interface{}{
									[]interface{}{
										float32(-80.66236877441406),
										float32(35.05950927734375),
									},
									[]interface{}{
										float32(-80.6626968383789),
										float32(35.05926513671875),
									},
									[]interface{}{
										float32(-80.662841796875),
										float32(35.058929443359375),
									},
									[]interface{}{
										float32(-80.66307830810547),
										float32(35.05833435058594),
									},
									[]interface{}{
										float32(-80.6635971069336),
										float32(35.05775451660156),
									},
									[]interface{}{
										float32(-80.66387176513672),
										float32(35.057403564453125),
									},
									[]interface{}{
										float32(-80.66441345214844),
										float32(35.05703353881836),
									},
									[]interface{}{
										float32(-80.66486358642578),
										float32(35.056785583496094),
									},
									[]interface{}{
										float32(-80.66542053222656),
										float32(35.0565071105957),
									},
									[]interface{}{
										float32(-80.66563415527344),
										float32(35.056312561035156),
									},
									[]interface{}{
										float32(-80.66602325439453),
										float32(35.05589294433594),
									},
									[]interface{}{
										float32(-80.66619110107422),
										float32(35.055450439453125),
									},
									[]interface{}{
										float32(-80.66619110107422),
										float32(35.055171966552734),
									},
									[]interface{}{
										float32(-80.666259765625),
										float32(35.05488967895508),
									},
									[]interface{}{
										float32(-80.66621398925781),
										float32(35.054222106933594),
									},
									[]interface{}{
										float32(-80.66621398925781),
										float32(35.053924560546875),
									},
									[]interface{}{
										float32(-80.66595458984375),
										float32(35.05290603637695),
									},
									[]interface{}{
										float32(-80.66569519042969),
										float32(35.05204391479492),
									},
									[]interface{}{
										float32(-80.6655044555664),
										float32(35.051483154296875),
									},
									[]interface{}{
										float32(-80.66576385498047),
										float32(35.050479888916016),
									},
									[]interface{}{
										float32(-80.66616821289062),
										float32(35.04972457885742),
									},
									[]interface{}{
										float32(-80.66651153564453),
										float32(35.049285888671875),
									},
									[]interface{}{
										float32(-80.66692352294922),
										float32(35.04853057861328),
									},
									[]interface{}{
										float32(-80.66700744628906),
										float32(35.048213958740234),
									},
									[]interface{}{
										float32(-80.66706848144531),
										float32(35.04777526855469),
									},
									[]interface{}{
										float32(-80.66705322265625),
										float32(35.04738998413086),
									},
									[]interface{}{
										float32(-80.66696166992188),
										float32(35.0469856262207),
									},
									[]interface{}{
										float32(-80.66681671142578),
										float32(35.04635238647461),
									},
									[]interface{}{
										float32(-80.66659545898438),
										float32(35.04596710205078),
									},
									[]interface{}{
										float32(-80.6664047241211),
										float32(35.045616149902344),
									},
									[]interface{}{
										float32(-80.66600036621094),
										float32(35.04519271850586),
									},
									[]interface{}{
										float32(-80.66552734375),
										float32(35.04487609863281),
									},
									[]interface{}{
										float32(-80.66499328613281),
										float32(35.0445442199707),
									},
									[]interface{}{
										float32(-80.66449737548828),
										float32(35.04417419433594),
									},
									[]interface{}{
										float32(-80.66384887695312),
										float32(35.04387664794922),
									},
									[]interface{}{
										float32(-80.66304016113281),
										float32(35.04371643066406),
									},
								},
							},
						},
					},
				},
				map[string]interface{}{
					"description": "Returns a GeoJSON Point object with GeoJSON Named CRS",
					"name":        "geojson_crs_named_crs",
					"data": map[string]interface{}{
						"coordinates": []interface{}{
							float32(-105.0162124633789),
							float32(39.57421875),
						},
						"type": "Point",
					},
				},
				map[string]interface{}{
					"name":        "geojson_crs_linked_crs",
					"description": "Returns a GeoJSON Point object with GeoJSON Linked CRS",
					"data": map[string]interface{}{
						"type": "Point",
						"coordinates": []interface{}{
							float32(-105.0162124633789),
							float32(39.57421875),
						},
					},
				},
			},
		},
		Errors: nil,
	}

	allData := []interface{}{}
	for _, item := range testDataAll.Items {
		allData = append(allData, item.Data)
	}
	res := doGraphQL(TestGeoJSONSchema, query)

	assertGraphQLResultEqual(t, expected, res)
}

func TestPoint(t *testing.T) {
	query := `
	{
		point {
			name
			description
			data {
			   type
			   ... on GeoJSONPoint {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"point": map[string]interface{}{
				"name": "geojson_point",
				"data": map[string]interface{}{
					"coordinates": []interface{}{
						float32(-105.01621),
						float32(39.57422),
					},
					"type": "Point",
				},
				"description": "Returns a GeoJSON Point object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultEqual(t, expected, res)
}
func TestPoint_WrongCoordinates(t *testing.T) {
	query := `
	{
		point {
			name
			description
			data {
			   type
			   ... on GeoJSONPoint {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"point": map[string]interface{}{
				"name": "geojson_point",
				"data": map[string]interface{}{
					"coordinates": []interface{}{
						float32(66),
					},
					"type": "Point",
				},
				"description": "Returns a GeoJSON Point object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultNotEqual(t, expected, res)
}

func TestMultiPoint(t *testing.T) {
	query := `
	{
		multipoint {
			name
			description
			data {
			   type
			   ... on GeoJSONMultiPoint {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"multipoint": map[string]interface{}{
				"name": "geojson_multipoint",
				"data": map[string]interface{}{
					"coordinates": []interface{}{
						[]interface{}{
							float32(-105.01621),
							float32(39.57422),
						},
						[]interface{}{
							float32(-80.6665134),
							float32(35.0539943),
						},
					},
					"type": "MultiPoint",
				},
				"description": "Returns a GeoJSON MultiPoint object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultEqual(t, expected, res)
}
func TestMultiPoint_WrongCoordinatesStructure(t *testing.T) {
	query := `
	{
		multipoint {
			name
			description
			data {
			   type
			   ... on GeoJSONMultiPoint {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"multipoint": map[string]interface{}{
				"name": "geojson_multipoint",
				"data": map[string]interface{}{
					"coordinates": []interface{}{
						float32(-105.01621),
						float32(39.57422),
						float32(-80.6665134),
						float32(35.0539943),
					},
					"type": "MultiPoint",
				},
				"description": "Returns a GeoJSON MultiPoint object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultNotEqual(t, expected, res)
}

func TestLineString(t *testing.T) {
	query := `
	{
		linestring {
			name
			description
			data {
			   type
			   ... on GeoJSONLineString {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"linestring": map[string]interface{}{
				"name": "geojson_linestring",
				"data": map[string]interface{}{
					"type": "LineString",
					"coordinates": []interface{}{
						[]interface{}{
							float32(-101.744384765625),
							float32(39.32154846191406),
						},
						[]interface{}{
							float32(-101.5521240234375),
							float32(39.330047607421875),
						},
						[]interface{}{
							float32(-101.40380859375),
							float32(39.330047607421875),
						},
						[]interface{}{
							float32(-101.3323974609375),
							float32(39.36403274536133),
						},
						[]interface{}{
							float32(-101.041259765625),
							float32(39.36827850341797),
						},
						[]interface{}{
							float32(-100.975341796875),
							float32(39.30455017089844),
						},
						[]interface{}{
							float32(-100.9149169921875),
							float32(39.245018005371094),
						},
						[]interface{}{
							float32(-100.843505859375),
							float32(39.16414260864258),
						},
						[]interface{}{
							float32(-100.8050537109375),
							float32(39.104488372802734),
						},
						[]interface{}{
							float32(-100.491943359375),
							float32(39.10022735595703),
						},
						[]interface{}{
							float32(-100.43701171875),
							float32(39.09596252441406),
						},
						[]interface{}{
							float32(-100.338134765625),
							float32(39.09596252441406),
						},
						[]interface{}{
							float32(-100.1953125),
							float32(39.02771759033203),
						},
						[]interface{}{
							float32(-100.008544921875),
							float32(39.01064682006836),
						},
						[]interface{}{
							float32(-99.86572265625),
							float32(39.00210952758789),
						},
						[]interface{}{
							float32(-99.6844482421875),
							float32(38.97222137451172),
						},
						[]interface{}{
							float32(-99.51416015625),
							float32(38.929500579833984),
						},
						[]interface{}{
							float32(-99.38232421875),
							float32(38.920955657958984),
						},
						[]interface{}{
							float32(-99.3218994140625),
							float32(38.89530944824219),
						},
						[]interface{}{
							float32(-99.1131591796875),
							float32(38.869651794433594),
						},
						[]interface{}{
							float32(-99.0802001953125),
							float32(38.85681915283203),
						},
						[]interface{}{
							float32(-98.822021484375),
							float32(38.85681915283203),
						},
						[]interface{}{
							float32(-98.448486328125),
							float32(38.848262786865234),
						},
						[]interface{}{
							float32(-98.206787109375),
							float32(38.848262786865234),
						},
						[]interface{}{
							float32(-98.02001953125),
							float32(38.878204345703125),
						},
						[]interface{}{
							float32(-97.635498046875),
							float32(38.87392807006836),
						},
					},
				},
				"description": "Returns a GeoJSON LineString object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultEqual(t, expected, res)
}
func TestLineString_WrongCoordinatesStructure(t *testing.T) {
	query := `
	{
		linestring {
			name
			description
			data {
			   type
			   ... on GeoJSONLineString {
			   		coordinates
			   }
			}
		}
	}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"linestring": map[string]interface{}{
				"name": "geojson_linestring",
				"data": map[string]interface{}{
					"type": "LineString",
					"coordinates": []interface{}{
						float32(-105.01331329345703),
						float32(39.5744514465332),
						float32(-105.01621),
						float32(39.57422),
						float32(-80.6665134),
						float32(35.0539943),
					},
				},
				"description": "Returns a GeoJSON LineString object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultNotEqual(t, expected, res)
}

func TestMultiLineString(t *testing.T) {
	query := `
	{
		multilinestring {
			name
			description
			data {
			   type
			   ... on GeoJSONMultiLineString {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"multilinestring": map[string]interface{}{
				"name": "geojson_multilinestring",
				"data": map[string]interface{}{
					"type": "MultiLineString",
					"coordinates": []interface{}{
						[]interface{}{
							[]interface{}{
								float32(-105.02144622802734),
								float32(39.57805633544922),
							},
							[]interface{}{
								float32(-105.0215072631836),
								float32(39.57780838012695),
							},
							[]interface{}{
								float32(-105.02157592773438),
								float32(39.57749557495117),
							},
							[]interface{}{
								float32(-105.02157592773438),
								float32(39.57716369628906),
							},
							[]interface{}{
								float32(-105.02157592773438),
								float32(39.57703399658203),
							},
							[]interface{}{
								float32(-105.02153015136719),
								float32(39.5767822265625),
							},
						},
						[]interface{}{
							[]interface{}{
								float32(-105.0198974609375),
								float32(39.57499694824219),
							},
							[]interface{}{
								float32(-105.01959991455078),
								float32(39.57489776611328),
							},
							[]interface{}{
								float32(-105.01905822753906),
								float32(39.57478332519531),
							},
						},
						[]interface{}{
							[]interface{}{
								float32(-105.01717376708984),
								float32(39.57440185546875),
							},
							[]interface{}{
								float32(-105.01698303222656),
								float32(39.57438659667969),
							},
							[]interface{}{
								float32(-105.01663970947266),
								float32(39.57438659667969),
							},
							[]interface{}{
								float32(-105.01651000976562),
								float32(39.57440185546875),
							},
							[]interface{}{
								float32(-105.01595306396484),
								float32(39.57426834106445),
							},
						},
						[]interface{}{
							[]interface{}{
								float32(-105.01427459716797),
								float32(39.573970794677734),
							},
							[]interface{}{
								float32(-105.01412963867188),
								float32(39.574039459228516),
							},
							[]interface{}{
								float32(-105.01382446289062),
								float32(39.57416915893555),
							},
							[]interface{}{
								float32(-105.01331329345703),
								float32(39.5744514465332),
							},
						},
					},
				},
				"description": "Returns a GeoJSON MultiLineString object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultEqual(t, expected, res)
}
func TestMultiLineString_WrongCoordinatesStructure(t *testing.T) {
	query := `
	{
		multilinestring {
			name
			description
			data {
			   type
			   ... on GeoJSONMultiLineString {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"multilinestring": map[string]interface{}{
				"name": "geojson_multilinestring",
				"data": map[string]interface{}{
					"type": "MultiLineString",
					"coordinates": []interface{}{
						float32(-105.01331329345703),
						float32(39.5744514465332),
						float32(-105.01621),
						float32(39.57422),
						float32(-80.6665134),
						float32(35.0539943),
					},
				},
				"description": "Returns a GeoJSON MultiLineString object",
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultNotEqual(t, expected, res)
}

func TestCRS_NamedCRS(t *testing.T) {
	query := `
	{
		crs_named_crs {
			name
			description
			data {
			   type
			   crs {
			   		type
			   		properties {
			   			... on GeoJSONNamedCRSProperties {
			   				name
			   			}
			   			... on GeoJSONLinkedCRSProperties {
			   				href
			   				type
			   			}
			   		}
			   }
			   ... on GeoJSONPoint {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"crs_named_crs": map[string]interface{}{
				"name":        "geojson_crs_named_crs",
				"description": "Returns a GeoJSON Point object with GeoJSON Named CRS",
				"data": map[string]interface{}{
					"type": "Point",
					"coordinates": []interface{}{
						float32(-105.0162124633789),
						float32(39.57421875),
					},
					"crs": map[string]interface{}{
						"properties": map[string]interface{}{
							"name": "urn:ogc:def:crs:OGC:1.3:CRS84",
						},
						"type": "name",
					},
				},
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultEqual(t, expected, res)
}
func TestCRS_LinkedCRS(t *testing.T) {
	query := `
	{
		crs_linked_crs {
			name
			description
			data {
			   type
			   crs {
			   		type
			   		properties {
			   			... on GeoJSONNamedCRSProperties {
			   				name
			   			}
			   			... on GeoJSONLinkedCRSProperties {
			   				href
			   				type
			   			}
			   		}
			   }
			   ... on GeoJSONPoint {
			   		coordinates
			   }
			}
		}
	}
	`
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"crs_linked_crs": map[string]interface{}{
				"name":        "geojson_crs_linked_crs",
				"description": "Returns a GeoJSON Point object with GeoJSON Linked CRS",
				"data": map[string]interface{}{
					"type": "Point",
					"crs": map[string]interface{}{
						"type": "link",
						"properties": map[string]interface{}{
							"href": "http://example.com/crs/42",
							"type": "proj4",
						},
					},
					"coordinates": []interface{}{
						float32(-105.0162124633789),
						float32(39.57421875),
					},
				},
			},
		},
		Errors: nil,
	}
	res := doGraphQL(TestGeoJSONSchema, query)
	assertGraphQLResultEqual(t, expected, res)
}
