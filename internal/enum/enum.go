package enum

import (
	"errors"
	"fmt"
	"strings"

	"github.com/autrk/atam/internal/set"
)

// Member represents an enum member with a specific value.
type Member[T comparable] struct {
	Value T
}

// ValueHolder is an interface that constrains the types used in Enum.
// It requires the type to have a Value field and implement the Stringer interface.
type ValueHolder[T comparable] interface {
	~struct{ Value T }
	fmt.Stringer
}

// Enum is a collection of unique enum members.
// It provides methods to manage and query these members.
//
// After creating an Enum using New, no new members can be added.
type Enum[V ValueHolder[T], T comparable] struct {
	members set.Set[V]
}

// New creates a new Enum instance from the provided members.
//
// After creating an Enum using New, no new members can be added.
func New[V ValueHolder[T], T comparable](members ...V) Enum[V, T] {
	return Enum[V, T]{members: set.New(members...)}
}

// Contains checks if the enum contains the specified member.
func (e Enum[V, T]) Contains(member V) bool {
	for _, m := range e.members.Members() {
		if m == member {
			return true
		}
	}
	return false
}

// Parse converts a raw value into an enum member.
// If the value is not a member of the enum, it returns an error.
func (e Enum[V, T]) Parse(value T) (V, error) {
	member := V{value}
	if e.Contains(member) {
		return member, nil
	}
	return member, errors.New("the value is not an enum member")
}

// Members returns a slice of all members of the enum.
func (e Enum[V, T]) Members() []V {
	return e.members.Members()
}

// Values returns a slice of values of all members of the enum.
func (e Enum[V, T]) Values() []T {
	res := make([]T, 0, len(e.members))
	for _, m := range e.members.Members() {
		member := Member[T](m)
		res = append(res, member.Value)
	}
	return res
}

// String returns a comma-separated string of the values of all enum members.
func (e Enum[V, T]) String() string {
	values := make([]string, 0, len(e.members))
	for _, m := range e.Members() {
		values = append(values, m.String())
	}
	return strings.Join(values, ", ")
}

// Builder is a helper for constructing Enum instances.
// It allows adding members incrementally before creating the Enum.
type Builder[V ValueHolder[T], T comparable] struct {
	members  []V
	finished bool
}

// NewBuilder creates a new Builder instance.
func NewBuilder[V ValueHolder[T], T comparable]() Builder[V, T] {
	return Builder[V, T]{make([]V, 0), false}
}

// Add registers a new member with the Builder.
func (b *Builder[V, T]) Add(member V) V {
	b.members = append(b.members, member)
	return member
}

// Enum creates a new Enum with all members added to the Builder.
func (b *Builder[V, T]) Enum() Enum[V, T] {
	b.finished = true
	return New(b.members...)
}
