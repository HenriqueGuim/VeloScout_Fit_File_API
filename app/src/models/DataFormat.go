package models

import "time"

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
