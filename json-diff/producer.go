package main

import "sync"

func producer(rc chan RequestData, wg *sync.WaitGroup) {
	defer wg.Done()

	//TODO: implement your own producer fn
	// The main job of producer is to genetate different combinations of
	// requests and write it to the Request Channle (rc)

	endpoints := &Endpoints{
		Live:  "http://local.api.endpoint",
		Local: "https://live.api.endpoint",
	}

	queryParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	headers := map[string]string{
		"x-api-token": "XXXXXXXXX",
		"x-device":    "YYYYYYYYYY",
	}

	params := &RequestParams{
		QueryParams: queryParams,
		Headers:     headers,
	}

	//This is just a single request, you could have multiple requests in a loop
	rc <- RequestData{
		Endpoints:     endpoints,
		RequestParams: params,
	}
	//Don't forget to add to the wait group after writing to the channel
	wg.Add(1)

	//Note: Don't forget to close the channel once you are done.
	close(rc)
}
