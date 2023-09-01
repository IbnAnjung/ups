package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		fmt.Println("fail to create http client")
		os.Exit(0)
	}

	request.Header.Set("Content-Type", "application/json")

	client := new(http.Client)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("fail to get data")
		os.Exit(0)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("fail to get http request")
		os.Exit(0)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("fail to read body")
		os.Exit(0)
	}

	fmt.Println(string(body))

}
