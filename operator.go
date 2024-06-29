package query

const (
	Eq Operator = iota + 1
	Ne
	Gt
	Gte
	Lt
	Lte
	In
	NotIn
)

type Operator int

func (o Operator) String() string {
	switch o {
	case Eq:
		return "eq"
	case Ne:
		return "ne"
	case Gt:
		return "gt"
	case Gte:
		return "gte"
	case Lt:
		return "lt"
	case Lte:
		return "lte"
	case In:
		return "in"
	case NotIn:
		return "nin"
	default:
		return "undefined"
	}
}
