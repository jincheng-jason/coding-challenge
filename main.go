package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"
)

type Ledger struct {
	ObjectCategory       string    `json:"object_category"`
	ConnectionID         string    `json:"connection_id"`
	User                 string    `json:"user"`
	ObjectCreationDate   time.Time `json:"object_creation_date"`
	Data                 []Data    `json:"data"`
	Currency             string    `json:"currency"`
	ObjectOriginType     string    `json:"object_origin_type"`
	ObjectOriginCategory string    `json:"object_origin_category"`
	ObjectType           string    `json:"object_type"`
	ObjectClass          string    `json:"object_class"`
	BalanceDate          time.Time `json:"balance_date"`
}
type Data struct {
	AccountCategory   string  `json:"account_category"`
	AccountCode       string  `json:"account_code"`
	AccountCurrency   string  `json:"account_currency"`
	AccountIdentifier string  `json:"account_identifier"`
	AccountStatus     string  `json:"account_status"`
	ValueType         string  `json:"value_type"`
	AccountName       string  `json:"account_name"`
	AccountType       string  `json:"account_type"`
	AccountTypeBank   string  `json:"account_type_bank"`
	SystemAccount     string  `json:"system_account"`
	TotalValue        float64 `json:"total_value"`
}

func main() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ledger Ledger

	json.Unmarshal(byteValue, &ledger)

	for i := 0; i < len(ledger.Data); i++ {
		fmt.Println("AccountCategory: " + ledger.Data[i].AccountCategory)
		fmt.Println("AccountCode: " + ledger.Data[i].AccountCode)
		fmt.Println("AccountCurrency: " + ledger.Data[i].AccountCurrency)
		fmt.Println("AccountIdentifier: " + ledger.Data[i].AccountIdentifier)
		fmt.Println("AccountStatus: " + ledger.Data[i].AccountStatus)
		fmt.Println("ValueType: " + ledger.Data[i].ValueType)
		fmt.Println("AccountName: " + ledger.Data[i].AccountName)
		fmt.Println("AccountType: " + ledger.Data[i].AccountType)
		fmt.Println("AccountTypeBank: " + ledger.Data[i].AccountTypeBank)
		fmt.Println("SystemAccount: " + ledger.Data[i].SystemAccount)
		fmt.Println("TotalValue: " + strconv.FormatFloat(ledger.Data[i].TotalValue, 'f', -1, 64))
		fmt.Println("--------------------")
	}

	//Revenue

	//Expenses

	//Gross Profit Margin

	//Net Profit Margin

	//Working Capital Ratio

}

type stream struct {
	list []interface{}
}

func Stream(arrs ...interface{}) *stream {

	st := new(stream)

	if len(arrs) > 0 {
		if x, ok := arrs[0].([]interface{}); ok {
			st.list = make([]interface{}, len(x))
			copy(st.list, x)
		} else {
			st.list = make([]interface{}, len(arrs))
			copy(st.list, arrs)
		}
	}

	return st
}

func (s *stream) Filter(fn func(each interface{}) bool) *stream {
	list := make([]interface{}, 0, len(s.list))
	for _, x := range s.list {
		if fn(x) {
			list = append(list, x)
		}
	}
	s.list = list
	return s
}

func (s *stream) ForEach(fn func(each interface{})) {
	list := s.list
	for _, x := range list {
		fn(x)
	}
}

func (s *stream) Collect(r interface{}) {
	bytes, _ := json.Marshal(s.list)
	json.Unmarshal(bytes, &r)
}

func (s *stream) FindAny() (interface{}, bool) {
	if len(s.list) > 0 {
		return s.list[0], true
	}
	return nil, false
}

func (s *stream) AnyMatch(fn func(each interface{}) bool) bool {
	for _, x := range s.list {
		if fn(x) {
			return true
		}
	}
	return false
}

func (s *stream) Map(fn func(each interface{}) interface{}) *stream {
	for i, x := range s.list {
		s.list[i] = fn(x)
	}
	return s
}

func (s *stream) Count() int {
	return len(s.list)
}

func (s *stream) Distinct() []interface{} {
	m := make(map[interface{}][]interface{})
	for _, x := range s.list {
		m[x] = nil
	}

	r := make([]interface{}, 0, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (s *stream) GroupByInt(fn func(each interface{}) int64, r interface{}) {
	m := make(map[int64][]interface{})

	for _, x := range s.list {
		key := fn(x)
		if l, ok := m[key]; ok {
			l = append(l, x)
			m[key] = l
		} else {
			l = make([]interface{}, 0, 0)
			l = append(l, x)
			m[key] = l
		}
	}
	bytes, _ := json.Marshal(m)
	_ = json.Unmarshal(bytes, &r)
}

func (s *stream) GroupByString(fn func(each interface{}) string, r interface{}) {
	m := make(map[string][]interface{})

	for _, x := range s.list {
		key := fn(x)
		if l, ok := m[key]; ok {
			l = append(l, x)
			m[key] = l
		} else {
			l = make([]interface{}, 0, 0)
			l = append(l, x)
			m[key] = l
		}
	}
	bytes, _ := json.Marshal(m)
	_ = json.Unmarshal(bytes, &r)

}

func (s *stream) Sum(fn func(each interface{}) interface{}) float64 {
	var r float64 = 0
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = r + (float64)(p.(int))
			break
		case float64:
			r = r + p.(float64)
			break
		}
	}
	return r
}

func (s *stream) Average(fn func(each interface{}) interface{}) float64 {
	var r float64 = 0
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = r + (float64)(p.(int))
			break
		case float64:
			r = r + p.(float64)
			break
		}
	}
	return r / float64(len(s.list))
}

func (s *stream) Max(fn func(each interface{}) interface{}) float64 {
	var r float64 = math.MinInt64
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = math.Max(r, (float64)(p.(int)))
			break
		case float64:
			r = math.Max(r, p.(float64))
			break
		}
	}
	return r
}

func (s *stream) Min(fn func(each interface{}) interface{}) float64 {
	var r = math.MaxFloat64
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = math.Min(r, (float64)(p.(int)))
			break
		case float64:
			r = math.Min(r, p.(float64))
			break
		}
	}
	return r
}

func (s *stream) Reduce(initialValue interface{}, fn func(pre interface{}, cur interface{}) interface{}) interface{} {
	for i := 0; i < len(s.list); i++ {
		initialValue = fn(initialValue, s.list[i])
	}
	return initialValue
}
