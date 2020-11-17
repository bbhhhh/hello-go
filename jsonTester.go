package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int `json:"released"`
	Actors []string
	aa     string
}

func main() {
	movies := []Movie{
		{Title: "starwars", Year: 1980, Actors: []string{"aaa", "bbb"}, aa: "111"},
		{Title: "starwars2", Year: 1990, Actors: []string{"ccc", "ddd"}},
	}
	data, err := json.MarshalIndent(movies, "", "  ")

	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	data, err = json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	var yyy []Movie
	if err := json.Unmarshal(data, &yyy); err != nil {
		log.Fatalf("json unmarshaling failed: %s", err)
	}
	fmt.Println(yyy)

}
