package main

import (
	"github.com/sirupsen/logrus"
	"context"
	"fmt"
	"time"
)


var pricesProducts = map[string]int64{
	"NATURAL_EARTH_DOG_3KG":              			17900,
	"NATURAL_EARTH_DOG_15KG":             			56900,
	"ARENA_NATURAL_EARTH_ECOLOGIC_4.6KG": 			8900,
	"BATHROOM_FOR_CAT_WHALE_BLUE":        			29900,
	"EVOLVE_CAT_CLASSIC_DEBONED_CHICKEN_1.36KG": 	9500,
	"NATURAL_EARTH_DOG_ADULT_3KG":              	17900,
}


type PricePlataform interface {
	FeaturedProduct(context.Context, string) ([]int64, error)
}


type PricePlataforms struct {
	PricePlataform
}


// The `func (pp *PricePlataforms) FeaturedProduct(_ context.Context, ticket string) (int64, error)`
// function is implementing the `FeaturedProduct` method of the `PricePlataform` interface for the
// `PricePlataforms` struct.
func (pp *PricePlataforms) FeaturedProduct(_ context.Context, ticket string) (int64, error) {
	price, ok := pricesProducts[ticket]
	if !ok {
		return 0, fmt.Errorf("Price for ticket %s not found", ticket)
	}
	return price, nil
}


type LogPlataform struct {
	PricePlataform
}


// The `func (lp LogPlataform) FeaturedProduct(ctx context.Context, ticket string) (price int64, err
// error)` function is implementing the `FeaturedProduct` method of the `PricePlataform` interface for
// the `LogPlataform` struct.
func (lp LogPlataform) FeaturedProduct(ctx context.Context, ticket string) (price int64, err error) {
	defer func(begin time.Time) {
		initID := ctx.Value("initialID")

		logrus.WithFields(logrus.Fields{
			"initialID": initID,
			"session": time.Since(begin),
			"ticket": ticket,
			"price": price,
			"err": err,
		}).Info("FeaturedProduct")
	}(time.Now())

	return lp.FeaturedProduct(ctx, ticket)
}