package handlers

import (
	"net/http"

	"github.com/goocarry/mstemplate/data"
)

// AddProduct ...
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(&prod)
}
