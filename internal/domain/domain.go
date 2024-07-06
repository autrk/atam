package domain

import (
	"fmt"

	"github.com/autrk/atam/internal/enum"
)

// Define the Property type
type Property enum.Member[string]

func (p Property) String() string {
	return fmt.Sprintf("%v", p.Value)
}

// Define Properties
var (
	builder = enum.NewBuilder[Property]()

	property_efficient = builder.Add(Property{"efficient"})
	property_flexible  = builder.Add(Property{"flexible"})
	property_operable  = builder.Add(Property{"operable"})
	property_reliable  = builder.Add(Property{"reliable"})
	property_safe      = builder.Add(Property{"safe"})
	property_secure    = builder.Add(Property{"secure"})
	property_suitable  = builder.Add(Property{"suitable"})
	property_usable    = builder.Add(Property{"usable"})

	Properties = builder.Enum()
)

type Quality struct {
	Name        string
	Description string
	Related     []string
	Properties  []Property
	Stakeholder []string
}

/*
func (o Operation) MarshalJSON() ([]byte, error) {
    return []byte(`"` + t.String() + `"`), nil
}

func (t *Operation) UnmarshalJSON(b []byte) error {
    *t = Parse(string(b))
    return nil
}
*/
