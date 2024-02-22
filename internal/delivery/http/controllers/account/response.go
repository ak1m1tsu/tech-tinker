package account

type InfoResponse struct {
	ID        string         `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Role      string         `json:"role"`
	CreatedAt string         `json:"created_at"`
	Orders    OrderResponses `json:"orders,omitempty"`
}

type OrderResponses []OrderResponse

type OrderResponse struct {
	ID         string `json:"id"`
	Number     int    `json:"number"`
	PriceLimit int    `json:"price_limit"`
	Comment    string `json:"comment"`
	Address    string `json:"address"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}

type StatisticResponse struct {
	From    string          `json:"from"`
	To      string          `json:"to"`
	Total   int             `json:"total"`
	Budgets BudgetResponses `json:"budgets"`
}

type BudgetResponses []BudgetResponse

type BudgetResponse struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}
