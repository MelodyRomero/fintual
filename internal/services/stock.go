package services

import (
	e "fintual/internal/err"
	"time"
)

//Stock interface
type Stock interface {
	Price(date time.Time) (float64, error)
}

//Stock struct
type stock struct {
	Name      string
	Date      time.Time
	SalePrice float64
}

//NewStock construct
func NewStock(name string, date time.Time, price float64) Stock {

	return &stock{
		Name:      name,
		Date:      date,
		SalePrice: price,
	}
}

//Price func export
func (s *stock) Price(date time.Time) (float64, error) {

	if &date == nil {
		return 0.0, e.ErrGetStockPrice
	}

	if s.Date == date {
		return s.SalePrice, nil
	}

	return 0.0, nil
}
