package set

import "fmt"

// Set is a collection of unique elements of a specified type E.
//
// This generic implementation uses a map to ensure all elements are unique.
// The zero value of a Set is an empty set ready to use.
//
// Example usage:
//     s := set.New(1, 2, 3)
//     s.Add(4)
//     if s.Contains(2) {
//         fmt.Println("Set contains 2")
//     }
//
// E must be a comparable type, which means it supports the == and != operators.
type Set[E comparable] map[E]struct{}
type void struct{}

// New creates a new Set containing the provided members.
func New[E comparable](member ...E) Set[E] {
    set := Set[E]{}
    for _, m := range member {
        set[m] = void{}
    }
    return set
}

// Add inserts the provided members into the Set.
func (s Set[E]) Add(member ...E) {
    for _, m := range member {
        s[m] = void{}
    }
}

// Contains checks if the Set contains the specified member.
func (s Set[E]) Contains(member E) bool {
    _, ok := s[member]
    return ok
}

// Members returns a slice containing all members of the Set.
func (s Set[E]) Members() []E {
    result := make([]E, 0, len(s))
    for member := range s {
        result = append(result, member)
    }
    return result
}

// Union returns a new Set containing all members from both Sets.
func (s Set[E]) Union(s2 Set[E]) Set[E] {
    result := New(s.Members()...)
    result.Add(s2.Members()...)
    return result
}

// Intersection returns a new Set containing only the members present in both Sets.
func (s Set[E]) Intersection(s2 Set[E]) Set[E] {
    result := New[E]()
    for _, member := range s.Members() {
        if s2.Contains(member) {
            result.Add(member)
        }
    }
    return result
}

// String returns a string representation of the Set.
func (s Set[E]) String() string {
    return fmt.Sprintf("%v", s.Members())
}