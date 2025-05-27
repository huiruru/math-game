package problems

import (
	"fmt"
	"math/rand"
	"time"
)

// MultiplicationGenerator generates multiplication problems
type MultiplicationGenerator struct {
	maxFactor int
	random    *rand.Rand
}

// NewMultiplicationGenerator creates a new multiplication problem generator
func NewMultiplicationGenerator(maxFactor int) *MultiplicationGenerator {
	return &MultiplicationGenerator{
		maxFactor: maxFactor,
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Generate creates a new multiplication problem
func (g *MultiplicationGenerator) Generate() Problem {
	// Generate two random factors from 1 to maxFactor
	factor1 := g.random.Intn(g.maxFactor) + 1
	factor2 := g.random.Intn(g.maxFactor) + 1

	return Problem{
		Question: fmt.Sprintf("%d × %d", factor1, factor2),
		Answer:   factor1 * factor2,
		Type:     Multiplication,
	}
}

// Type returns the type of problems this generator creates
func (g *MultiplicationGenerator) Type() ProblemType {
	return Multiplication
}

// Name returns a human-readable name for this problem type
func (g *MultiplicationGenerator) Name() string {
	return "Multiplication"
}
