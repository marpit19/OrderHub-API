package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	stopServerAfterDelay(server, 10*time.Second)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Print("failed to listen", err)
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
