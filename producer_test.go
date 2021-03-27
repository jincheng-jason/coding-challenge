package main

import (
	"testing"
	"time"
)

func TestProduce(t *testing.T) {
	ledgers := make(chan Ledger, 1)
	stop := false
	go Produce(ledgers, &stop)
	time.Sleep(1000 * time.Millisecond)
	l := <-ledgers
	if len(l.Data) == 0 {
		t.Errorf("Producer expected produce ledger, but got empty")
	}
	stop = true
	close(ledgers)
}
