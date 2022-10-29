package server

import (
	"context"

	"github.com/goocarry/mstemplate/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

// Currency ...
type Currency struct {
	log hclog.Logger
	currency.UnimplementedCurrencyServer
}

// NewCurrency ...
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{
		log: l,
	}
}

// GetRate ...
func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &currency.RateResponse{Rate: 0.5}, nil
}