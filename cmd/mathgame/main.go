package main

import (
	"fmt"
	"os"
	"path/filepath"

	"math-game/internal/game"
	"math-game/internal/history"
	"math-game/internal/problems"
	"math-game/internal/ui"
)

const (
	totalProblems = 20
)

func main() {
	// Create UI
	userInterface := ui.NewTerminalUI()
	userInterface.Clear()

	// Welcome message
	fmt.Println("Welcome to Math Game!")
	fmt.Println("=====================")
	fmt.Println("Practice your math skills with fun challenges!")
	fmt.Println()

	// Create data directory
	dataDir := getDataDir()

	// Create history storage
	storage, err := history.NewFileStorage(dataDir)
	if err != nil {
		fmt.Printf("Error initializing storage: %v\n", err)
		os.Exit(1)
	}

	// Main game loop
	for {
		mainMenu(userInterface, storage)
	}
}

// getDataDir returns the path to the data directory
func getDataDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	dataDir := filepath.Join(homeDir, ".mathgame")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("Error creating data directory: %v\n", err)
		os.Exit(1)
	}

	return dataDir
}

// mainMenu displays the main menu and handles user selection
func mainMenu(userInterface ui.UI, storage history.Storage) {
	options := []string{
		"Play Addition",
		"Play Subtraction",
		"Play Multiplication",
		"Play Division",
		"View Addition History",
		"View Subtraction History",
		"View Multiplication History",
		"View Division History",
		"Exit",
	}

	choice, err := userInterface.ShowMenu(options)
	if err != nil {
		userInterface.ShowMessage(fmt.Sprintf("Error: %v", err))
		return
	}

	switch choice {
	case 0: // Addition
		playGame(userInterface, storage, problems.NewAdditionGenerator(2))
	case 1: // Subtraction
		playGame(userInterface, storage, problems.NewSubtractionGenerator(2))
	case 2: // Multiplication
		playGame(userInterface, storage, problems.NewMultiplicationGenerator(12))
	case 3: // Division
		playGame(userInterface, storage, problems.NewDivisionGenerator(12))
	case 4: // View Addition History
		showHistory(userInterface, storage, problems.Addition)
	case 5: // View Subtraction History
		showHistory(userInterface, storage, problems.Subtraction)
	case 6: // View Multiplication History
		showHistory(userInterface, storage, problems.Multiplication)
	case 7: // View Division History
		showHistory(userInterface, storage, problems.Division)
	case 8: // Exit
		fmt.Println("Thank you for playing Math Game!")
		os.Exit(0)
	}
}

// playGame runs a game session with the given problem generator
func playGame(userInterface ui.UI, storage history.Storage, generator problems.Generator) {
	userInterface.Clear()

	// Show game start message
	fmt.Printf("Starting %s Game\n", generator.Name())
	fmt.Printf("You will be given %d problems to solve.\n", totalProblems)
	fmt.Println("Press Enter to start...")
	fmt.Scanln()

	// Create and start a new game session
	session := game.NewSession(generator, totalProblems)
	session.Start()

	// Present each problem
	for i := 0; i < totalProblems; i++ {
		problem := generator.Generate()
		userAnswer, err := userInterface.DisplayProblem(problem, i+1, totalProblems)

		if err != nil {
			userInterface.ShowMessage(fmt.Sprintf("Error: %v", err))
			i-- // Retry the same problem
			continue
		}

		// Check answer and record result
		correct := userAnswer == problem.Answer
		session.AddResult(correct)

		// Show feedback
		if correct {
			userInterface.ShowMessage("Correct!")
		} else {
			userInterface.ShowMessage(fmt.Sprintf("Incorrect. The correct answer is %d.", problem.Answer))
		}
	}

	// End the session and get results
	session.End()
	result := session.GetResult()

	// Save result to history
	if err := storage.SaveResult(result); err != nil {
		userInterface.ShowMessage(fmt.Sprintf("Failed to save result: %v", err))
	}

	// Show results
	userInterface.ShowResults(result)
}

// showHistory displays the history for a specific problem type
func showHistory(userInterface ui.UI, storage history.Storage, problemType problems.ProblemType) {
	// Get the last 10 results for this problem type
	results, err := storage.GetResults(problemType, 10)
	if err != nil {
		userInterface.ShowMessage(fmt.Sprintf("Error retrieving history: %v", err))
		return
	}

	// Show history
	userInterface.ShowHistory(results)
}
