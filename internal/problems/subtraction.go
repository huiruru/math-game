package problems

import (
	"fmt"
	"math/rand"
	"time"
)

// SubtractionGenerator generates subtraction problems
type SubtractionGenerator struct {
	maxDigits int
	random    *rand.Rand
}

// NewSubtractionGenerator creates a new subtraction problem generator
func NewSubtractionGenerator(maxDigits int) *SubtractionGenerator {
	return &SubtractionGenerator{
		maxDigits: maxDigits,
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Generate creates a new subtraction problem
func (g *SubtractionGenerator) Generate() Problem {
	// Generate the first number with up to maxDigits
	maxNum1 := pow10(g.maxDigits) - 1
	num1 := g.random.Intn(maxNum1) + 1

	// Generate the second number less than num1 to ensure positive result
	num2 := g.random.Intn(num1) + 1

	return Problem{
		Question: fmt.Sprintf("%d - %d", num1, num2),
		Answer:   num1 - num2,
		Type:     Subtraction,
	}
}

// Type returns the type of problems this generator creates
func (g *SubtractionGenerator) Type() ProblemType {
	return Subtraction
}

// Name returns a human-readable name for this problem type
func (g *SubtractionGenerator) Name() string {
	return "Subtraction"
}
