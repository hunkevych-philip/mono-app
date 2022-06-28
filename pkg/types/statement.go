package types

type Statement struct {
	StatementRecords []*StatementRecord `json:"statementRecords,omitempty"`
}

type StatementRecord struct {
	Time            int    `json:"time"`
	Description     string `json:"description"`
	Amount          int    `json:"amount"`          // actual price
	OperationAmount int    `json:"operationAmount"` // initial price
	Balance         int    `json:"balance"`
}
