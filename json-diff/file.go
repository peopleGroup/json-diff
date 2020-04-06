package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func prettyfy(jstr string) (string, error) {
	b := []byte(jstr)
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.String(), err
}

func writeJSONStringToFile(filename, jsonString string) {
	fn, err := os.Create(filename)
	defer fn.Close()
	prettyJSON, err := prettyfy(jsonString)
	if err != nil {
		fmt.Println(err)
	}
	fn.WriteString(prettyJSON)
}
