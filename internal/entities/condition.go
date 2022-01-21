package entities

type ConditionType int

const (
	_ ConditionType = iota
	ConditionTypeEq
	ConditionTypeNotEq
)
