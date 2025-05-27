package problems

import (
	"fmt"
	"math/rand"
	"time"
)

// AdditionGenerator generates addition problems
type AdditionGenerator struct {
	maxDigits int
	random    *rand.Rand
}

// NewAdditionGenerator creates a new addition problem generator
func NewAdditionGenerator(maxDigits int) *AdditionGenerator {
	return &AdditionGenerator{
		maxDigits: maxDigits,
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Generate creates a new addition problem
func (g *AdditionGenerator) Generate() Problem {
	// Generate the first number with up to maxDigits
	maxNum1 := pow10(g.maxDigits) - 1
	num1 := g.random.Intn(maxNum1) + 1

	// Generate the second number with up to maxDigits
	maxNum2 := pow10(g.maxDigits) - 1
	num2 := g.random.Intn(maxNum2) + 1

	return Problem{
		Question: fmt.Sprintf("%d + %d", num1, num2),
		Answer:   num1 + num2,
		Type:     Addition,
	}
}

// pow10 returns 10^n
func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

// Type returns the type of problems this generator creates
func (g *AdditionGenerator) Type() ProblemType {
	return Addition
}

// Name returns a human-readable name for this problem type
func (g *AdditionGenerator) Name() string {
	return "Addition"
}
