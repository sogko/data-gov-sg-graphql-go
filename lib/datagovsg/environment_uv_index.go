package datagovsg

type UVIndexOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
	Date     string `json:"date,omitempy" url:"date,omitempy"`
}

type UVIndexReading struct {
	Value     int    `json:"value,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type UVIndexReadingsResultItem struct {
	UpdateTimestamp string           `json:"update_timestamp,omitempty"`
	Timestamp       string           `json:"timestamp,omitempty"`
	Index           []UVIndexReading `json:"index,omitempty"`
}

type UVIndexReadingsResult struct {
	APIInfo APIInfo                     `json:"api_info,omitempty"`
	Items   []UVIndexReadingsResultItem `json:"items,omitempty"`
}

func (resp *UVIndexReadingsResult) ToGraphQL() interface{} {
	return resp
}
