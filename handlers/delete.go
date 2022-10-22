package handlers

import (
	"net/http"
	"strconv"

	"github.com/goocarry/mstemplate/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} deleteProduct
// Deletes product from productList
// responses:
//   201: noContent

// DeleteProduct deletes a product from productList
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	// this will always convert because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Something went wrong...", http.StatusInternalServerError)
		return
	}
}