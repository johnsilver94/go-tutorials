package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name       string
	items      map[string]float64
	tipPercent float64
}

func createBill(name string) Bill {
	b := Bill{
		name:       name,
		items:      map[string]float64{},
		tipPercent: 0.0,
	}

	return b
}

func (bill *Bill) format() string {
	fs := "Bill breakdown for:\n"
	var total float64 = 0
	var tip float64 = 0

	for item, price := range bill.items {
		fs += fmt.Sprintf("%-25v ...%v$ \n", item+":", price)
		total += price
	}

	tip = (total * bill.tipPercent) / 100
	fs += fmt.Sprintf("%-25v ...%0.2f$ \n", "Tip:", tip)
	total += tip

	fs += fmt.Sprintf("%-25v ...%0.2f$", "Total:", total)

	return fs
}

func (bill *Bill) updateTipPercent(tipPercent float64) {
	bill.tipPercent = tipPercent
}

func (bill *Bill) addItem(name string, price float64) {
	bill.items[name] = price
}

func (bill *Bill) save() {
	data := []byte(bill.format())

	err := os.WriteFile("data/"+bill.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill saved successfully.")
}
