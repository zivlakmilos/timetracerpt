package timetrace

import (
	"encoding/json"
	"os"
	"time"
)

type Record struct {
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Project struct {
		Key string `json:"key"`
	} `json:"project"`
	IsBillable bool     `json:"is_billable"`
	Tags       []string `json:"tags"`
}

func (r *Record) UnmarshalJSON(data []byte) error {
	type Alias Record
	aux := &struct {
		Start string `json:"start"`
		End   string `json:"end"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	r.Start, err = time.Parse(time.RFC3339, aux.Start)
	if err != nil {
		return err
	}
	if aux.End != "" {
		r.End, err = time.Parse(time.RFC3339, aux.End)
		if err != nil {
			return err
		}
	} else {
		r.End = time.Time{}
	}
	return nil
}

func ParseRecord(filePath string) (*Record, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rec Record
	if err := json.Unmarshal(data, &rec); err != nil {
		return nil, err
	}

	return &rec, nil
}
