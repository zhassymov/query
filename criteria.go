package query

// Criteria represents a query criteria.
type Criteria[T comparable] struct {
	op    Operator
	value T
}

func (c Criteria[T]) Operator() Operator {
	return c.op
}

func (c Criteria[T]) Value() T {
	return c.value
}
