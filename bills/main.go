package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func newBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter the name of the bill: ", reader)

	b := createBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b *Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		floatPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Invalid price. Please enter a valid number.")
			promptOptions(b)
		}
		b.addItem(name, floatPrice)

		fmt.Printf("item added: %v -  %v \n", name, price)
		promptOptions(b)
	case "t":
		tipPercent, _ := getInput("Tip percent: ", reader)

		percentFloat, err := strconv.ParseFloat(tipPercent, 64)
		if err != nil {
			fmt.Println("Invalid tip percent. Please enter a valid number.")
			promptOptions(b)
		}
		b.updateTipPercent(percentFloat)

		fmt.Printf("tip added: %v \n", percentFloat)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("Bill saved to file - ", b.name)
	default:
		fmt.Println("Invalid option. Please choose a valid option.")
		promptOptions(b)
	}
}

func main() {
	myBill := newBill()
	promptOptions(&myBill)

	fmt.Println(myBill.format())
}
