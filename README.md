# Math Game

A simple terminal-based math game designed for 3rd graders to practice basic arithmetic operations.

## Features

- Four game variations: Addition, Subtraction, Multiplication, and Division
- 20 problems per game session
- Timed sessions to track progress
- History tracking of the last 10 game sessions per variation
- Simple terminal UI

## Requirements

- Go 1.17 or higher
- macOS (tested on macOS 12+)

## Installation

```bash
git clone https://github.com/huiruru/math-game.git
cd math-game
go build -o mathgame ./cmd/mathgame
```

## How to Build and Run

### Build

```bash
cd math-game

# Build using go build
go build -o mathgame ./cmd/mathgame

# Or use go install to install to your $GOPATH/bin
go install ./cmd/mathgame
```

### Run

```bash
# Run the executable directly
./mathgame

# Or if you used go install
mathgame
```

## How to Test

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./internal/problems

# Run tests with coverage
go test -cover ./...

# Run verbose tests
go test -v ./...
```

## Usage

Follow the on-screen instructions to select a game variation and play.

## Game Variations

- **Addition**: Problems with positive numbers up to 3 digits
- **Subtraction**: Problems with positive numbers up to 3 digits (results always positive)
- **Multiplication**: Problems from the multiplication table up to 12×12
- **Division**: Problems derived from the multiplication table up to 12×12 