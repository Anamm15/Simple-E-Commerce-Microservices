package algorithm

import (
	"math"
	"strings"
)

type ShippingOption struct {
	Courier       string
	Service       string
	Cost          int64
	EstimatedDays int32
}

type service struct {
	name       string
	multiplier float64
	baseETD    int32
}

type courier struct {
	code     string
	services []service
}

var cityZone = map[string]int{
	"jakarta":    1,
	"bogor":      1,
	"bandung":    2,
	"surabaya":   2,
	"denpasar":   3,
	"medan":      3,
	"balikpapan": 4,
	"makassar":   5,
}

var basePricePerKg = map[int]int64{
	1: 8000,
	2: 10000,
	3: 13000,
	4: 16000,
	5: 20000,
}

var etdAdjustment = map[int]int32{
	1: 0,
	2: 1,
	3: 2,
	4: 3,
	5: 4,
}

var couriers = []courier{
	{
		code: "JNE",
		services: []service{
			{name: "REG", multiplier: 1.0, baseETD: 3},
			{name: "YES", multiplier: 1.4, baseETD: 1},
		},
	},
	{
		code: "JNT",
		services: []service{
			{name: "REG", multiplier: 0.95, baseETD: 3},
			{name: "EXP", multiplier: 1.3, baseETD: 1},
		},
	},
	{
		code: "SICEPAT",
		services: []service{
			{name: "REG", multiplier: 0.9, baseETD: 3},
			{name: "BEST", multiplier: 1.25, baseETD: 1},
		},
	},
}

func CalculateShipping(
	originCity string,
	destinationCity string,
	totalWeightG int32,
) []ShippingOption {
	originZone, ok := cityZone[strings.ToLower(originCity)]
	if !ok {
		return nil
	}

	destZone, ok := cityZone[strings.ToLower(destinationCity)]
	if !ok {
		return nil
	}

	zoneDistance := math.Abs(float64(originZone-destZone)) + 1

	weightKg := math.Ceil(float64(totalWeightG))
	pricePerKg := basePricePerKg[int(zoneDistance)]

	baseCost := pricePerKg * int64(weightKg)

	var results []ShippingOption

	for _, c := range couriers {
		for _, s := range c.services {
			cost := int64(float64(baseCost) * s.multiplier)
			etd := s.baseETD + etdAdjustment[int(zoneDistance)]

			results = append(results, ShippingOption{
				Courier:       c.code,
				Service:       s.name,
				Cost:          cost,
				EstimatedDays: etd,
			})
		}
	}

	return results
}
