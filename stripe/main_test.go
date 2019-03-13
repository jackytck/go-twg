package main

import (
	"fmt"
	"net/http"
	"testing"

	stripe "github.com/jackytck/twg/stripe/v1"
)

type App struct {
	Stripe *stripe.Client
}

func (a *App) Run() {}

func TestApp(t *testing.T) {
	client, mux, teardown := stripe.TestClient(t)
	defer teardown()

	mux.HandleFunc("/v1/charges", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id":"ch_1EDZBv2eZvKYlo2CtEST0JbK","amount":2000,"description":"Charge for demo purposes.","status":"succeeded"}`)
	})

	// Now inject client into your app and run your tests - they will use your
	// local test server using this mux
	app := App{
		Stripe: client,
	}
	app.Run()
	// ...
}