package test

import (
	"coding-challenge/formular"
	. "coding-challenge/models"
	"testing"
)

func TestCalcRevenue(t *testing.T) {
	data1 := Data{AccountCategory: "revenue", TotalValue: 1}
	data2 := Data{AccountCategory: "revenue", TotalValue: 1}
	data3 := Data{AccountCategory: "revenue", TotalValue: 1}
	data4 := Data{AccountCategory: "revenue", TotalValue: 1}
	data5 := Data{AccountCategory: "expense", TotalValue: 1}
	datas := []interface{}{data1, data2, data3, data4, data5}
	AssertEqual(t, formular.CalcRevenue(datas), float64(4))
}

func TestCalcExpenses(t *testing.T) {
	data1 := Data{AccountCategory: "expense", TotalValue: 1}
	data2 := Data{AccountCategory: "expense", TotalValue: 1}
	data3 := Data{AccountCategory: "expense", TotalValue: 1}
	data4 := Data{AccountCategory: "expense", TotalValue: 1}
	data5 := Data{AccountCategory: "revenue", TotalValue: 1}
	datas := []interface{}{data1, data2, data3, data4, data5}
	AssertEqual(t, formular.CalcExpenses(datas), float64(4))
}

func TestCalcGrossProfitMargin(t *testing.T) {
	data1 := Data{AccountType: "sales", ValueType: "debit", TotalValue: 1}
	data2 := Data{AccountType: "sales", ValueType: "debit", TotalValue: 1}
	data3 := Data{AccountType: "sales", ValueType: "debit", TotalValue: 1}
	data4 := Data{AccountType: "sales", ValueType: "debit", TotalValue: 1}
	data5 := Data{AccountType: "whatever", ValueType: "whatever", TotalValue: 1}
	datas := []interface{}{data1, data2, data3, data4, data5}
	AssertEqual(t, formular.CalcGrossProfitMargin(datas, 8), float64(4)/float64(8))
}

func TestCalcNetProfitMargin(t *testing.T) {
	AssertEqual(t, formular.CalcNetProfitMargin(8, 4), (float64(8)-float64(4))/float64(8))
}

func TestCalcWorkingCapitalRatio(t *testing.T) {
	assetsAddData1 := Data{AccountType: "current", AccountCategory: "assets", ValueType: "debit", TotalValue: 2}
	assetsAddData2 := Data{AccountType: "bank", AccountCategory: "assets", ValueType: "debit", TotalValue: 2}
	assetsAddData3 := Data{AccountType: "current_accounts_receivable", AccountCategory: "assets", ValueType: "debit", TotalValue: 2}
	assetsSubData1 := Data{AccountType: "current", AccountCategory: "assets", ValueType: "credit", TotalValue: 1}
	assetsSubData2 := Data{AccountType: "bank", AccountCategory: "assets", ValueType: "credit", TotalValue: 1}
	assetsSubData3 := Data{AccountType: "current_accounts_receivable", AccountCategory: "assets", ValueType: "credit", TotalValue: 1}

	liaAddData1 := Data{AccountType: "current", AccountCategory: "liability", ValueType: "credit", TotalValue: 2}
	liaAddData2 := Data{AccountType: "bank", AccountCategory: "liability", ValueType: "whatever", TotalValue: 2}
	liaAddData3 := Data{AccountType: "current_accounts_payable", AccountCategory: "liability", ValueType: "credit", TotalValue: 2}
	liaSubData1 := Data{AccountType: "current", AccountCategory: "liability", ValueType: "debit", TotalValue: 1}
	liaSubData2 := Data{AccountType: "bank", AccountCategory: "liability", ValueType: "whatever", TotalValue: 1}
	liaSubData3 := Data{AccountType: "current_accounts_payable", AccountCategory: "liability", ValueType: "debit", TotalValue: 1}

	datas := []interface{}{assetsAddData1, assetsAddData2, assetsAddData3,
		assetsSubData1, assetsSubData2, assetsSubData3,
		liaAddData1, liaAddData2, liaAddData3,
		liaSubData1, liaSubData2, liaSubData3}
	AssertEqual(t, formular.CalcWorkingCapitalRatio(datas), (float64((2+2+2)-(1+1+1)))/float64((2+2)-(1+1)))
}
