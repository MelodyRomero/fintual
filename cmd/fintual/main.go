package main

import (
	s "fintual/internal/services"
	"fmt"
	"log"
	"time"
)

func main() {
	var inventory []s.Stock

	dayStock1, err := time.Parse("020106 150405", "020106 150405")

	if err != nil {
		log.Fatal(err)
	}

	stock1 := s.NewStock("stock1", dayStock1, 0.80)

	inventory = append(inventory, stock1)

	dayStock2, err := time.Parse("020106 150405", "020109 150405")

	if err != nil {
		log.Fatal(err)
	}

	stock2 := s.NewStock("stock2", dayStock2, 0.70)

	inventory = append(inventory, stock1, stock2)

	miPortfolio := s.NewPortfolio(inventory)

	profit, err := miPortfolio.Profit("020106 150405", "020109 150405")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("El profit total corresponde a :" + fmt.Sprintf("%f", profit))
}
