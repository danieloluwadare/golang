package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

func topArticles(authorUsername string, limit int) []string {
	url := "https://jsonmock.hackerrank.com/api/articles"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Add query parameters
	q := req.URL.Query()
	q.Add("author", authorUsername)
	req.URL.RawQuery = q.Encode()

	// Create request client
	client := http.Client{Timeout: 5 * time.Second}

	// Make desired call
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp.Body)

	// Store json data
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}
	fmt.Println(data["data"])

	//Main Algorithm
	fcts := make([]BookObj, 0) // same as []int{0, 0}

	for _, s := range data["data"].([]interface{}) {

		var bookTitle string
		var count int = 0
		//fmt.Println(i, s)

		obj:=s.(map[string]interface{})

		if obj["title"] == nil && obj["story_title"] == nil {
			continue
		}

		if obj["title"] == nil {
			bookTitle = obj["story_title"].(string)
		} else {
			bookTitle = obj["title"].(string)
		}

		if obj["num_comments"] != nil{
			c := obj["num_comments"].(float64)
			//fmt.Println(c)
			count=int(c)
		}

		//fmt.Println(bookTitle)
		//fmt.Println(count)
		fct := BookObj{}
		fct.bookTitle=bookTitle
		fct.comments=count

		fcts = append(fcts, fct)

	}

	fcts = append(fcts, BookObj{"Daba",3})
	fcts = append(fcts, BookObj{"Aba",3})
	fcts = append(fcts, BookObj{"Baba",3})




	fmt.Println("UnSorted Array")

	fmt.Println(fcts)

	sort.SliceStable(fcts, func(i, j int) bool {
		if fcts[i].comments == fcts[j].comments{
			return fcts[i].bookTitle < fcts[j].bookTitle
		}
		return fcts[i].comments > fcts[j].comments
	})
	fmt.Println("Sorted Array")
	fmt.Println(fcts)

	books := make([]string, 0) // same as []int{0, 0}


	for i, s := range fcts {
		checkLimit := i+1
		if checkLimit==limit{
			break
		}
		books = append(books, s.bookTitle)
		fmt.Println(i)
	}

	return books

}

func topStories(name string) []string {
	baseUrl:="https://jsonmock.hackerrank.com/api/movies/search?Title="+name

	data, _ := performHttpRequest(baseUrl);



	movieTitles := make([]string, 0) // same as []int{0, 0}
	for _, s := range data["data"].([]interface{}) {
		obj:=s.(map[string]interface{})
		movieTitles = append(movieTitles, obj["Title"].(string))
	}
	sort.Strings(movieTitles)
	fmt.Println(movieTitles)

	return movieTitles;

}

func performHttpRequest(url string) (map[string]interface{}, error){

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
	//fmt.Println(data["data"])

	return data, err
}

type BookObj struct {
	bookTitle string
	comments int
}
