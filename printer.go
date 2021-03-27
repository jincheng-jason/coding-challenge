package main

import (
	"fmt"
	"github.com/leekchan/accounting"
	"math"
)

func PrintResult(formula FormulaType, result float64) {
	switch formula {
	case REVENUE, EXPENSES:
		ac := accounting.Accounting{Symbol: "$", Precision: 3}
		fmt.Printf("%v: %v \n", formula, ac.FormatMoney(result))
	case GROSS_PROFIT_MARGIN, NET_PROFIT_MARGIN, WORKING_CAPITAL_RATIO:
		fmt.Printf("%v: %v%% \n", formula, roundTo(result*100, 1))
	}
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
