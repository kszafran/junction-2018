package main

import "github.com/kszafran/junction-2018/models"

type CiscoHealth *models.GetClientDetailResponseResponseTopologyNodesItems0

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
