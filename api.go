package main

type ClientHealth struct {
	Name        string         `json:"name,omitempty"`
	MAC         string         `json:"mac,omitempty"`
	Type        string         `json:"type,omitempty"`
	Sensors     SensorReadings `json:"data_history,omitempty"`
	CiscoHealth CiscoHealth    `json:"cisco_health_data"`
}

type SensorReadings map[string][]Reading

type Reading struct {
	Date  int64  `json:"date"`
	Value string `json:"value"`
}

type CiscoHealth struct {
	Current int             `json:"current_health_data"`
	History []HistoryHealth `json:"history_health_data,omitempty"`
}

type HistoryHealth struct {
	Date int64 `json:"date"`
	Data int   `json:"data"`
}
