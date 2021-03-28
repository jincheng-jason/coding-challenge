package workers

import (
	"coding-challenge/formular"
	"fmt"
	"github.com/leekchan/accounting"
	"math"
)

func PrintResult(formula formular.FormulaType, result float64) {
	switch formula {
	case formular.REVENUE, formular.EXPENSES:
		ac := accounting.Accounting{Symbol: "$", Precision: 3}
		fmt.Printf("%v: %v \n", formula, ac.FormatMoney(result))
	case formular.GROSS_PROFIT_MARGIN, formular.NET_PROFIT_MARGIN, formular.WORKING_CAPITAL_RATIO:
		fmt.Printf("%v: %v%% \n", formula, roundTo(result*100, 1))
	}
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
