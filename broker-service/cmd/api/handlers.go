package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RequestPayload struct {
	Action string `json:"action"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
	}

	switch requestPayload.Action {
	case "survey/languages":
		app.languages(w)
	case "survey/webframeworks":
		app.webframeworks(w)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) languages(w http.ResponseWriter) {
	surveyServiceURL := "http://survey-service/"

	languages, err := surveyRequest(surveyServiceURL + "languages")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get language results success!"
	payload.Data = languages

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) webframeworks(w http.ResponseWriter) {
	surveyServiceURL := "http://survey-service/"

	webframeworks, err := surveyRequest(surveyServiceURL + "webframeworks")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get webframeworks results success!"
	payload.Data = webframeworks

	app.writeJSON(w, http.StatusOK, payload)
}

func surveyRequest(URL string) (any, error) {
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("internal server error")
	}

	var jsonFromService jsonResponse

	// decode the json from the survey service
	fmt.Printf("%v", response.Body)
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		return nil, err
	}

	if jsonFromService.Error {
		return nil, errors.New("internal server error")
	}

	fmt.Printf("%v", jsonFromService.Data)
	return jsonFromService.Data, nil
}
