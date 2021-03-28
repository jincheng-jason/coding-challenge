package workers

import (
	"coding-challenge/formular"
	. "coding-challenge/models"
)

func Consume(ledgers <-chan Ledger, formulas []formular.FormulaType) {
	for ledger := range ledgers {
		originDatas := ledger.Data
		datas := make([]interface{}, len(originDatas))
		for i := range originDatas {
			datas[i] = originDatas[i]
		}
		for _, formula := range formulas {
			var result float64
			switch formula {
			case formular.REVENUE:
				result = formular.CalcRevenue(datas)
			case formular.EXPENSES:
				result = formular.CalcExpenses(datas)
			case formular.GROSS_PROFIT_MARGIN:
				result = formular.CalcGrossProfitMargin(datas, formular.CalcRevenue(datas))
			case formular.NET_PROFIT_MARGIN:
				result = formular.CalcNetProfitMargin(formular.CalcRevenue(datas), formular.CalcExpenses(datas))
			case formular.WORKING_CAPITAL_RATIO:
				result = formular.CalcWorkingCapitalRatio(datas)
			}
			PrintResult(formula, result)
		}
	}
}
