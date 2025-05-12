package server

import (
	"net/http"
	"net/url"
	"strings"
)

/*
* Checks if the received URL is ending with a video extension
 */
func isVideoURL(u *url.URL) bool {
	var videoExtensions = [...]string{
		".mp4", ".webm", ".mov", ".avi", ".mkv", ".flv", ".wmv", ".mpeg",
		".3gp", ".ogg", ".m4v", ".ts", ".vob", ".m2ts", ".mts", ".f4v",
		".rm", ".rmvb",
	}
	var path string = strings.ToLower(u.Path)

	for _, extension := range videoExtensions {
		if strings.HasSuffix(path, extension) {
			return true
		}
	}

	return false
}

/*
* Checks if the URL Content-Type is a video
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
* Checks if the link is from known sites (actually Youtube)
 */
func isYoutubeLink(urlString string) bool {
	return strings.HasPrefix(urlString, "https://www.youtube.com/watch?") || strings.HasPrefix(urlString, "https://www.youtube.com/short?")
}

/*
* Function using all the validations above to check if any given URL is about video
 */
func isAboutVideo(urlString string) (bool, error) {
	parsedURL, err := url.Parse(urlString)

	if err != nil {
		return false, err
	}

	/* First check */
	if isVideoURL(parsedURL) {
		return true, nil
	}

	/* Second check */
	if contentIsVideo, err := isVideoContent(urlString); err != nil {
		return false, err
	} else if contentIsVideo {
		return true, nil
	}

	/* Third check */
	return isYoutubeLink(urlString), nil
}
