package yt_go_server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

/*
* Data structure of the JSON coming from web client
 */
type DataFromClient struct {
	URL string `json:"url"`
}

/*
* HTTP request handler function
 */
func handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var data DataFromClient
	var decoder *json.Decoder = json.NewDecoder(request.Body)
	var errDecoding error = decoder.Decode(&data)

	if errDecoding != nil {
		http.Error(response, "Invalid JSON", http.StatusBadRequest)
		return
	}

	parsedURL, errParsing := url.Parse(data.URL)

	if errParsing != nil {
		http.Error(response, "Invalid URL", http.StatusBadRequest)
	}

	if isUrlSafe(parsedURL) {
		videoURL, err := isAboutVideo(data.URL)

		if err != nil {
			http.Error(response, "Error parsing URL or checking its content", http.StatusBadRequest)
		}

		if !videoURL {
			http.Error(response, "URL not about video", http.StatusBadRequest)
		}

		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{
			"status": "ok",
			"url":    data.URL,
		})
		fmt.Println(data.URL)
	}

	/* Not a safe URL ⚠️ Miala malaky :O */
	http.Error(response, "", http.StatusForbidden)
}

func startServer() {
	http.HandleFunc("/_yt_", handler)

	fmt.Println("Listening on http://localhost:43214")

	/* Returns Fatal error if...who knows */
	log.Fatal(http.ListenAndServe(":43214", nil))
}
