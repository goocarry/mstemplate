package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/goocarry/mstemplate/currency/protos/currency"
	"github.com/goocarry/mstemplate/data"
	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct {}

// Products handler for getting and updating products
type Products struct {
	l *log.Logger
	v *data.Validation
	cc currency.CurrencyClient
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation, cc currency.CurrencyClient) *Products {
	return &Products{l: l, v: v, cc: cc}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}


// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product if from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
