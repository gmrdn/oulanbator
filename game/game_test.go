package game

import (
	"bytes"
	"strings"
	"testing"
)

func TestGameStartShouldGreetPlayers(t *testing.T) {
	input := strings.NewReader("")
	output := &bytes.Buffer{}

	g := NewGame(input, output)

	g.Start()

	expected := "Welcome to the game!\n"

	if output.String() != expected {
		t.Errorf("Expected %s, got %s", expected, output.String())
	}
}

func TestGameStartShouldInitializeGameState(t *testing.T) {
	input := strings.NewReader("")
	output := &bytes.Buffer{}

	g := NewGame(input, output)

	g.Start()

	expected := GameState{turn: 0, players: []Player{
		{name: "Alice", startingCube: Cube{-3, 3, 0}, objective: Cube{3, -3, 0}},
		{name: "Bob", startingCube: Cube{3, -3, 0}, objective: Cube{-3, 3, 0}},
	}, dangerToken: Cube{0, 0, 0}, playerTokens: [][]Cube{}}

	for i, player := range g.state.players {
		if player != expected.players[i] {
			t.Errorf("Expected %v, got %v", expected, g.state)
		}
	}
}
