package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"url/models"
	"url/utils"
)

func ShorthandUrl(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	short := &models.ShortenRequest{}
	url := &models.Response{
		ShortenRequest: *short,
	}
	if err := json.NewDecoder(r.Body).Decode(&short); err != nil {
		http.Error(w, "Invalid Json format", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(short.Url, "http://") && !strings.HasPrefix(short.Url, "https://") {
		short.Url = "https://" + short.Url
	}
	url.ShortenRequest.Url = short.Url

	var code string
	if url.UrlValidate() {
		http.Error(w, "url already exists", http.StatusConflict)
		return
	}

	for {
		code = utils.GenerateCode()
		url.ResponseUrl = code
		if !url.CodeValidate() {

			break
		}
		fmt.Println("short code already exists, generating a new one...")
	}

	if err := url.Insert(); err != nil {
		http.Error(w, "Error inserting the url", http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		ResponseUrl:    "http://localhost:3051/" + code,
		ShortenRequest: models.ShortenRequest{Url: short.Url},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func Redirect(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	requestedcode := strings.TrimPrefix(r.URL.Path, "/")

	url := &models.Response{}
	url.ResponseUrl = requestedcode

	if err := url.Get(); err != nil {
		http.Error(w, "Redirect Url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.ShortenRequest.Url, http.StatusFound)

}
