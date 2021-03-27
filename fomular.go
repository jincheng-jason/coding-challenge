package main

func CalcRevenue(datas []interface{}) float64 {
	return Stream(datas).Filter(func(each interface{}) bool {
		return each.(Data).AccountCategory == "revenue"
	}).Reduce(0.0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(float64) + cur.(Data).TotalValue
	}).(float64)
}

func CalcExpenses(datas []interface{}) float64 {
	return Stream(datas).Filter(func(each interface{}) bool {
		return each.(Data).AccountCategory == "expense"
	}).Reduce(0.0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(float64) + cur.(Data).TotalValue
	}).(float64)
}

func CalcGrossProfitMargin(datas []interface{}, revenue float64) float64 {
	gross := Stream(datas).Filter(func(each interface{}) bool {
		return each.(Data).AccountType == "sales" && each.(Data).ValueType == "debit"
	}).Reduce(0.0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(float64) + cur.(Data).TotalValue
	})
	return gross.(float64) / revenue
}

func CalcNetProfitMargin(revenue float64, expense float64) float64 {
	return (revenue - expense) / revenue
}

func CalcWorkingCapitalRatio(datas []interface{}) float64 {
	assetsAccountTypes := []string{"current", "bank", "current_accounts_receivable"}
	assetsAdd := Stream(datas).Filter(func(each interface{}) bool {
		return each.(Data).AccountCategory == "assets" && each.(Data).ValueType == "debit" && elementInSlice(each.(Data).AccountType, assetsAccountTypes)
	}).Reduce(0.0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(float64) + cur.(Data).TotalValue
	})

	assetsSub := Stream(datas).Filter(func(each interface{}) bool {
		return each.(Data).AccountCategory == "assets" && each.(Data).ValueType == "credit" && elementInSlice(each.(Data).AccountType, assetsAccountTypes)
	}).Reduce(0.0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(float64) + cur.(Data).TotalValue
	})

	assets := assetsAdd.(float64) - assetsSub.(float64)

	liabilitiesAccountTypes := []string{"current", "current_accounts_payable"}
	liaAdd := Stream(datas).Filter(func(each interface{}) bool {
		return each.(Data).AccountCategory == "liability" && each.(Data).ValueType == "credit" && elementInSlice(each.(Data).AccountType, liabilitiesAccountTypes)
	}).Reduce(0.0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(float64) + cur.(Data).TotalValue
	})

	liaSub := Stream(datas).Filter(func(each interface{}) bool {
		return each.(Data).AccountCategory == "liability" && each.(Data).ValueType == "debit" && elementInSlice(each.(Data).AccountType, liabilitiesAccountTypes)
	}).Reduce(0.0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(float64) + cur.(Data).TotalValue
	})

	liabilities := liaAdd.(float64) - liaSub.(float64)

	return assets / liabilities
}

func elementInSlice(target string, strArray []string) bool {
	for _, e := range strArray {
		if target == e {
			return true
		}
	}
	return false
}
