package models

import "time"

// DataFormat is a struct that contains the data that will be sent to the frontend
// The json tags are used to make the json response more readable
// The omitempty tag is used to not send empty values
type DataFormat struct {
	Date           time.Time
	Duration       uint32            `json:"duration,omitempty"`
	Distance       float64           `json:"distance,omitempty"`
	ElevationGain  uint16            `json:"elevation_gain,omitempty"`
	TSS            uint16            `json:"t_ss,omitempty"`
	IF             float32           `json:"i_f,omitempty"`
	HeartRate      map[string]uint8  `json:"heart_rate,omitempty"`
	Power          map[string]uint16 `json:"power,omitempty"`
	AverageCadence uint8             `json:"average_cadence,omitempty"`
}

func NewDataFormat() *DataFormat {
	dataformat := DataFormat{HeartRate: make(map[string]uint8), Power: make(map[string]uint16)}
	return &dataformat
}
