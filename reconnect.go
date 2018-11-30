package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReconRequest for recon struct
type ReconRequest struct {
	ByUser struct {
		ByChannel string `json:"byChannel"`
		ByUser    string `json:"byUser"`
	} `json:"ByUser"`
	Customer struct {
		CustomerID int `json:"CustomerID"`
	} `json:"Customer"`
	Product struct {
		Product []struct {
			ProductID int `json:"ProductId"`
		} `json:"Product"`
	} `json:"Product"`
	Reason int `json:"Reason"`
	Target struct {
		Target     int    `json:"Target"`
		TargetDate string `json:"TargetDate"`
	} `json:"Target"`
}

// ReconResult for recon result
type ReconResult struct {
	ErrorCode   int    `json:"ErrorCode"`
	ErrorDesc   string `json:"ErrorDesc"`
	ResultValue string `json:"ResultValue"`
	ProductID   int    `json:"ProductId"`
}

//ReconnectProduct for ...
func ReconnectProduct(req ReconRequest) ReconResult {
	var myReturn ReconResult

	/*{
		"ByUser":{
			"byChannel":"9912",
			"byUser":"9912"
		},
		"Customer":{
			"CustomerID":60646187
		},
		"Product":{
			"Product":[{
				"ProductId":265164823
			}]
		},
		"Reason":487,
		"Target":{
			"Target":0,
			"TargetDate":""
		}
	}*/

	//Call Rest API ICC Reconnect
	//jsonData := map[string]string{"ThaiId": "3909800183384"}
	jsonValue, _ := json.Marshal(req)
	response, err := http.Post("http://172.22.203.68//ConvergenceIBSTVG/ICC/reconnectproduct", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		//myReturn = json.Unmarshal(string(data))
		err = json.Unmarshal(data, &myReturn)
		if err != nil {
			panic(err)
		}
	}

	return myReturn
}
