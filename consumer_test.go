package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestConsume(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ledgers := make(chan Ledger, 1)
	go Consume(ledgers, []FormulaType{REVENUE, EXPENSES, GROSS_PROFIT_MARGIN, NET_PROFIT_MARGIN, WORKING_CAPITAL_RATIO})
	ledger := Ledger{Data: []Data{{AccountCategory: "revenue", TotalValue: 10000},
		{AccountCategory: "expense", TotalValue: 10000},
		{AccountType: "sales", ValueType: "debit", TotalValue: 1},
		{AccountType: "current", AccountCategory: "assets", ValueType: "debit", TotalValue: 2},
		{AccountType: "current", AccountCategory: "liability", ValueType: "credit", TotalValue: 2}}}
	ledgers <- ledger
	time.Sleep(100 * time.Millisecond)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	AssertEqual(t, string(out), "Revenue: $10,000.000 \nExpenses: $10,000.000 \nGross Profit Margin: 0% \nNet Profit Margin: 0% \nWorking Capital Ratio: 100% \n")
}
