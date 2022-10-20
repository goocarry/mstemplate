package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/goocarry/mstemplate/data"
	"github.com/gorilla/mux"
)

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

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct ...
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(&prod)
}

// UpdateProduct ...
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) 
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest) 
		return
	}

	p.l.Printf("Prod: %#v", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return 
	}

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return 
	}
}

// KeyProduct ...
type KeyProduct struct {}

// MiddlewareProductValidation ...
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest) 
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r =r.WithContext(ctx)
	
		next.ServeHTTP(rw, r)
	})
}
