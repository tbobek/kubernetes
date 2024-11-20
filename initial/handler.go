package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	klog "k8s.io/klog/v2"
)

const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"
const workerUrl = "http://172.17.0.3:8765/calc"

func handleCalc(c *gin.Context) {
	t0 := time.Now()
	var requestBody Request
	var response Response
	var body []byte
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	if cb, ok := c.Get(BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if err == nil {

		results := make([]int, 0)
		ch := make(chan Result, requestBody.Calls)
		for i := 0; i < requestBody.Calls; i++ {
			// make request to other service
			requestBody.Id = uuid.NewString()
			klog.Infof("making request to %s: %+v", workerUrl, requestBody)

			go handleSubRequest(requestBody, ch)
		}
		for i := 0; i < requestBody.Calls; i++ {
			myResult := <-ch
			results = append(results, myResult.Result)
		}
		response = Response{
			Id:       uuid.NewString(),
			Result:   results,
			Status:   "ok",
			Error:    false,
			ErrorMsg: "",
			Input:    string(body),
			Duration: time.Now().Sub(t0).String(),
		}
		c.IndentedJSON(http.StatusOK, response)
	} else {
		response = Response{
			Id:       uuid.NewString(),
			Result:   []int{},
			Status:   "nok",
			Error:    true,
			ErrorMsg: "bad request",
			Input:    string(body),
		}
		c.IndentedJSON(http.StatusBadRequest, response)
	}
}

func handleSubRequest(requestBody Request, chResult chan Result) {
	requestBody.Id = uuid.NewString()
	var result Result
	klog.Infof("making request to %s: %+v", workerUrl, requestBody)
	res, err := MakeRequest(requestBody, workerUrl)
	if err != nil {
		klog.Warning("request terminated with error")
		result.Result = -2
	} else {
		data, _ := io.ReadAll(res.Body)
		json.Unmarshal(data, &result)
	}
	chResult <- result
}

func MakeRequest(r Request, myUrl string) (*http.Response, error) {
	jsonData, err := json.Marshal(r)
	client := &http.Client{}
	rq, _ := http.NewRequest(http.MethodPost, myUrl, bytes.NewReader(jsonData))
	rq.Header.Add("Content-Type", "application/json")
	res, err := client.Do(rq)
	return res, err
}

// this service triggers the
