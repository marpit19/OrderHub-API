package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	stopServerAfterDelay(server, 10*time.Second)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yo bro!"))
}

func stopServerAfterDelay(server *http.Server, delay time.Duration) {
	time.AfterFunc(delay, func() {
		fmt.Println("Stopping the server.....")
		err := server.Close()
		if err != nil {
			fmt.Println("Error while stopping the server:", err)
		}
	})
}
