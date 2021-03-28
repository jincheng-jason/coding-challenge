package models

import "time"

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
