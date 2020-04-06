package main

import (
	"sync"
)

func consumer(rc chan RequestData, wg *sync.WaitGroup) {

	//TODO: Implement your own consumner fn
	// The main job of consumer is to make a call to API & write the response to a file

	for request := range rc {
		someUniqueFileName := "unique_file_name.json"

		liveResponse, err := callAPI(request.Endpoints.Live, request.RequestParams)
		if err != nil {
			writeJSONStringToFile("./live/"+someUniqueFileName, liveResponse)
		}
		localResponse, err := callAPI(request.Endpoints.Local, request.RequestParams)
		if err != nil {
			writeJSONStringToFile("./local/"+someUniqueFileName, localResponse)
		}

		wg.Done()
	}
}
