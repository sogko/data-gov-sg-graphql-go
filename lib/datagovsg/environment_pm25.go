package datagovsg

type PM25ReadingsOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
	Date     string `json:"date,omitempy" url:"date,omitempy"`
}

type PM25Readings struct {
	Value int  `json:"value,omitempty"`
	Area  Area `json:"area,omitempty"`
}

type PM25ReadingRegions struct {
	South   int `json:"south,omitempty"`
	North   int `json:"north,omitempty"`
	East    int `json:"east,omitempty"`
	Central int `json:"central,omitempty"`
	West    int `json:"west,omitempty"`
}
type PM25ReadingIntervals struct {
	PM25OneHourly PM25ReadingRegions `json:"pm25_one_hourly,omitempty"`
}
type PM25ReadingsResultItem struct {
	UpdateTimestamp string               `json:"update_timestamp,omitempty"`
	Timestamp       string               `json:"timestamp,omitempty"`
	Readings        PM25ReadingIntervals `json:"readings,omitempty"`
}

type PM25ReadingsResult struct {
	APIInfo        APIInfo                  `json:"api_info,omitempty"`
	RegionMetadata []Area                   `json:"region_metadata,omitempty"`
	Items          []PM25ReadingsResultItem `json:"items,omitempty"`
}

func (resp *PM25ReadingsResult) AreaByName(name string) Area {
	for _, area := range resp.RegionMetadata {
		if area.Name == name {
			return area
		}
	}
	return Area{}
}

func (resp *PM25ReadingsResult) ToGraphQL() interface{} {

	items := []PM25ReadingsResultItemGraphQL{}
	for _, i := range resp.Items {
		item := PM25ReadingsResultItemGraphQL{
			UpdateTimestamp: i.UpdateTimestamp,
			Timestamp:       i.Timestamp,
			Readings: PM25ReadingIntervalsGraphQL{
				PM25OneHourly: PM25ReadingRegionsGraphQL{
					South: PM25Readings{
						Value: i.Readings.PM25OneHourly.South,
						Area:  resp.AreaByName("south"),
					},
					North: PM25Readings{
						Value: i.Readings.PM25OneHourly.North,
						Area:  resp.AreaByName("north"),
					},
					East: PM25Readings{
						Value: i.Readings.PM25OneHourly.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PM25Readings{
						Value: i.Readings.PM25OneHourly.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PM25Readings{
						Value: i.Readings.PM25OneHourly.West,
						Area:  resp.AreaByName("west"),
					},
				},
			},
		}
		items = append(items, item)
	}
	return PM25ReadingsResultGraphQL{
		APIInfo: resp.APIInfo,
		Items:   items,
	}
}

type PM25ReadingRegionsGraphQL struct {
	South   PM25Readings `json:"south,omitempty"`
	North   PM25Readings `json:"north,omitempty"`
	East    PM25Readings `json:"east,omitempty"`
	Central PM25Readings `json:"central,omitempty"`
	West    PM25Readings `json:"west,omitempty"`
}
type PM25ReadingIntervalsGraphQL struct {
	PM25OneHourly PM25ReadingRegionsGraphQL `json:"pm25_one_hourly,omitempty"`
}
type PM25ReadingsResultItemGraphQL struct {
	UpdateTimestamp string                      `json:"update_timestamp,omitempty"`
	Timestamp       string                      `json:"timestamp,omitempty"`
	Readings        PM25ReadingIntervalsGraphQL `json:"readings,omitempty"`
}

type PM25ReadingsResultGraphQL struct {
	APIInfo APIInfo                         `json:"api_info,omitempty"`
	Items   []PM25ReadingsResultItemGraphQL `json:"items,omitempty"`
}
