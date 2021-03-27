package main

import "time"

func main() {

	stop := false
	ledgers := make(chan Ledger, 1)

	go Produce(ledgers, &stop)
	go Consume(ledgers, []FormulaType{REVENUE, EXPENSES, GROSS_PROFIT_MARGIN, NET_PROFIT_MARGIN, WORKING_CAPITAL_RATIO})

	time.Sleep(100 * time.Millisecond)
	stop = true
	close(ledgers)

}
