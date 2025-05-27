package problems

// ProblemType represents the type of math problem
type ProblemType string

const (
	Addition       ProblemType = "addition"
	Subtraction    ProblemType = "subtraction"
	Multiplication ProblemType = "multiplication"
	Division       ProblemType = "division"
)

// Problem represents a single math problem
type Problem struct {
	Question string
	Answer   int
	Type     ProblemType
}

// String returns a string representation of the problem
func (p Problem) String() string {
	return p.Question
}

// Generator defines the interface for problem generators
type Generator interface {
	// Generate creates a new math problem
	Generate() Problem

	// Type returns the type of problems this generator creates
	Type() ProblemType

	// Name returns a human-readable name for this problem type
	Name() string
}
