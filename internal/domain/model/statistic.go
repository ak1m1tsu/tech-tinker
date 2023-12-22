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
	BudgetType_LowerThan50K BudgetType = iota
	BudgetType_Between50KAnd100K
	BudgetType_Between100KAnd500K
	BudgetType_Greater500K
)

var budgetTypeNames = map[BudgetType]string{
	BudgetType_LowerThan50K:       "Lower than 50K",
	BudgetType_Between50KAnd100K:  "Between 50K and 100K",
	BudgetType_Between100KAnd500K: "Between 100K and 500K",
	BudgetType_Greater500K:        "Greater than 500K",
}
