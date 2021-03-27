package main

func Consume(ledgers <-chan Ledger, formulas []FormulaType) {
	for ledger := range ledgers {
		originDatas := ledger.Data
		datas := make([]interface{}, len(originDatas))
		for i := range originDatas {
			datas[i] = originDatas[i]
		}
		for _, formula := range formulas {
			var result float64
			switch formula {
			case REVENUE:
				result = CalcRevenue(datas)
			case EXPENSES:
				result = CalcExpenses(datas)
			case GROSS_PROFIT_MARGIN:
				result = CalcGrossProfitMargin(datas, CalcRevenue(datas))
			case NET_PROFIT_MARGIN:
				result = CalcNetProfitMargin(CalcRevenue(datas), CalcExpenses(datas))
			case WORKING_CAPITAL_RATIO:
				result = CalcWorkingCapitalRatio(datas)
			}
			PrintResult(formula, result)
		}
	}
}
