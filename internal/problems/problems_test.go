package problems

import (
	"fmt"
	"testing"
)

func TestMultiplicationGenerator(t *testing.T) {
	generator := NewMultiplicationGenerator(12)

	// Test type and name
	if generator.Type() != Multiplication {
		t.Errorf("Expected problem type %s, got %s", Multiplication, generator.Type())
	}

	if generator.Name() != "Multiplication" {
		t.Errorf("Expected name 'Multiplication', got '%s'", generator.Name())
	}

	// Generate and test 100 problems
	for i := 0; i < 100; i++ {
		problem := generator.Generate()

		// Verify problem type
		if problem.Type != Multiplication {
			t.Errorf("Expected problem type %s, got %s", Multiplication, problem.Type)
		}

		// Parse problem and verify answer
		var factor1, factor2 int
		n, err := fmt.Sscanf(problem.Question, "%d × %d", &factor1, &factor2)
		if err != nil || n != 2 {
			t.Errorf("Failed to parse problem: %s", problem.Question)
			continue
		}

		expectedAnswer := factor1 * factor2
		if problem.Answer != expectedAnswer {
			t.Errorf("Problem: %s, expected answer %d, got %d", problem.Question, expectedAnswer, problem.Answer)
		}

		// Verify factors are within range
		if factor1 < 1 || factor1 > 12 {
			t.Errorf("Factor1 out of range: %d", factor1)
		}
		if factor2 < 1 || factor2 > 12 {
			t.Errorf("Factor2 out of range: %d", factor2)
		}
	}
}

func TestDivisionGenerator(t *testing.T) {
	generator := NewDivisionGenerator(12)

	// Test type and name
	if generator.Type() != Division {
		t.Errorf("Expected problem type %s, got %s", Division, generator.Type())
	}

	if generator.Name() != "Division" {
		t.Errorf("Expected name 'Division', got '%s'", generator.Name())
	}

	// Generate and test 100 problems
	for i := 0; i < 100; i++ {
		problem := generator.Generate()

		// Verify problem type
		if problem.Type != Division {
			t.Errorf("Expected problem type %s, got %s", Division, problem.Type)
		}

		// Parse problem and verify answer
		var dividend, divisor int
		n, err := fmt.Sscanf(problem.Question, "%d ÷ %d", &dividend, &divisor)
		if err != nil || n != 2 {
			t.Errorf("Failed to parse problem: %s", problem.Question)
			continue
		}

		expectedAnswer := dividend / divisor
		if problem.Answer != expectedAnswer {
			t.Errorf("Problem: %s, expected answer %d, got %d", problem.Question, expectedAnswer, problem.Answer)
		}

		// Verify clean division
		if dividend%divisor != 0 {
			t.Errorf("Division with remainder: %d ÷ %d = %d r%d",
				dividend, divisor, expectedAnswer, dividend%divisor)
		}
	}
}

func TestAdditionGenerator(t *testing.T) {
	generator := NewAdditionGenerator(3)

	// Test type and name
	if generator.Type() != Addition {
		t.Errorf("Expected problem type %s, got %s", Addition, generator.Type())
	}

	if generator.Name() != "Addition" {
		t.Errorf("Expected name 'Addition', got '%s'", generator.Name())
	}

	// Generate and test 100 problems
	for i := 0; i < 100; i++ {
		problem := generator.Generate()

		// Verify problem type
		if problem.Type != Addition {
			t.Errorf("Expected problem type %s, got %s", Addition, problem.Type)
		}

		// Parse problem and verify answer
		var addend1, addend2 int
		n, err := fmt.Sscanf(problem.Question, "%d + %d", &addend1, &addend2)
		if err != nil || n != 2 {
			t.Errorf("Failed to parse problem: %s", problem.Question)
			continue
		}

		expectedAnswer := addend1 + addend2
		if problem.Answer != expectedAnswer {
			t.Errorf("Problem: %s, expected answer %d, got %d", problem.Question, expectedAnswer, problem.Answer)
		}

		// Verify number of digits
		if addend1 <= 0 || addend1 >= 1000 {
			t.Errorf("Addend1 out of range: %d", addend1)
		}
		if addend2 <= 0 || addend2 >= 1000 {
			t.Errorf("Addend2 out of range: %d", addend2)
		}
	}
}

func TestSubtractionGenerator(t *testing.T) {
	generator := NewSubtractionGenerator(3)

	// Test type and name
	if generator.Type() != Subtraction {
		t.Errorf("Expected problem type %s, got %s", Subtraction, generator.Type())
	}

	if generator.Name() != "Subtraction" {
		t.Errorf("Expected name 'Subtraction', got '%s'", generator.Name())
	}

	// Generate and test 100 problems
	for i := 0; i < 100; i++ {
		problem := generator.Generate()

		// Verify problem type
		if problem.Type != Subtraction {
			t.Errorf("Expected problem type %s, got %s", Subtraction, problem.Type)
		}

		// Parse problem and verify answer
		var minuend, subtrahend int
		n, err := fmt.Sscanf(problem.Question, "%d - %d", &minuend, &subtrahend)
		if err != nil || n != 2 {
			t.Errorf("Failed to parse problem: %s", problem.Question)
			continue
		}

		expectedAnswer := minuend - subtrahend
		if problem.Answer != expectedAnswer {
			t.Errorf("Problem: %s, expected answer %d, got %d", problem.Question, expectedAnswer, problem.Answer)
		}

		// Verify numbers and result are positive
		if minuend <= 0 || minuend >= 1000 {
			t.Errorf("Minuend out of range: %d", minuend)
		}
		if subtrahend <= 0 || subtrahend >= 1000 {
			t.Errorf("Subtrahend out of range: %d", subtrahend)
		}
		if expectedAnswer < 0 {
			t.Errorf("Negative answer: %d", expectedAnswer)
		}
	}
}
