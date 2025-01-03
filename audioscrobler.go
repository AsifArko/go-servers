package main

import "fmt"
import "net/http"
import (
	"io/ioutil"
	"encoding/json"
	"os"
)


type Tracks struct {
	Toptracks []Toptracks_info
}

type Toptracks_info struct {
	Track []Track_info
	Attr  []Attr_info
}

type Track_info struct {
	Name       string
	Duration   string
	Listeners  string
	Mbid       string
	Url        string
	Streamable []Streamable_info
	Artist     []Artist_info
	Attr       []Track_attr_info
}

type Attr_info struct {
	Country    string
	Page       string
	PerPage    string
	TotalPages string
	Total      string
}

type Streamable_info struct {
	Text      string
	Fulltrack string
}

type Artist_info struct {
	Name string
	Mbid string
	Url  string
}

type Track_attr_info struct {
	Rank string
}

func get_content() {
	// json data
	url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}


	data := new(Tracks)

	err = json.Unmarshal(body, &data)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(data)

	fmt.Println(string(body))

	os.Exit(0)
}

func main() {
	get_content()
}
