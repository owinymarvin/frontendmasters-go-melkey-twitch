package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/owinymarvin/frontendmasters-go-melkey-twitch/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "GO backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

	app.Logger.Printf("we are running on port %d\n", port)
	http.HandleFunc("/health", HealthCheck)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
