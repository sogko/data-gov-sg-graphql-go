package datagovsg

type TrafficImagesOptions struct {
	DateTime string `json:"date_time,omitempy" url:"date_time,omitempy"`
}

type TrafficImageMetadata struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	MD5    string `json:"md5,omitempty"`
}

type TrafficImageCamera struct {
	Timestamp     string               `json:"timestamp,omitempty"`
	Image         string               `json:"image,omitempty"`
	Location      Location             `json:"location,omitempty"`
	CameraID      int                  `json:"camera_id,omitempty"`
	ImageID       int                  `json:"image_id,omitempty"`
	ImageMetadata TrafficImageMetadata `json:"image_metadata,omitempty"`
}

type TrafficImagesResultItem struct {
	Timestamp string               `json:"timestamp,omitempty"`
	Cameras   []TrafficImageCamera `json:"cameras,omitempty"`
}

type TrafficImagesResult struct {
	APIInfo APIInfo                   `json:"api_info,omitempty"`
	Items   []TrafficImagesResultItem `json:"items,omitempty"`
}

func (resp *TrafficImagesResult) ToGraphQL() interface{} {
	return resp
}
