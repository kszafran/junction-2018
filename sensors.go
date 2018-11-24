package main

import (
	"encoding/json"
	"errors"
	"github.com/ssimunic/gosensors"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"
)

var readingRegex = regexp.MustCompile(`(\s*[+-]?[\d\s.]+[^(\s]+)(\s+\(.*\))?`)

// MAC -> sensor data
var sensorHistory = make(map[string][]SensorEntry)

// adapter -> values
type SensorEntry map[string]TsReadings

func GetSensorEntries(ip string) []SensorEntry {
	return sensorHistory[ip]
}

func StoreSensorData(mac string, info io.Reader) error {
	infostring, err := ioutil.ReadAll(info)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Got sensor information from %s: %s\n", mac, infostring)

	var sensors gosensors.Sensors
	err = json.NewDecoder(strings.NewReader(string(infostring))).Decode(&sensors)
	if err != nil {
		return err
	}
	if len(sensors.Chips) == 0 {
		return errors.New("no sensor information found")
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
	log.Printf("[INFO] Registered sensor reading from %s: %v\n", mac, string(bytes))

	sensorHistory[mac] = append(sensorHistory[mac], readings)
	return nil
}

