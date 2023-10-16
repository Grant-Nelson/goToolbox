package collections

import "goToolbox/utils"

// Collection is a collection of values.
type Collection[T any] interface {
	Enumerable[T]
	Countable
	utils.Stringer
	utils.Equatable

	// Empty indicates if the collection is empty.
	Empty() bool
}
