package main

import (
	"encoding/json"
	"github.com/ssimunic/gosensors"
	"io"
	"log"
	"regexp"
	"sort"
	"time"
)

var readingRegex = regexp.MustCompile(`(\s*[+-]?[\d\s.]+[^(\s]+)(\s+\(.*\))?`)

// ip -> sensor data
var sensorHistory = make(map[string][]SensorEntry)

// adapter -> values
type SensorEntry map[string]TsReadings

func GetSensorEntries(ip string) []SensorEntry {
	return sensorHistory[ip]
}

func StoreSensorData(ip string, info io.Reader) error {
	log.Printf("[INFO] Got sensor information from %s\n", ip)

	var sensors gosensors.Sensors
	err := json.NewDecoder(info).Decode(&sensors)
	if err != nil {
		return err
	}

	date := time.Now().Unix()

	readings := make(SensorEntry)
	for _, entries := range sensors.Chips {
		if _, ok := entries["Adapter"]; !ok {
			continue
		}
		adapter := entries["Adapter"]
		delete(entries, "Adapter")

		hr := readings[adapter]
		if hr.Date == 0 {
			hr = TsReadings{Date: date}
		}

		for name, value := range entries {
			submatch := readingRegex.FindSubmatch([]byte(value))
			if len(submatch) <= 1 {
				continue
			}
			hr.Data = append(hr.Data, Reading{
				Name: name,
				Value: string(submatch[1]),
			})
		}

		sort.Slice(hr.Data, func(i, j int) bool {
			return hr.Data[i].Name < hr.Data[j].Name
		})

		readings[adapter] = hr
	}

	bytes, _ := json.Marshal(readings)
	log.Printf("[INFO] Registered sensor reading: %v\n", string(bytes))

	sensorHistory[ip] = append(sensorHistory[ip], readings)
	return nil
}

