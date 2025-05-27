package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"math-game/internal/game"
	"math-game/internal/problems"
)

// UI represents the user interface for the game
type UI interface {
	// ShowMenu displays the main menu and returns the selected option
	ShowMenu(options []string) (int, error)

	// DisplayProblem shows a problem to the user and gets their answer
	DisplayProblem(problem problems.Problem, problemNum, total int) (int, error)

	// ShowResults displays the results of a completed game session
	ShowResults(result game.Result)

	// ShowHistory displays historical game results
	ShowHistory(results []game.Result)

	// ShowMessage displays a message to the user
	ShowMessage(message string)

	// Clear clears the screen
	Clear()
}

// TerminalUI implements a simple terminal-based UI
type TerminalUI struct {
	reader *bufio.Reader
}

// NewTerminalUI creates a new terminal UI
func NewTerminalUI() *TerminalUI {
	return &TerminalUI{
		reader: bufio.NewReader(os.Stdin),
	}
}

// readInput reads a line of input from the user
func (ui *TerminalUI) readInput() (string, error) {
	input, err := ui.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// Clear clears the terminal screen
func (ui *TerminalUI) Clear() {
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen
}

// ShowMessage displays a message to the user
func (ui *TerminalUI) ShowMessage(message string) {
	fmt.Println(message)
}

// ShowMenu displays the main menu and returns the selected option
func (ui *TerminalUI) ShowMenu(options []string) (int, error) {
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	fmt.Print("\nEnter your choice (1-", len(options), "): ")

	input, err := ui.readInput()
	if err != nil {
		return 0, err
	}

	var choice int
	if _, err := fmt.Sscanf(input, "%d", &choice); err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	if choice < 1 || choice > len(options) {
		return 0, fmt.Errorf("choice out of range")
	}

	return choice - 1, nil
}

// DisplayProblem shows a problem to the user and gets their answer
func (ui *TerminalUI) DisplayProblem(problem problems.Problem, problemNum, total int) (int, error) {
	fmt.Printf("\nProblem %d of %d:\n", problemNum, total)
	fmt.Printf("%s = ? ", problem.Question)

	input, err := ui.readInput()
	if err != nil {
		return 0, err
	}

	var answer int
	if _, err := fmt.Sscanf(input, "%d", &answer); err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	return answer, nil
}

// ShowResults displays the results of a completed game session
func (ui *TerminalUI) ShowResults(result game.Result) {
	ui.Clear()
	fmt.Println("Game Results:")
	fmt.Println("-------------")
	fmt.Printf("Game Type: %s\n", result.ProblemType)
	fmt.Printf("Score: %d / %d (%.1f%%)\n",
		result.CorrectCount,
		result.TotalCount,
		result.PercentCorrect())
	fmt.Printf("Time: %s\n", formatDuration(result.Duration))
	fmt.Println("\nPress Enter to continue...")
	ui.readInput()
}

// formatDuration formats a duration as MM:SS
func formatDuration(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := seconds / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

// ShowHistory displays historical game results
func (ui *TerminalUI) ShowHistory(results []game.Result) {
	ui.Clear()

	if len(results) == 0 {
		fmt.Println("No history available.")
		fmt.Println("\nPress Enter to continue...")
		ui.readInput()
		return
	}

	fmt.Printf("History for %s:\n", results[0].ProblemType)
	fmt.Println("-------------------")

	for i, result := range results {
		fmt.Printf("%d. Score: %d/%d (%.1f%%) - Time: %s - %s\n",
			i+1,
			result.CorrectCount,
			result.TotalCount,
			result.PercentCorrect(),
			formatDuration(result.Duration),
			result.CompletionTime.Format("Jan 02, 2006 15:04"))
	}

	fmt.Println("\nPress Enter to continue...")
	ui.readInput()
}
