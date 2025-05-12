package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"yt-go/api/handler"
)

var Srv *http.Server

func setEndpointHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("POST /_yt_", handler.URLHandler)
}

func StopServer() error {
	if Srv == nil {
		fmt.Printf("HTTP server Shutdown error. Server not found")
		return http.ErrServerClosed
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err = Srv.Shutdown(ctx)
	fmt.Println("Server shutdown gracefully")
	return err
}

func StartServer() {
	var mux *http.ServeMux = http.NewServeMux()
	setEndpointHandlers(mux)

	Srv = &http.Server{
		Addr:    ":43214",
		Handler: mux,
	}

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := StopServer(); err != nil {
			fmt.Printf("HTTP server Shutdown error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	fmt.Println("Server listening on http://localhost:43214")
	if err := Srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server ListenAndServe error: %v", err)
	}

	<-idleConnectionsClosed
	fmt.Println("Server shutdown gracefully")
}
