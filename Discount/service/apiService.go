package service

import (
	"Discount/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CallAnotherAPI(endpoint string, method string, requestBody interface{}) (models.ApiResponse, error) {
	println("api called")
	url := "http://localhost:8081/" + endpoint

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return models.ApiResponse{}, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.ApiResponse{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.ApiResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.ApiResponse{StatusCode: resp.StatusCode}, err
	}

	return models.ApiResponse{
		Bytes:      body,
		StatusCode: resp.StatusCode,
	}, nil
}
