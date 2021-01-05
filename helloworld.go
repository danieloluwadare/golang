package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// This contains the main method
func main() {
	// fmt.Println("Hello World! This is Daniel first Go program\n More Codes to come")
	// testDataTypeExample()
	// testconstantExamples()
	// testForLoopsExamples()
	// testWhileLoopsExamples()
	// testForloopOnArray()
	// testForloopOnArrayMutation()
	// testForEachLoopUsingRangeOnArray()
	// testIfStatements()
	// testSwitchCaseStatemets()
	// testArraysExample()
	// testSliceExample()
	// testUpdateSliceExample()
	// testMapExample()
	// testString()
	// testFuctionWithFunctionArgument(1, 3)
	// testVariadicFunction()
	// testStructs()
	// testInterface()
	// testTypePlayGround()
	// testDerivedTypesOfmapOfEvents()

	// testPointers()
	//testMethodsOfStructs()
	//fmt.Println("Top articles by author")
	//fmt.Println(topArticles("olalonde", 3))

	topStories("spiderman")

	TestError()
}



func performHttpRequest2(url string) (map[string]interface{}, error){

	baseUrl:= url

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		panic(err)
		return nil, err
	}

	//// Add query parameters
	//q := req.URL.Query()
	//q.Add("author", authorUsername)
	//req.URL.RawQuery = q.Encode()

	// Create request client
	client := http.Client{Timeout: 5 * time.Second}

	// Make desired call
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		return nil, err
	}
	//fmt.Println(resp.Body)

	// Store json data
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
		return nil, err

	}
	fmt.Println(data["data"])

	return data, err
}