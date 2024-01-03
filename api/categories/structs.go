package categories

type CreateCategoryReq struct {
	Label     string `json:"label" validate:"min=1,max=24"`
	Icon      string `json:"icon"`
	IsExpense *bool  `json:"is_expense" validate:"nonnil"`
}

type UpdateCategoryReq struct {
	ID        string  `json:"id" validate:"len=36"`
	Label     *string `json:"label" validate:"min=1,max=24"`
	Icon      *string `json:"icon"`
	IsExpense *bool   `json:"is_expense"`
}
