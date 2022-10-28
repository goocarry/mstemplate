package handlers

import (
	"net/http"
	"strconv"

	"github.com/goocarry/mstemplate/data"
	"github.com/gorilla/mux"
)

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

	err = data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return 
	}

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return 
	}
}