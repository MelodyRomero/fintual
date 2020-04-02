package services

import (
	e "fintual/internal/err"
	"log"
	"time"
)

//Portfolio interface
type Portfolio interface {
	Profit(from, to string) (float64, error)
}

//Portfolio struct
type portfolio struct {
	Stocks []Stock
}

//NewPortfolio construct
func NewPortfolio(stocks []Stock) Portfolio {

	return &portfolio{
		Stocks: stocks,
	}
}

//Profit function export
func (p *portfolio) Profit(from, to string) (float64, error) {
	var profit float64
	dateRange, err := getDateRange(from, to)

	if err != nil {
		return 0.0, e.ErrDateRange
	}

	for _, stock := range p.Stocks {
		for _, day := range dateRange {

			dayprofit, err := stock.Price(day)
			if err != nil {
				return 0.0, e.ErrGetStockPrice
			}
			profit = profit + dayprofit
		}
	}

	return profit, nil
}

//Receive format day month year min sec ms
func getDateRange(from, to string) ([]time.Time, error) {

	initRange, err := time.Parse("020106 150405", from)

	if err != nil {
		log.Fatal(err)
	}

	endRange, err := time.Parse("020106 150405", to)

	if err != nil {
		log.Fatal(err)
	}

	var dateRange []time.Time

	dateRange = append(dateRange, initRange)

	for id, d := range dateRange {

		dateRange = append(dateRange, d.AddDate(0, 0, 1))

		if dateRange[id].AddDate(0, 0, 1) == endRange {
			break
		}
	}
	return dateRange, nil
}
