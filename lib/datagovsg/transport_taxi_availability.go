package datagovsg

import (
	"encoding/json"
)

type TaxiAvailabilityOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
}

type TaxiAvailabilityResultProperties struct {
	Timestamp string  `json:"timestamp,omitempty"`
	TaxiCount int     `json:"taxi_count,omitempty"`
	APIInfo   APIInfo `json:"api_info,omitempty"`
}
type TaxiAvailabilityResultItem struct {
	Type       string                           `json:"type,omitempty"`
	Geometry   interface{}                      `json:"geometry,omitempty"`
	Properties TaxiAvailabilityResultProperties `json:"properties,omitempty"`
}

// TaxiAvailabilityResult interim struct until we model GeoJSON struct
type TaxiAvailabilityResult struct {
	Type     string                       `json:"type,omitempty"`
	CRS      interface{}                  `json:"crs,omitempty"`
	Features []TaxiAvailabilityResultItem `json:"features,omitempty"`
}

func (resp *TaxiAvailabilityResult) ToGraphQL() interface{} {
	if resp == nil {
		return TaxiAvailabilityResultGraphQL{}
	}
	// This assumes that /transport/taxi-availability always returns
	// a FeatureCollection with a single Feature of MultiPoint
	feature := TaxiAvailabilityResultItem{}
	if len(resp.Features) > 0 {
		feature = resp.Features[0]
	}

	// marshal-unmarshal to map[string]interface{}
	geoJSONMap := map[string]interface{}{}
	b, _ := json.Marshal(resp)
	json.Unmarshal(b, &geoJSONMap)

	return TaxiAvailabilityResultGraphQL{
		Timestamp: feature.Properties.Timestamp,
		TaxiCount: feature.Properties.TaxiCount,
		APIInfo:   feature.Properties.APIInfo,
		Result:    geoJSONMap,
	}
}

type TaxiAvailabilityResultGraphQL struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	TaxiCount int                    `json:"taxi_count,omitempty"`
	APIInfo   APIInfo                `json:"api_info,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty"`
}
