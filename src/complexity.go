package sudoku

import (
	"bytes"
	"encoding/json"
	"strings"
)

//ComplexityLevel enum of complex game
type ComplexityLevel byte

var (
	complexityLevelToString = [...]string{"Invalid", "Easy", "Basic", "Medium", "Hard", "Master", "Empty", "Test"}
	complexityLevelToID     = map[string]ComplexityLevel{
		"Invalid": InvalidLevel,
		"Easy":    EasyLevel,
		"Basic":   BasicLevel,
		"Medium":  MediumLevel,
		"Hard":    HardLevel,
		"Master":  MasterLevel,
		"Empty":   EmptyLevel,
		"Test":    TestLevel,
	}
)

//GetComplexities returns the list of difficulties
func GetComplexities() [8]string {
	return complexityLevelToString
}

func (l ComplexityLevel) String() string {
	return complexityLevelToString[l]
}

type complexity [9]*complexData

type complexData struct {
	randCount int
	xMax      int
	yMax      int
}

type square struct {
	xMin int
	yMin int
	xMax int
	yMax int
}

const (
	InvalidLevel ComplexityLevel = iota
	EasyLevel
	BasicLevel
	MediumLevel
	HardLevel
	MasterLevel
	EmptyLevel
	TestLevel
)

//StringToComplexity parses string to ComplexityLevel, Default: EasyLevel
func StringToComplexity(c string) ComplexityLevel {
	for i, v := range complexityLevelToString {
		if strings.ToLower(v) == strings.ToLower(c) {
			return ComplexityLevel(i)
		}
	}

	return InvalidLevel
}

func buildComplexity(level ComplexityLevel) *complexity {
	var randCount [9]int
	switch level {
	case TestLevel:
		randCount = [9]int{1, 4, 1, 1, 1, 4, 1, 1, 1}
	case BasicLevel:
		randCount = [9]int{4, 5, 4, 5, 5, 5, 5, 5, 5}
	case MediumLevel:
		randCount = [9]int{2, 3, 4, 4, 1, 3, 2, 1, 3}
	case HardLevel:
		randCount = [9]int{3, 2, 3, 4, 2, 3, 3, 2, 3}
	case MasterLevel:
		randCount = [9]int{1, 4, 3, 3, 2, 3, 3, 3, 3}
	case EasyLevel:
		fallthrough
	default:
		randCount = [9]int{4, 5, 3, 5, 5, 5, 5, 5, 5}
	}

	var c complexity = [9]*complexData{}

	for i, n := range randCount {
		xMax, yMax := squareIndexLimit(i)

		c[i] = &complexData{
			randCount: n,
			xMax:      xMax,
			yMax:      yMax,
		}
	}

	return &c
}

func squareIndexLimit(i int) (int, int) {
	return ((i / 3) * 3) + 2, ((i % 3) * 3) + 2
}

// MarshalJSON marshals the enum as a quoted json string
func (l ComplexityLevel) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(complexityLevelToString[l])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (l *ComplexityLevel) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*l = complexityLevelToID[j]
	return nil
}
