package query

import (
	"cmp"
	"errors"
	"strconv"

	"github.com/zhassymov/please"
)

type Query struct {
	offset   int
	limit    int
	cursor   string
	criteria map[string][]Criteria[any]
}

// New creates a new query.
func New(opts ...please.Validate[*Query]) (Query, error) {
	q := Query{criteria: make(map[string][]Criteria[any], len(opts))}
	return q, please.Join(&q, opts...)
}

// Offset returns the offset value and the set indication.
func (q *Query) Offset() (int, bool) {
	return q.offset, q.offset > 0
}

// Offset sets the offset value with validation.
func Offset(offset int, opts ...please.Validate[int]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(offset, opts...); err != nil {
			return err
		}
		q.offset = offset
		return nil
	}
}

// OffsetString sets the offset value with validation from a string.
func OffsetString(offset string, opts ...please.Validate[int]) please.Validate[*Query] {
	return func(q *Query) error {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return err
		}
		if err = please.Join(o, opts...); err != nil {
			return err
		}
		q.offset = o
		return nil
	}
}

// Limit returns the limit value and the set indication.
func (q *Query) Limit() (int, bool) {
	return q.limit, q.limit > 0
}

// Limit sets the limit value with validation.
func Limit(limit int, opts ...please.Validate[int]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(limit, opts...); err != nil {
			return err
		}
		q.limit = limit
		return nil
	}
}

// LimitString sets the limit value with validation from a string.
func LimitString(limit string, opts ...please.Validate[int]) please.Validate[*Query] {
	return func(q *Query) error {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return err
		}
		if err = please.Join(l, opts...); err != nil {
			return err
		}
		q.limit = l
		return nil
	}
}

// Cursor returns the cursor value and the set indication.
func (q *Query) Cursor() (string, bool) {
	return q.cursor, q.cursor != ""
}

// Cursor sets the cursor value with validation.
func Cursor(cursor string, opts ...please.Validate[string]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(cursor, opts...); err != nil {
			return err
		}
		q.cursor = cursor
		return nil
	}
}

// Criteria returns the criteria for a field and the set indication.
func (q *Query) Criteria(field string) ([]Criteria[any], bool) {
	if q.criteria == nil {
		return nil, false
	}

	cs, ok := q.criteria[field]
	if len(cs) == 0 || !ok {
		return nil, false
	}

	return cs, ok
}

// Equal matches values that are equal to a specified value.
func Equal[T comparable](field string, value T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(value, opts...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Eq, value})
		return nil
	}
}

// NotEqual matches all values that are not equal to a specified value.
func NotEqual[T comparable](field string, value T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(value, opts...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Ne, value})
		return nil
	}
}

// Greater matches values that are greater than a specified value.
func Greater[T comparable](field string, value T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(value, opts...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Gt, value})
		return nil
	}
}

// GreaterOrEqual matches values that are greater than or equal to a specified value.
func GreaterOrEqual[T comparable](field string, value T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(value, opts...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Gte, value})
		return nil
	}
}

// Less matches values that are less than a specified value.
func Less[T comparable](field string, value T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(value, opts...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Lt, value})
		return nil
	}
}

// LessOrEqual matches values that are less than or equal to a specified value.
func LessOrEqual[T comparable](field string, value T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		if err := please.Join(value, opts...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Lte, value})
		return nil
	}
}

// OneOf matches any of the values specified in an array.
func OneOf[T comparable](field string, values []T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		errs := make([]error, 0, len(values))
		for _, value := range values {
			if err := please.Join(value, opts...); err != nil {
				errs = append(errs, err)
			}
		}
		if err := errors.Join(errs...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{In, values})
		return nil
	}
}

// NotOneOf matches none of the values specified in an array.
func NotOneOf[T comparable](field string, values []T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		errs := make([]error, 0, len(values))
		for _, value := range values {
			if err := please.Join(value, opts...); err != nil {
				errs = append(errs, err)
			}
		}
		if err := errors.Join(errs...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Nin, values})
		return nil
	}
}

// Between matches values that are between the specified values.
func Between[T cmp.Ordered](field string, x, y T, opts ...please.Validate[T]) please.Validate[*Query] {
	return func(q *Query) error {
		minimal := min(x, y)
		maximal := max(x, y)
		if err := please.Join(minimal, opts...); err != nil {
			return err
		}
		if err := please.Join(maximal, opts...); err != nil {
			return err
		}
		q.criteria[field] = append(q.criteria[field], Criteria[any]{Gte, minimal}, Criteria[any]{Lte, maximal})
		return nil
	}
}
