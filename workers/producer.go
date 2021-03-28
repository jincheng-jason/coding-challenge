package workers

import (
	. "coding-challenge/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func Produce(filePath string, ledgers chan<- Ledger, stop *bool) {
	for !*stop {
		ledger, err := readFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		ledgers <- ledger
		time.Sleep(1000 * time.Millisecond)
	}
}

func readFile(filePath string) (Ledger, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return Ledger{}, err
	}
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)
	var ledger Ledger
	jsonErr := json.Unmarshal(byteValue, &ledger)
	if jsonErr != nil {
		return Ledger{}, jsonErr
	}
	return ledger, nil
}
