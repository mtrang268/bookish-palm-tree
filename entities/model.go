package entities

import "fmt"

// Level defines the level of coverage a plan provides
type Level string

const (
	INVALID = Level("INVALID")
	Bronze = Level("Bronze")
	Silver = Level("Silver")
	Gold = Level("Gold")
	Platinum = Level("Platinum")
	Catastrophic = Level("Catastrophic")
)

// StateCode defines the abbreviation for a state, e.g Colorado = CO
type StateCode string

type RateNumber uint16

// RateArea defines the geographic region
type RateArea struct {
	State StateCode
	Number RateNumber
}

type ZipCode string

type Rate float32

type Rates []Rate
func (r Rates) Len() int {
	return len(r)
}
func (r Rates) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r Rates) Less(i, j int) bool {
	return r[i] < r[j]
}

// ParseLevel parses the level string
func ParseLevel(level string) (Level, error) {
	switch level {
	case string(Bronze):
		return Bronze, nil
	case string(Silver):
		return Silver, nil
	case string(Gold):
		return Gold, nil
	case string(Platinum):
		return Platinum, nil
	case string(Catastrophic):
		return Catastrophic, nil
	}

	return INVALID, fmt.Errorf("unsupported level %s", level)
}
