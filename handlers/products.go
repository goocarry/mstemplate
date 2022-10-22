// Package handlers classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
//
// Consumes:
// -application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/goocarry/mstemplate/data"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productsNoContent struct {}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}


// Replace all of the code below by gorilla mux
// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	// handle the request for a list of products
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		// expect the id in the URI
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1) 

// 		if len(g) != 1 {
// 			p.l.Println("Invalid URI more than one id")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		if len(g[0]) != 2 {
// 			p.l.Println("Invalid URI more than one capture group")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			p.l.Println("Invalid URI unable to convert to number")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return 
// 		}
// 		p.updateProduct(id, rw, r)

// 	}

// 	// catch all
// 	// if no method is satisfied return an error
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }




// KeyProduct ...
type KeyProduct struct {}

// MiddlewareProductValidation ...
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product")
			http.Error(rw, "Error reading product", http.StatusBadRequest) 
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product")
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err) , http.StatusBadRequest) 
			return
		}


		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r =r.WithContext(ctx)
	
		next.ServeHTTP(rw, r)
	})
}
