package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kszafran/junction-2018/models"
	"log"
	"net/http"
)

type H map[string]string

const (
	host = "82.203.192.245"
	user = "developer"
	pass = "Hackjunction2018"
)

// TODO refresh token
var token string

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		err := authenticate()
		if err != nil {
			c.AbortWithError(500, errors.New(fmt.Sprintf("failed to authenticate: %v", err)))
			return
		}
		var health models.GetOverallNetworkHealthResponseResponse
		err = get("/network-health?timestamp", H{"__runsync": "true", "__timeout": "60", "__persistbapioutput": "true"}, &health)
		if err != nil {
			c.AbortWithError(500, errors.New(fmt.Sprintf("failed to get health: %v", err)))
			return
		}
		c.JSON(200, health)
	})
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func authenticate() error {
	req, err := http.NewRequest("POST", "https://"+host+"/dna/system/api/v1/auth/token", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(user+":"+pass)))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var auth models.AuthenticationAPIResponse
	err = json.NewDecoder(res.Body).Decode(&auth)
	if err != nil {
		return err
	}
	token = auth.Token
	return nil
}

func get(path string, headers H, response interface{}) error {
	req, err := http.NewRequest("GET", "https://"+host+"/dna/intent/api/v1"+path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Auth-Token", token)
	for name, value := range headers {
		req.Header.Add(name, value)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&response)
}
