package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

/*
* Data structure of the JSON coming from web client
 */
type URLFromClient struct {
	URL string `json:"url"`
}

/*
* HTTP request handler function
 */
func handler(response http.ResponseWriter, request *http.Request) {
	/* Allow client to perform a POST request if it asks through OPTIONS request */
	if request.Method == http.MethodOptions {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		response.WriteHeader(http.StatusNoContent)
		return
	}

	if request.Method != http.MethodPost {
		http.Error(response, "Only POST method allowed", http.StatusMethodNotAllowed)
		fmt.Println(request)
		return
	}

	fmt.Println(request.Method)

	var data URLFromClient
	var decoder *json.Decoder = json.NewDecoder(request.Body)
	var errDecoding error = decoder.Decode(&data)

	if errDecoding != nil {
		http.Error(response, "Invalid JSON", http.StatusBadRequest)
		return
	}

	parsedURL, errParsing := url.Parse(data.URL)

	if errParsing != nil {
		http.Error(response, "Invalid URL", http.StatusBadRequest)
		return
	}

	if isUrlSafe(parsedURL) {
		if urlIsAboutVideo, err := isAboutVideo(data.URL); err != nil {
			http.Error(response, "Error parsing URL or checking its content", http.StatusBadRequest)
			return
		} else if !urlIsAboutVideo {
			http.Error(response, "URL not about video", http.StatusBadRequest)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{
			"status": "ok",
			"url":    data.URL,
		})
	}

	/* Not a safe URL ⚠️ Miala malaky :O */
	http.Error(response, "", http.StatusForbidden)
}

func StartServer() {
	http.HandleFunc("/_yt_", handler)

	fmt.Println("Listening on http://localhost:43214")

	/* Returns Fatal error if server is closed unexpectedly */
	if err := http.ListenAndServe(":43214", nil); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("ListenAndServer() : %v", err)
	}
}
