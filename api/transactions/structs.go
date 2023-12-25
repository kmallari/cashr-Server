package transactions

type CreateTransactionReq struct {
	Date        int64   `json:"date" validate:"min=1"`
	Amount      float64 `json:"amount" validate:"min=0.01"`
	IsExpense   *bool   `json:"is_expense" validate:"nonnil"`
	CategoryID  string  `json:"category" validate:"len=36"`
	Description string  `json:"description" validate:"max=250"`
}

type UpdateTransactionReq struct {
	Date        int64   `json:"date,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	Description string  `json:"description,omitempty"`
	IsExpense   bool    `json:"is_expense,omitempty"`
	Amount      string `json:"amount,omitempty"`
}
