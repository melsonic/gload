package util

import (
	"fmt"
	"net/http"
	"time"
	"sync"
)

type CoreResult struct {
	success      int
	totalRequest int
	totalTime    int
}

func makeRequest(client *http.Client, gla *GLoadArgs, totalRequest int, result *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < totalRequest; i++ {
		resp, reqError := client.Get(gla.Url)
		statusCode := http.StatusOK
		if reqError != nil {
			statusCode = http.StatusNotFound
		} else {
			statusCode = resp.StatusCode
		}
		if statusCode == http.StatusOK {
			*result = *result + 1
		}
	}
}

func CoreFunction(gla *GLoadArgs) {
	var wg = &sync.WaitGroup{}
	client := &http.Client{}
	var requestPerUser int = gla.RequestCount / gla.ConcurrentUsers
	resultCollector := make([]int, gla.ConcurrentUsers)
	var startTime time.Time = time.Now()
	// let first user make requestPerUser+additional requests
	wg.Add(1)
	go makeRequest(client, gla, requestPerUser+(gla.RequestCount%gla.ConcurrentUsers), &resultCollector[0], wg)
	for i := 1; i < gla.ConcurrentUsers; i++ {
		wg.Add(1)
		go makeRequest(client, gla, requestPerUser, &resultCollector[i], wg)
	}
	wg.Wait()
	var endTime time.Time = time.Now()
	var result CoreResult = CoreResult{totalRequest: gla.RequestCount, totalTime: int(endTime.Sub(startTime).Milliseconds())}
	computeResult(resultCollector, &result)
	displayResult(result)
}

func computeResult(resultCollector []int, result *CoreResult) {
	for _, v := range resultCollector {
		result.success += v
	}
}

func displayResult(result CoreResult) {
	fmt.Println("=========== Result ===========")
	fmt.Println("Success    : ", result.success)
	fmt.Println("Failure    : ", result.totalRequest-result.success)
	fmt.Println("Hit Ratio  : ", (float64(result.success)/float64(result.totalRequest))*100)
	fmt.Println("Total time : ", result.totalTime)
}
