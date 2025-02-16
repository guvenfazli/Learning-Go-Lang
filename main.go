package main

import "fmt"

func main() {

	enteredRevenue, enteredExpenses, enteredTaxRate := gatherUserValues("Revenue: ", "Expenses: ", "Tax Rate: ")

	ebt, profit, ratio := calculateWithValues(enteredRevenue, enteredExpenses, enteredTaxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}

func gatherUserValues(revenueMessage, expensesMessage, taxRateMessage string) (enteredRevenue, enteredExpenses, enteredTaxRate float64) {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print(revenueMessage)
	fmt.Scan(&revenue)
	enteredRevenue = revenue

	fmt.Print(expensesMessage)
	fmt.Scan(&expenses)
	enteredExpenses = expenses

	fmt.Print(taxRateMessage)
	fmt.Scan(&taxRate)
	enteredTaxRate = taxRate

	return enteredRevenue, enteredExpenses, enteredTaxRate

}

func calculateWithValues(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
