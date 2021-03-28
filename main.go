package main

import (
	"coding-challenge/formular"
	. "coding-challenge/models"
	"coding-challenge/workers"
	"time"
)

func main() {

	stop := false
	ledgers := make(chan Ledger, 1)

	go workers.Produce("data.json", ledgers, &stop)
	go workers.Consume(ledgers, []formular.FormulaType{formular.REVENUE, formular.EXPENSES, formular.GROSS_PROFIT_MARGIN, formular.NET_PROFIT_MARGIN, formular.WORKING_CAPITAL_RATIO})

	time.Sleep(100 * time.Millisecond)
	stop = true
	close(ledgers)

}
