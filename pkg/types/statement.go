package types

type Statement struct {
	Time            int    `json:"time"`
	Description     string `json:"description"`
	Amount          int    `json:"amount"`          // actual price
	OperationAmount int    `json:"operationAmount"` // initial price
	Balance         int    `json:"balance"`
}
