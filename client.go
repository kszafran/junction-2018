package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kszafran/junction-2018/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type H map[string]string

const (
	host = "82.203.192.245"
	user = "developer"
	pass = "Hackjunction2018"
)

var (
	sharedToken string
	mutex       = &sync.RWMutex{}
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func KeepTokenFresh() {
	auth := func() {
		log.Println("[INFO] Refreshing authorization token...")
		token, err := authenticate()
		if err != nil {
			log.Printf("[ERROR] failed to authenticate: %v\n", err)
		}
		mutex.Lock()
		sharedToken = token
		mutex.Unlock()
		log.Println("[INFO] New token obtained!")
	}

	auth()
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			auth()
		}
	}()
}

func authenticate() (token string, err error) {
	req, err := http.NewRequest("POST", "https://"+host+"/dna/system/api/v1/auth/token", nil)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(user+":"+pass)))
	var auth models.AuthenticationAPIResponse
	err = send(req, &auth)
	if err != nil {
		return
	}
	return auth.Token, nil
}

func getToken() string {
	mutex.RLock()
	defer mutex.RUnlock()
	return sharedToken
}

func Get(path string, headers H, response interface{}) error {
	req, err := http.NewRequest("GET", "https://"+host+"/dna/intent/api/v1"+path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Auth-Token", getToken())
	for name, value := range headers {
		req.Header.Add(name, value)
	}
	return send(req, response)
}

func Post(path string, headers H, response interface{}, request interface{}) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://"+host+"/dna/intent/api/v1"+path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("X-Auth-Token", getToken())
	req.Header.Add("Content-Type", "application/json")
	for name, value := range headers {
		req.Header.Add(name, value)
	}
	return send(req, response)
}

func send(req *http.Request, response interface{}) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(fmt.Sprintf("unexpected response code (%d): %s", res.StatusCode, string(body)))
	}
	defer res.Body.Close()
	if response == nil {
		return nil
	}
	return json.NewDecoder(res.Body).Decode(&response)
}
