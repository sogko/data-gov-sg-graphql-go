package datagovsg

type PSIReadingsOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
	Date     string `json:"date,omitempy" url:"date,omitempy"`
}

type PSIReadings struct {
	Value float32 `json:"value,omitempty"`
	Area  Area    `json:"area,omitempty"`
}

type PSIReadingRegions struct {
	National float32 `json:"national,omitempty"`
	South    float32 `json:"south,omitempty"`
	North    float32 `json:"north,omitempty"`
	East     float32 `json:"east,omitempty"`
	Central  float32 `json:"central,omitempty"`
	West     float32 `json:"west,omitempty"`
}

type PSIReadingIntervals struct {
	PSITwentyFourHourly  PSIReadingRegions `json:"psi_twenty_four_hourly,omitempty"`
	PM10TwentyFourHourly PSIReadingRegions `json:"pm10_twenty_four_hourly,omitempty"`
	PM10SubIndex         PSIReadingRegions `json:"pm10_sub_index,omitempty"`
	PM25TwentyFourHourly PSIReadingRegions `json:"pm25_twenty_four_hourly,omitempty"`
	PSIThreeHourly       PSIReadingRegions `json:"psi_three_hourly,omitempty"`
	SO2TwentyFourHourly  PSIReadingRegions `json:"so2_twenty_four_hourly,omitempty"`
	O3SubIndex           PSIReadingRegions `json:"o3_sub_index,omitempty"`
	NO2OneHourMax        PSIReadingRegions `json:"no2_one_hour_max,omitempty"`
	SO2SubIndex          PSIReadingRegions `json:"so2_sub_index,omitempty"`
	PM2SubIndex          PSIReadingRegions `json:"pm25_sub_index,omitempty"`
	COEightHourMax       PSIReadingRegions `json:"co_eight_hour_max,omitempty"`
	COSubIndex           PSIReadingRegions `json:"co_sub_index,omitempty"`
	O3EightHourMax       PSIReadingRegions `json:"o3_eight_hour_max,omitempty"`
}

type PSIReadingsResultItem struct {
	UpdateTimestamp string              `json:"update_timestamp,omitempty"`
	Timestamp       string              `json:"timestamp,omitempty"`
	Readings        PSIReadingIntervals `json:"readings,omitempty"`
}

type PSIReadingsResult struct {
	APIInfo        APIInfo                 `json:"api_info,omitempty"`
	RegionMetadata []Area                  `json:"region_metadata,omitempty"`
	Items          []PSIReadingsResultItem `json:"items,omitempty"`
}

func (resp *PSIReadingsResult) AreaByName(name string) Area {
	for _, area := range resp.RegionMetadata {
		if area.Name == name {
			return area
		}
	}
	return Area{}
}

func (resp *PSIReadingsResult) ToGraphQL() interface{} {

	// I know this is extremely ripe for some refactoring work, but patience, my young grasshoppa
	items := []PSIReadingsResultItemGraphQL{}
	for _, i := range resp.Items {
		item := PSIReadingsResultItemGraphQL{
			UpdateTimestamp: i.UpdateTimestamp,
			Timestamp:       i.Timestamp,
			Readings: PSIReadingIntervalsGraphQL{
				PSITwentyFourHourly: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.PSITwentyFourHourly.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.PSITwentyFourHourly.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.PSITwentyFourHourly.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.PSITwentyFourHourly.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.PSITwentyFourHourly.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.PSITwentyFourHourly.West,
						Area:  resp.AreaByName("west"),
					},
				},
				PM10TwentyFourHourly: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.PM10TwentyFourHourly.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.PM10TwentyFourHourly.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.PM10TwentyFourHourly.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.PM10TwentyFourHourly.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.PM10TwentyFourHourly.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.PM10TwentyFourHourly.West,
						Area:  resp.AreaByName("west"),
					},
				},
				PM10SubIndex: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.PM10SubIndex.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.PM10SubIndex.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.PM10SubIndex.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.PM10SubIndex.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.PM10SubIndex.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.PM10SubIndex.West,
						Area:  resp.AreaByName("west"),
					},
				},
				PM25TwentyFourHourly: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.PM25TwentyFourHourly.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.PM25TwentyFourHourly.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.PM25TwentyFourHourly.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.PM25TwentyFourHourly.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.PM25TwentyFourHourly.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.PM25TwentyFourHourly.West,
						Area:  resp.AreaByName("west"),
					},
				},
				PSIThreeHourly: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.PSIThreeHourly.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.PSIThreeHourly.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.PSIThreeHourly.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.PSIThreeHourly.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.PSIThreeHourly.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.PSIThreeHourly.West,
						Area:  resp.AreaByName("west"),
					},
				},
				SO2TwentyFourHourly: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.SO2TwentyFourHourly.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.SO2TwentyFourHourly.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.SO2TwentyFourHourly.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.SO2TwentyFourHourly.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.SO2TwentyFourHourly.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.SO2TwentyFourHourly.West,
						Area:  resp.AreaByName("west"),
					},
				},
				O3SubIndex: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.O3SubIndex.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.O3SubIndex.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.O3SubIndex.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.O3SubIndex.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.O3SubIndex.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.O3SubIndex.West,
						Area:  resp.AreaByName("west"),
					},
				},
				NO2OneHourMax: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.NO2OneHourMax.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.NO2OneHourMax.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.NO2OneHourMax.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.NO2OneHourMax.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.NO2OneHourMax.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.NO2OneHourMax.West,
						Area:  resp.AreaByName("west"),
					},
				},
				SO2SubIndex: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.SO2SubIndex.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.SO2SubIndex.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.SO2SubIndex.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.SO2SubIndex.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.SO2SubIndex.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.SO2SubIndex.West,
						Area:  resp.AreaByName("west"),
					},
				},
				PM2SubIndex: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.PM2SubIndex.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.PM2SubIndex.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.PM2SubIndex.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.PM2SubIndex.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.PM2SubIndex.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.PM2SubIndex.West,
						Area:  resp.AreaByName("west"),
					},
				},
				COEightHourMax: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.COEightHourMax.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.COEightHourMax.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.COEightHourMax.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.COEightHourMax.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.COEightHourMax.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.COEightHourMax.West,
						Area:  resp.AreaByName("west"),
					},
				},
				COSubIndex: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.COSubIndex.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.COSubIndex.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.COSubIndex.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.COSubIndex.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.COSubIndex.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.COSubIndex.West,
						Area:  resp.AreaByName("west"),
					},
				},
				O3EightHourMax: PSIReadingRegionsGraphQL{
					National: PSIReadings{
						Value: i.Readings.O3EightHourMax.National,
						Area:  resp.AreaByName("national"),
					},
					South: PSIReadings{
						Value: i.Readings.O3EightHourMax.South,
						Area:  resp.AreaByName("south"),
					},
					North: PSIReadings{
						Value: i.Readings.O3EightHourMax.North,
						Area:  resp.AreaByName("north"),
					},
					East: PSIReadings{
						Value: i.Readings.O3EightHourMax.East,
						Area:  resp.AreaByName("east"),
					},
					Central: PSIReadings{
						Value: i.Readings.O3EightHourMax.Central,
						Area:  resp.AreaByName("central"),
					},
					West: PSIReadings{
						Value: i.Readings.O3EightHourMax.West,
						Area:  resp.AreaByName("west"),
					},
				},
			},
		}
		items = append(items, item)
	}
	return PSIReadingsResultGraphQL{
		APIInfo: resp.APIInfo,
		Items:   items,
	}
}

type PSIReadingRegionsGraphQL struct {
	National PSIReadings `json:"national,omitempty"`
	South    PSIReadings `json:"south,omitempty"`
	North    PSIReadings `json:"north,omitempty"`
	East     PSIReadings `json:"east,omitempty"`
	Central  PSIReadings `json:"central,omitempty"`
	West     PSIReadings `json:"west,omitempty"`
}
type PSIReadingIntervalsGraphQL struct {
	PSITwentyFourHourly  PSIReadingRegionsGraphQL `json:"psi_twenty_four_hourly,omitempty"`
	PM10TwentyFourHourly PSIReadingRegionsGraphQL `json:"pm10_twenty_four_hourly,omitempty"`
	PM10SubIndex         PSIReadingRegionsGraphQL `json:"pm10_sub_index,omitempty"`
	PM25TwentyFourHourly PSIReadingRegionsGraphQL `json:"pm25_twenty_four_hourly,omitempty"`
	PSIThreeHourly       PSIReadingRegionsGraphQL `json:"psi_three_hourly,omitempty"`
	SO2TwentyFourHourly  PSIReadingRegionsGraphQL `json:"so2_twenty_four_hourly,omitempty"`
	O3SubIndex           PSIReadingRegionsGraphQL `json:"o3_sub_index,omitempty"`
	NO2OneHourMax        PSIReadingRegionsGraphQL `json:"no2_one_hour_max,omitempty"`
	SO2SubIndex          PSIReadingRegionsGraphQL `json:"so2_sub_index,omitempty"`
	PM2SubIndex          PSIReadingRegionsGraphQL `json:"pm25_sub_index,omitempty"`
	COEightHourMax       PSIReadingRegionsGraphQL `json:"co_eight_hour_max,omitempty"`
	COSubIndex           PSIReadingRegionsGraphQL `json:"co_sub_index,omitempty"`
	O3EightHourMax       PSIReadingRegionsGraphQL `json:"o3_eight_hour_max,omitempty"`
}
type PSIReadingsResultItemGraphQL struct {
	UpdateTimestamp string                     `json:"update_timestamp,omitempty"`
	Timestamp       string                     `json:"timestamp,omitempty"`
	Readings        PSIReadingIntervalsGraphQL `json:"readings,omitempty"`
}

type PSIReadingsResultGraphQL struct {
	APIInfo APIInfo                        `json:"api_info,omitempty"`
	Items   []PSIReadingsResultItemGraphQL `json:"items,omitempty"`
}
