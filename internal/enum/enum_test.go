package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestMember Member[string]

func (m TestMember) String() string {
	return m.Value
}

func TestEnum(t *testing.T) {
	m1 := TestMember{"A"}
	m2 := TestMember{"B"}
	m3 := TestMember{"C"}
	e := New(m1, m2, m3)

	assert.True(t, e.Contains(m1))
	assert.False(t, e.Contains(TestMember{"D"}))

	member, err := e.Parse("A")
	assert.NoError(t, err)
	assert.Equal(t, m1, member)

	_, err = e.Parse("D")
	assert.Error(t, err)

	assert.ElementsMatch(t, []TestMember{m1, m2, m3}, e.Members())
	assert.ElementsMatch(t, []string{"A", "B", "C"}, e.Values())

	expectedString := "A, B, C"
	assert.Equal(t, expectedString, e.String())
}

func TestBuilder(t *testing.T) {
	b := NewBuilder[TestMember, string]()
	m1 := b.Add(TestMember{"A"})
	m2 := b.Add(TestMember{"B"})
	m3 := b.Add(TestMember{"C"})
	e := b.Enum()

	assert.True(t, e.Contains(m1))
	assert.True(t, e.Contains(m2))
	assert.True(t, e.Contains(m3))
}
