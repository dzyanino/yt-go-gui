package handler

import (
	"yt-go/api/middleware"
	"yt-go/internal/downloader"
	"yt-go/internal/types"

	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

/*
 *
 * HTTP request handler function
 *
 */
func URLHandler(response http.ResponseWriter, request *http.Request) {
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
		return
	}

	fmt.Println(request.Method)

	var data types.URLData
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

	if !middleware.IsUrlSafe(parsedURL) {
		/* Not a safe URL ⚠️ Miala malaky :O */
		http.Error(response, "", http.StatusForbidden)
		return
	}

	if urlIsAboutVideo, err := middleware.IsAboutVideo(data.URL); err != nil {
		http.Error(response, "Error parsing URL or checking its content", http.StatusBadRequest)
		return
	} else if !urlIsAboutVideo {
		http.Error(response, "URL not about video", http.StatusBadRequest)
		return
	}

	downloader.DownloadVideo(data.URL)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]string{
		"status": "ok",
		"url":    data.URL,
	})
}
