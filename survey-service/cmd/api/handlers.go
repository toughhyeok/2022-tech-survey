package main

import (
	"net/http"
)

func (app *Config) GetLanguages(w http.ResponseWriter, r *http.Request) {
	languages, err := app.Models.Language.GetLanguages()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get languages success",
		Data:    languages,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) GetWebframeworks(w http.ResponseWriter, r *http.Request) {
	webframeworks, err := app.Models.Webframework.GetWebframeworks()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get webframeworks success",
		Data:    webframeworks,
	}

	app.writeJSON(w, http.StatusOK, payload)
}
