package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)


type APISetting func(ctx context.Context, w http.ResponseWriter, r *http.Request) error


// The APIStart function is a wrapper that adds an initialization ID to the context and handles any
// errors that occur during the execution of the provided function.
func APIStart(fn APISetting) http.HandlerFunc {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx = context.WithValue(ctx, "initID", rand.Intn(1000000))

		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}


type JSONAPIServer struct {
	address string
	network PricePlataform
}


// The NewJSONAPIServer function creates a new instance of a JSONAPIServer with the specified address
// and network.
func NewJSONAPIServer(address string, network PricePlataform) *JSONAPIServer {
	return &JSONAPIServer{
		address: address,
		network: network,
	}
}


// The `Run` method of the `JSONAPIServer` struct starts the HTTP server and listens for incoming
// requests on the specified address. It uses the `http.ListenAndServe` function to start the server
// and pass in the address and `nil` as the handler.
func (jas *JSONAPIServer) Run() {
	http.ListenAndServe(jas.address, nil)
}


type PriceResponse struct {
	Ticket string `json:"ticket"`
	Price  []int64  `json:"price"`
}


// The `HandleFeaturedProduct` function is a method of the `JSONAPIServer` struct. It handles the HTTP
// request for the featured product by extracting the `ticket` parameter from the URL query string. If
// the `ticket` is empty, it returns an error indicating that the ticket is invalid.
func (jas *JSONAPIServer) HandleFeaturedProduct(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticket := r.URL.Query().Get("ticket")
	if len(ticket) == 0 {
		return fmt.Errorf("Ticket invalid")
	}

	price, err := jas.network.FeaturedProduct(ctx, ticket)
	if err != nil {
		return err
	}

	resp := PriceResponse{
		Ticket: ticket, 
		Price: price,
	}

	return writeJSON(w, http.StatusOK, resp)
}


// The function writes a JSON response with a specified status code and value to an HTTP response
// writer.
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}