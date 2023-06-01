package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

//this is for test

func sendPatchRequest(url string, requestBody map[string]interface{}, wg *sync.WaitGroup, ch chan<- string, i int) {
	defer wg.Done()

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		ch <- fmt.Sprintf("Failed to marshal JSON: %s", err.Error())
		return
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		ch <- fmt.Sprintf("Failed to create request: %s", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Request failed: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("Id: %d\nStatus code: %d\n", i, resp.StatusCode)
}
func sendGetRequest(url string, wg *sync.WaitGroup, ch chan<- string, i int) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- fmt.Sprintf("Failed to create request: %s", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Request failed: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err.Error())
		return
	}

	ch <- fmt.Sprintf("Id: %d\nStatus code: %d\nwith body: %s", i, resp.StatusCode, string(body))
}

func main() {
	patchUrl := "http://localhost:8080/gifts"
	getUrl := "http://localhost:8080/gifts/5"
	balanceUrl := "http://localhost:8081/wallets/"

	numRequests := 10

	var wg, wg2, wg3 sync.WaitGroup
	wg.Add(numRequests)
	wg2.Add(numRequests)
	wg3.Add(numRequests)

	ch := make(chan string, numRequests)
	ch2 := make(chan string, numRequests)
	ch3 := make(chan string, numRequests)

	for i := 0; i < numRequests; i++ {
		phoneNumber := fmt.Sprintf("09135763%03d", i+1)

		requestBody := map[string]interface{}{
			"code":        "5",
			"phoneNumber": phoneNumber,
		}

		go sendPatchRequest(patchUrl, requestBody, &wg, ch, i)
		go sendGetRequest(getUrl, &wg2, ch2, i)
		go sendGetRequest(balanceUrl+phoneNumber, &wg3, ch3, i)
	}

	wg.Wait()

	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
	for msg := range ch2 {
		fmt.Println(msg)
	}
	for msg := range ch3 {
		fmt.Println(msg)
	}
}
