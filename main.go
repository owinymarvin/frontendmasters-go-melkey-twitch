package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/owinymarvin/frontendmasters-go-melkey-twitch/internal/app"
	"github.com/owinymarvin/frontendmasters-go-melkey-twitch/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "GO backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are running on port: %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
