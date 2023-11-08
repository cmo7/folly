package common

type SQLComparator string

const (
	Like             SQLComparator = "like"
	Equal            SQLComparator = "="
	ILike            SQLComparator = "ilike"
	GreaterThan      SQLComparator = ">"
	LessThan         SQLComparator = "<"
	GreaterEqualThan SQLComparator = ">="
	LessEqualThan    SQLComparator = "<="
	In               SQLComparator = "in"
)

type SQLConditionType string

const (
	And SQLConditionType = "and"
	Or  SQLConditionType = "or"
)

type SQLCondition struct {
	Type       SQLConditionType
	Field      string
	Value      string
	Comparator SQLComparator
}

type SQLConditions []SQLCondition

var NoConditions = SQLConditions{}

type OrderDirection string

const (
	Asc  OrderDirection = "asc"
	Desc OrderDirection = "desc"
)

type OrderBy struct {
	Field     string
	Direction OrderDirection
}

type OrderBys []OrderBy
