package http

import (
	"fmt"
	"net/http"
	"io/ioutil"
	Config "../config"
)

func HttpRequest(key string) []byte {

	apiKey := Config.ApiKey
	url := "https://api.themoviedb.org/3/movie/" + key + "?api_key=" + apiKey

	// fmt.Println(url)

	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Println("Error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error")
	}
	
	return body
}

