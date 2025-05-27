package game

import (
	"time"

	"math-game/internal/problems"
)

// Result represents the outcome of a game session
type Result struct {
	ProblemType    problems.ProblemType
	CorrectCount   int
	TotalCount     int
	Duration       time.Duration
	CompletionTime time.Time
}

// PercentCorrect returns the percentage of correct answers
func (r Result) PercentCorrect() float64 {
	if r.TotalCount == 0 {
		return 0
	}
	return float64(r.CorrectCount) / float64(r.TotalCount) * 100
}

// Session represents a single game session
type Session struct {
	ProblemType   problems.ProblemType
	Generator     problems.Generator
	TotalProblems int
	StartTime     time.Time
	EndTime       time.Time
	Answers       []bool // true for correct, false for incorrect
}

// NewSession creates a new game session with the given problem generator
func NewSession(generator problems.Generator, totalProblems int) *Session {
	return &Session{
		ProblemType:   generator.Type(),
		Generator:     generator,
		TotalProblems: totalProblems,
		Answers:       make([]bool, 0, totalProblems),
	}
}

// Start begins a new game session
func (s *Session) Start() {
	s.StartTime = time.Now()
}

// End completes a game session
func (s *Session) End() {
	s.EndTime = time.Now()
}

// Duration returns the total duration of the session
func (s *Session) Duration() time.Duration {
	if s.EndTime.IsZero() {
		return time.Since(s.StartTime)
	}
	return s.EndTime.Sub(s.StartTime)
}

// AddResult records the result of a single problem (correct or incorrect)
func (s *Session) AddResult(correct bool) {
	s.Answers = append(s.Answers, correct)
}

// CorrectCount returns the number of correct answers
func (s *Session) CorrectCount() int {
	count := 0
	for _, correct := range s.Answers {
		if correct {
			count++
		}
	}
	return count
}

// GetResult returns the final result of the session
func (s *Session) GetResult() Result {
	return Result{
		ProblemType:    s.ProblemType,
		CorrectCount:   s.CorrectCount(),
		TotalCount:     len(s.Answers),
		Duration:       s.Duration(),
		CompletionTime: s.EndTime,
	}
}
