package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://datausa.io/api/data?drilldowns=Nation&measures=Population"

func main() {
	fmt.Println("Handling URL in golang")

	//get information from url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	queryParms := result.Query()
	fmt.Println(queryParms)

	for _, val := range queryParms {
		fmt.Println(val) // key value pair map
	}

	//build your URL
	partsofURL := &url.URL{
		Scheme:   "https",
		Host:     "datausa.io",
		Path:     "api/data",
		RawQuery: "drilldowns=Nation&measures=Population",
	}

	fmt.Println(partsofURL.String())
}
