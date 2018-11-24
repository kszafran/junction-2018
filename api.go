package main

type SensorData struct {
	Name        string       `json:"name,omitempty"`
	IP          string       `json:"ip,omitempty"`
	Type        string       `json:"type,omitempty"`
	Current     []Reading    `json:"current,omitempty"`
	History     []TsReadings `json:"data_history,omitempty"`
	CiscoHealth CiscoHealth  `json:"cisco_health_data"`
}

type TsReadings struct {
	Date int64     `json:"date"`
	Data []Reading `json:"data,omitempty"`
}

type Reading struct {
	Name  string `json:"name"`
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
