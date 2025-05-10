package yt_go_server

import (
	"net/http"
	"net/url"
	"strings"
)

/*
* Checks if the received URL is ending with a video extension
 */
func isVideoURL(u *url.URL) bool {
	var videoExtensions []string = []string{".mp4", ".webm", ".mov", ".avi", ".mkv", ".flv", "wmv", "mpeg"}
	var path string = strings.ToLower(u.Path)

	for _, extension := range videoExtensions {
		if strings.HasSuffix(path, extension) {
			return true
		}
	}

	return false
}

/*
* Checks if the URL content a video
 */
func isVideoContent(urlString string) (bool, error) {
	response, err := http.Head(urlString)

	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	var contentType string = response.Header.Get("Content-Type")
	return strings.HasPrefix(contentType, "video/"), nil
}

/*
* Function using the two validations above to check if any given URL is about video
 */
func isAboutVideo(urlString string) (bool, error) {
	parsedURL, err := url.Parse(urlString)

	if err != nil {
		return false, err
	}

	if isVideoURL(parsedURL) {
		return true, nil
	}

	return isVideoContent(urlString)
}
