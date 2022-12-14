package handlers

import (
	"context"
	"net/http"

	"github.com/goocarry/mstemplate/currency/protos/currency"
	"github.com/goocarry/mstemplate/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

// ListAll handles GET request and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	rw.Header().Add("Content-Type", "application/json")
	
	lp := data.GetProducts()

	err := data.ToJSON(lp, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return single product
// responses: 
// 	200: productResponse
// 	400: errorResponse 

// ListSingle handles GET request and returns product
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("[DEBUG] get record id", id)

	lp, err := data.GetProductByID(id)

	switch err {
		case nil:
		
		case data.ErrProductNotFound:
			p.l.Println("[ERROR] fetching product", err)

			rw.WriteHeader(http.StatusNotFound)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

	// get exchange rate
	rr := &currency.RateRequest{
		Base: currency.Currencies(currency.Currencies_value["EUR"]),
		Destination: currency.Currencies(currency.Currencies_value["GBP"]),
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[ERROR] error getting new rate", err)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

	lp.Price =  lp.Price * resp.Rate

	err = data.ToJSON(lp, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

