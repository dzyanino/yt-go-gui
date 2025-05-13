package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"yt-go/api/handler"
)

var (
	Srv                     *http.Server
	stopOnce                sync.Once
	isRunning               bool
	ErrServerIsNotRunning   = errors.New("server is not running")
	ErrServerAlreadyStopped = errors.New("server already shutdown")
)

func setEndpointHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("POST /_yt_", handler.URLHandler)
}

func StopServer() error {
	var stopError error

	if !isRunning {
		return ErrServerAlreadyStopped
	}

	stopOnce.Do(func() {
		if Srv == nil {
			fmt.Println("HTTP server shutdown error: server not found")
			stopError = ErrServerIsNotRunning
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		stopError = Srv.Shutdown(ctx)
		isRunning = false
		Srv = nil
	})

	return stopError
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

	isRunning = true
	fmt.Println("Server listening on http://localhost:43214")

	if err := Srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server ListenAndServe error: %v", err)
	}

	<-idleConnectionsClosed
	fmt.Println("Server shutdown gracefully")
}
