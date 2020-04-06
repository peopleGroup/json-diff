package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

//Endpoints - contains live & local endpoints
type Endpoints struct {
	Live  string
	Local string
}

//RequestParams - contains params needed to make http request
type RequestParams struct {
	QueryParams map[string]string
	Headers     map[string]string
	Body        map[string]interface{}
}

//RequestData - contains data needs to make request to endpoints
type RequestData struct {
	Endpoints     *Endpoints
	RequestParams *RequestParams
}

func initializeDirectories() {
	os.MkdirAll("./log/local", os.ModePerm)
	os.MkdirAll("./log/live", os.ModePerm)
	os.MkdirAll("./log/diff", os.ModePerm)
}

func main() {
	initializeDirectories()

	var wg sync.WaitGroup
	wg.Add(1)

	// Buffered channel so that producer and consumers don't block,
	// feel free to increase the buffer size
	rc := make(chan RequestData, 100)

	start := time.Now()

	go producer(rc, &wg)

	go consumer(rc, &wg)
	go consumer(rc, &wg)
	go consumer(rc, &wg)
	go consumer(rc, &wg)

	fmt.Println("Waiting...")

	wg.Wait()

	fmt.Println("Running json-diff...")
	out, err := exec.Command("sh", "-c", "node diff.js").Output()
	fmt.Println("exec:", out, err)

	fmt.Println("Done...")
	fmt.Println(time.Now().Sub(start))

}
