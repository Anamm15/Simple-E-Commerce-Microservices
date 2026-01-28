package algorithm

import (
	"errors"
)

var PriceTable = map[string]map[string]int32{
	"JNE": {
		"kargo":     10000,
		"regular":   15000,
		"exclusive": 25000,
	},
	"JNT": {
		"kargo":     9000,
		"regular":   14000,
		"exclusive": 23000,
	},
}

// Estimated days delivery (custom)
var EstimatedDaysTable = map[string]map[string]int32{
	"JNE": {
		"kargo":     5,
		"regular":   3,
		"exclusive": 1,
	},
	"JNT": {
		"kargo":     4,
		"regular":   2,
		"exclusive": 1,
	},
}

type ShippingOption struct {
	Courier       string `json:"courier"`
	ServiceType   string `json:"service_type"`
	WeightKG      int32  `json:"weight_kg"`
	PricePerKG    int32  `json:"price_per_kg"`
	TotalCost     int32  `json:"total_cost"`
	EstimatedDays int32  `json:"estimated_days"`
}

func CalculateAllCosts(weightKG int32) ([]ShippingOption, error) {
	if weightKG <= 0 {
		return nil, errors.New("weight must be greater than zero")
	}

	if weightKG > 10 {
		return nil, errors.New("weight exceeds maximum limit of 10 kg")
	}

	var options []ShippingOption

	for courier, services := range PriceTable {
		for service, pricePerKG := range services {

			estimatedDays := EstimatedDaysTable[courier][service]
			totalCost := pricePerKG * weightKG

			options = append(options, ShippingOption{
				Courier:       courier,
				ServiceType:   service,
				WeightKG:      weightKG,
				PricePerKG:    pricePerKG,
				TotalCost:     totalCost,
				EstimatedDays: estimatedDays,
			})
		}
	}

	return options, nil
}
