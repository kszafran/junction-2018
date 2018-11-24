package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kszafran/junction-2018/models"
	"log"
)

func main() {
	KeepTokenFresh()
	r := gin.Default()
	r.GET("/test", testHandler)
	r.GET("/topology", topologyHandler)
	r.GET("/device-health/:ip", deviceHealthHandler)
	r.GET("/path-trace", pathTraceHandler)
	r.POST("/sensor/:ip", sensorHandler)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func testHandler(c *gin.Context) {
	var health models.GetOverallNetworkHealthResponseResponse
	err := Get("/network-health?timestamp", H{"__runsync": "true", "__timeout": "60", "__persistbapioutput": "true"}, &health)
	if err != nil {
		c.String(500, "failed to get health: %v", err)
		c.Error(err)
		return
	}
	c.JSON(200, health)
}

func deviceHealthHandler(c *gin.Context) {
	ip := c.Param("ip")
	entries := GetSensorEntries(ip)
	respMap := make(map[string]*SensorData)
	for i := len(entries) - 1; i >= 0; i-- {
		entry := entries[i]
		for name, value := range entry {
			if _, ok := respMap[name]; !ok {
				respMap[name] = &SensorData{Name: ip, IP: ip, Type: name, Current: value.Data}
			}
			respMap[name].History = append(respMap[name].History, value)
		}
	}
	var resp []*SensorData
	for _, data := range respMap {
		resp = append(resp, data)
	}
	c.JSON(200, resp)
}

func topologyHandler(c *gin.Context) {
	var topology models.TopologyResult
	err := Get("/topology/physical-topology", nil, &topology)
	if err != nil {
		c.String(500, "failed to get topology: %v", err)
		c.Error(err)
		return
	}
	c.JSON(200, topology.Response)
}

func pathTraceHandler(c *gin.Context) {
	source := c.Query("source")
	dest := c.Query("dest")
	protocol := c.Query("protocol")
	if source == "" {
		c.String(400, "'source' query param missing")
		return
	}
	if dest == "" {
		c.String(400, "'target' query param missing")
		return
	}
	if protocol == "" {
		protocol = "TCP"
	}
	var result models.FlowAnalysisRequestResultOutput
	err := Post("/flow-analysis", nil, &result, &models.FlowAnalysisRequest{
		Protocol:   protocol,
		SourceIP:   source,
		SourcePort: "80",
		DestIP:     dest,
		DestPort:   "80",
	})
	if err != nil {
		c.String(500, "failed to initiate path trace: %v", err)
		c.Error(err)
		return
	}
	var trace models.PathResponseResult
	status := "INPROGRESS"
	for status == "INPROGRESS" {
		err = Get("/flow-analysis/"+result.Response.FlowAnalysisID, nil, &trace)
		if err != nil {
			c.String(500, "failed to get path trace results: %v", err)
			c.Error(err)
			return
		}
		status = trace.Response.Request.Status
	}
	if status != "COMPLETED" {
		c.String(500, "path trace %s: %s", status, trace.Response.Request.FailureReason)
		return
	}
	c.JSON(200, trace.Response.NetworkElementsInfo)
}

func sensorHandler(c *gin.Context) {
	ip := c.Param("ip")
	log.Printf("[INFO] Got sensor information from %s\n", ip)
	err := StoreSensorData(ip, c.Request.Body)
	if err != nil {
		c.String(500, "failed to read decode request body: %v", err)
		c.Error(err)
		return
	}
	c.Status(204)
}
