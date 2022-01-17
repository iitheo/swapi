package bushaservices

import (
	"encoding/json"
	"fmt"
	"github.com/iitheo/theobusha/pkg/app/config/httpresponses"
	"github.com/iitheo/theobusha/pkg/app/models/filmsmodel"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetAllSwapiFilms(swapiResponse chan httpresponses.HttpResponse) {
	var (
		resp        httpresponses.HttpResponse
		listOfFilms filmsmodel.FilmsVM
	)
	baseUrl := strings.TrimSpace(os.Getenv("SWAPIBASEURL"))
	getUrl := baseUrl + "/films/"

	request, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}

	// Do sends an HTTP request and
	response, err := client.Do(request)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	err = json.Unmarshal(body, &listOfFilms)

	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	resp.Success = true
	resp.Message = fmt.Sprintf("films successfully fetched")
	resp.Data = listOfFilms
	swapiResponse <- resp
	return
}

func GetAllSwapiCharacters(swapiResponse chan httpresponses.HttpResponse) {
	var (
		resp             httpresponses.HttpResponse
		listOfCharacters filmsmodel.People
	)
	baseUrl := strings.TrimSpace(os.Getenv("SWAPIBASEURL"))
	getUrl := baseUrl + "/people/"

	request, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}

	// Do sends an HTTP request and
	response, err := client.Do(request)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	err = json.Unmarshal(body, &listOfCharacters)

	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		swapiResponse <- resp
		return
	}

	resp.Success = true
	resp.Message = fmt.Sprintf("characters successfully fetched")
	resp.Data = listOfCharacters
	swapiResponse <- resp
	return
}
