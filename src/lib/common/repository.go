package common

type SQLComparator string

const (
	Like  SQLComparator = "like"
	Equal SQLComparator = "="
	ILike SQLComparator = "ilike"
)

type SQLCondition struct {
	Field      string
	Value      string
	Comparator SQLComparator
}

type SQLConditions []SQLCondition

var NoConditions = SQLConditions{}
