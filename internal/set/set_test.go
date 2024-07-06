package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New(1, 2, 3)
	assert.Equal(t, 3, len(s))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}

func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3)
	assert.Equal(t, 3, len(s))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}

func TestContains(t *testing.T) {
	s := New(1, 2, 3)
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(4))
}

func TestMembers(t *testing.T) {
	s := New(1, 2, 3)
	members := s.Members()
	assert.ElementsMatch(t, []int{1, 2, 3}, members)
}

func TestUnion(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(3, 4, 5)
	unionSet := s1.Union(s2)
	expected := New(1, 2, 3, 4, 5)
	assert.ElementsMatch(t, expected.Members(), unionSet.Members())
}

func TestIntersection(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(3, 4, 5)
	intersectionSet := s1.Intersection(s2)
	expected := New(3)
	assert.ElementsMatch(t, expected.Members(), intersectionSet.Members())
}

func TestString(t *testing.T) {
	s := New(1, 2, 3)
	assert.Equal(t, "[1 2 3]", s.String())
}
