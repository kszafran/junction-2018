package main

import (
	"encoding/json"
	"errors"
	"github.com/ssimunic/gosensors"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

var readingRegex = regexp.MustCompile(`(\s*[+-]?[\d\s.]+[^(\s]+)(\s+\(.*\))?`)

// MAC -> sensor data
var sensorReadings = make(map[string]SensorReadings)

func GetSensorReadings(ip string) SensorReadings {
	return sensorReadings[ip]
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
	readings := sensorReadings[mac]
	if readings == nil {
		readings = SensorReadings{}
		sensorReadings[mac] = readings
	}

	for _, entries := range sensors.Chips {
		if _, ok := entries["Adapter"]; !ok {
			continue
		}
		adapter := entries["Adapter"]
		delete(entries, "Adapter")

		for name, value := range entries {
			submatch := readingRegex.FindSubmatch([]byte(value))
			if len(submatch) <= 1 {
				continue
			}
			fullName := adapter + " " + name
			readings[fullName] = append(readings[fullName], Reading{
				Date: date,
				Value: string(submatch[1]),
			})
		}
	}
	log.Printf("[INFO] Successfully registered sensor reading from %s\n", mac)
	return nil
}

