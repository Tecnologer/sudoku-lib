package sudoku

//ValidationErrors is map with the errors
type ValidationErrors struct {
	Errs  map[errorType][]*ErrorCoordinate `json:"errors"`
	Count int                              `json:"count"`
}

type errorType string

//ErrorCoordinate contains the values for X,Y where the error is
type ErrorCoordinate struct {
	*Coordinate
}

const (
	emptyError    errorType = "empty"
	invalidColumn errorType = "column"
	invalidRow    errorType = "row"
	invalidSquare errorType = "square"
)

func (e *ValidationErrors) appendError(t errorType, err *ErrorCoordinate) {
	if e.Errs == nil {
		e.Errs = make(map[errorType][]*ErrorCoordinate)
	}

	if _, ok := e.Errs[t]; !ok {
		e.Errs[t] = make([]*ErrorCoordinate, 0)
	}
	e.Count++
	e.Errs[t] = append(e.Errs[t], err)
}

//NewErrorCoordinate creates an instance of coordinate with error
func NewErrorCoordinate(x, y int) *ErrorCoordinate {
	return &ErrorCoordinate{
		NewCoordinate(x, y),
	}
}
