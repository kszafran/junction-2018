package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kszafran/junction-2018/models"
	"io/ioutil"
	"log"
)

func main() {
	refreshToken()
	r := gin.Default()
	r.GET("/test", testHandler)
	r.GET("/path-trace", pathTraceHandler)
	r.POST("/sensor/:ip", sensorHandler)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func testHandler(c *gin.Context) {
	var health models.GetOverallNetworkHealthResponseResponse
	err := get("/network-health?timestamp", H{"__runsync": "true", "__timeout": "60", "__persistbapioutput": "true"}, &health)
	if err != nil {
		c.String(500, "failed to get health: %v", err)
		c.Error(err)
		return
	}
	c.JSON(200, health)
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
	err := post("/flow-analysis", nil, &result, &models.FlowAnalysisRequest{
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
		err = get("/flow-analysis/"+result.Response.FlowAnalysisID, nil, &trace)
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
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(500, "failed to read request body: %v", err)
		c.Error(err)
		return
	}
	log.Printf("[INFO] Got sensor information from %s: %s\n", c.Param("ip"), body)
}
