package model

import "time"

type Stats []Stat

type Stat struct {
	Total    int
	From     time.Time
	To       time.Time
	Employee *Employee
	Budgets  Budgets
}

type Budgets []Budget

type Budget struct {
	Count int
	Type  BudgetType
}

type BudgetType uint

func (t BudgetType) String() string {
	return budgetTypeNames[t]
}

const (
	BudgetTypeLowerThan50K BudgetType = iota
	BudgetTypeBetween50KAnd100K
	BudgetTypeBetween100KAnd500K
	BudgetTypeGreaterThan500K
)

var budgetTypeNames = map[BudgetType]string{
	BudgetTypeLowerThan50K:       "Lower than 50K",
	BudgetTypeBetween50KAnd100K:  "Between 50K and 100K",
	BudgetTypeBetween100KAnd500K: "Between 100K and 500K",
	BudgetTypeGreaterThan500K:    "Greater than 500K",
}
