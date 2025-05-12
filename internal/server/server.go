package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"yt-go/api/handler"
)

func StartServer() {
	http.HandleFunc("/_yt_", handler.URLHandler)

	fmt.Println("Listening on http://localhost:43214")

	/* Returns Fatal error if server is closed unexpectedly */
	if err := http.ListenAndServe(":43214", nil); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("ListenAndServer() : %v", err)
	}
}
