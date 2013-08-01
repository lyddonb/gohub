package client

import (
	"encoding/json"
	"fmt"
	//"gohub/requests"
	//"gohub/user"
	"io/ioutil"
	"net/http"
)

const base_gh_url = "https://api.github.com/"

type GithubClient struct {
	Auth GithubAuth
	http.Client
}

func (gc *GithubClient) Get(url string) (*http.Response, error) {
	// Create a new GET request for the url.
	request, err := http.NewRequest("GET", base_gh_url+url, nil)

	// Pass the reference to the pointer of the request to the authenticate
	// method for the GithubAuth implementation.
	gc.Auth.Authenticate(&*request)

	// Make the request.
	response, err := gc.Do(request)

	if response.StatusCode != 200 {
		// TODO: Handle this error better.
		fmt.Println(err)
		return response, err
	}

	return response, err
}

func parseJson(content []byte, ghType interface{}) (interface{}, error) {
	err := json.Unmarshal(content, ghType)

	return ghType, err
}

func (gc *GithubClient) MakeRequest(url string, ghType interface{}) (interface{}, error) {
	response, err := gc.Get(url)

	defer response.Body.Close()
	b, _ := ioutil.ReadAll(response.Body)
	entity, err := parseJson(b, ghType)

	return entity, err
}
