package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("How to make GET, post, post fromdata file/ some think request in golang")
	//PerformGetRequest()
	//PerformPostRequest()
}

func PerformGetRequest() {
	const myurl = "https://reqres.in/api/users?page=2"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("content length:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)
	fmt.Println("Byte count : ", bytecount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	const myUrl = "https://reqres.in/api/users"

	//json payload

	requestbody := strings.NewReader(`
		{
			"name": "morpheus",
			"job": "leader"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestbody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("status :", response.Status)
	fmt.Println(string(content))
}

func PerformDatajsonRequest() {
	const myurl = "...."

	//formdata image upload or some think like this
	data := url.Values{}
	data.Add("firstname", "xyz")
	data.Add("lastname", "xyz")
	data.Add("email", "xyz")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("status :", response.Status)
	fmt.Println(string(content))

}
