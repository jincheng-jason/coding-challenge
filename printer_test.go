package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPrintResult(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	PrintResult(REVENUE, 10000)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	AssertEqual(t, string(out), "Revenue: $10,000.000 \n")
}

func TestPrintResultWithGROSS_PROFIT_MARGIN(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	PrintResult(GROSS_PROFIT_MARGIN, 0.2388)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	AssertEqual(t, string(out), "Gross Profit Margin: 23.9% \n")
}
