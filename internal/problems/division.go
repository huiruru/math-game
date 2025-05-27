package problems

import (
	"fmt"
	"math/rand"
	"time"
)

// DivisionGenerator generates division problems
type DivisionGenerator struct {
	maxFactor int
	random    *rand.Rand
}

// NewDivisionGenerator creates a new division problem generator
func NewDivisionGenerator(maxFactor int) *DivisionGenerator {
	return &DivisionGenerator{
		maxFactor: maxFactor,
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Generate creates a new division problem
func (g *DivisionGenerator) Generate() Problem {
	// For division, we'll generate a multiplication problem first,
	// then convert it to a division problem to ensure clean division
	factor1 := g.random.Intn(g.maxFactor) + 1
	factor2 := g.random.Intn(g.maxFactor) + 1

	product := factor1 * factor2

	// Create a division problem using the product and one of the factors
	return Problem{
		Question: fmt.Sprintf("%d ÷ %d", product, factor1),
		Answer:   factor2,
		Type:     Division,
	}
}

// Type returns the type of problems this generator creates
func (g *DivisionGenerator) Type() ProblemType {
	return Division
}

// Name returns a human-readable name for this problem type
func (g *DivisionGenerator) Name() string {
	return "Division"
}
