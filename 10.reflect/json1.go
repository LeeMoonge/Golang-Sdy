package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Price int `json:"rmb"`
	Actors []string `json:"actors"`
}

func main()  {
	movie := Movie{Title: "喜剧之王", Year: 2000, Price: 10, Actors: []string{"aab", "bba", "aabb", "周星驰"}}

	//编码的过程 结构体 --> json
	jsonStr, err := json.Marshal(movie)
	if err != nil{
		fmt.Println("json Marshal Error = ", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	//编码的过程 jsonstr --> 结构体
	//jsonStr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["aab","bba","aabb","周星驰"]}
	myMovie := movie
	unmarshalErr := json.Unmarshal(jsonStr, &myMovie)
	if unmarshalErr != nil{
		fmt.Println("json.Unmarshal Error", unmarshalErr)
		return
	}
	fmt.Printf("movie Struct = %v\n", myMovie)
}