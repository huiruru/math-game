package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"math-game/internal/game"
	"math-game/internal/problems"
)

// Storage interface for saving and loading game history
type Storage interface {
	SaveResult(result game.Result) error
	GetResults(problemType problems.ProblemType, limit int) ([]game.Result, error)
}

// FileStorage implements history storage using files
type FileStorage struct {
	BaseDir string
}

// NewFileStorage creates a new file-based storage for game history
func NewFileStorage(baseDir string) (*FileStorage, error) {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		if err := os.MkdirAll(baseDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create history directory: %w", err)
		}
	}

	return &FileStorage{
		BaseDir: baseDir,
	}, nil
}

// getFilePath returns the path to the history file for a specific problem type
func (s *FileStorage) getFilePath(problemType problems.ProblemType) string {
	return filepath.Join(s.BaseDir, fmt.Sprintf("%s.json", problemType))
}

// SaveResult saves a game result to storage
func (s *FileStorage) SaveResult(result game.Result) error {
	filePath := s.getFilePath(result.ProblemType)

	// Load existing results
	results, err := s.GetResults(result.ProblemType, 0)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load existing results: %w", err)
	}

	// Add new result
	results = append(results, result)

	// Sort by completion time (newest first)
	sort.Slice(results, func(i, j int) bool {
		return results[i].CompletionTime.After(results[j].CompletionTime)
	})

	// Keep only the most recent results (max 10)
	if len(results) > 10 {
		results = results[:10]
	}

	// Save to file
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal results: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write results file: %w", err)
	}

	return nil
}

// GetResults loads game results from storage
func (s *FileStorage) GetResults(problemType problems.ProblemType, limit int) ([]game.Result, error) {
	filePath := s.getFilePath(problemType)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []game.Result{}, nil
	}

	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read results file: %w", err)
	}

	// Unmarshal JSON
	var results []game.Result
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, fmt.Errorf("failed to unmarshal results: %w", err)
	}

	// Sort by completion time (newest first)
	sort.Slice(results, func(i, j int) bool {
		return results[i].CompletionTime.After(results[j].CompletionTime)
	})

	// Limit results if requested
	if limit > 0 && len(results) > limit {
		results = results[:limit]
	}

	return results, nil
}
