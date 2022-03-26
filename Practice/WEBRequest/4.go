package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to create json and consume json data")

	//encoding json
	//Encodejson()

	DecodeJson()
}

func Encodejson() {
	mahbubCourses := []course{
		{"c#", 200, "programming", "123", []string{"a", "b", "c"}},
		{"cpp", 200, "programming", "123", []string{"a", "b", "c"}},
		{"java", 200, "programming", "123", nil},
	}
	//pacakage this data as json data

	finalJson, err := json.MarshalIndent(mahbubCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsondatafromWeb := []byte(` [{
	"courseName": "c#",
	"Price": 200,
	"website": "programming",
	"tags": [ "a", "b", "c"]
	},
	{
		"courseName": "cpp",
		"Price": 200,
		"website": "programming",
		"tags": [ "a", "b", "c"]
	},
	{
		"courseName": "java",
		"Price": 200,
		"website": "programming"
	}]`)

	//json valid or not
	var mahbubCourse []course
	checkValid := json.Valid(jsondatafromWeb)

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsondatafromWeb, &mahbubCourse)
		fmt.Printf("%#v\n", mahbubCourse)
	} else {
		fmt.Println("not valid")
	}

	//some case store json in key value pair

	var myOnlineCourses []map[string]interface{}
	json.Unmarshal(jsondatafromWeb, &myOnlineCourses)
	fmt.Printf("%#v\n", myOnlineCourses)

 
}
