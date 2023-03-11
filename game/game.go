package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Game struct {
	input  io.Reader
	output io.Writer
	state  GameState
}

type GameState struct {
	turn         int
	players      []Player
	dangerToken  Cube
	playerTokens [][]Cube
}

type Player struct {
	name         string
	startingCube Cube
	objective    Cube
}

func NewGame(input io.Reader, output io.Writer) *Game {
	g := Game{
		input,
		output,
		GameState{dangerToken: Cube{0, 0, 0}, playerTokens: [][]Cube{}},
	}
	return &g
}

type Cube struct {
	q, r, s int
}

func (g *Game) Start() {
	g.output.Write([]byte("Welcome to the game!\n"))
	g.state.players = append(g.state.players, Player{
		name: "Alice",
		startingCube: Cube{
			q: -3,
			r: 3,
			s: 0,
		},
		objective: Cube{
			q: 3,
			r: -3,
			s: 0,
		},
	})
	g.state.players = append(g.state.players, Player{
		name: "Bob",
		startingCube: Cube{
			q: 3,
			r: -3,
			s: 0,
		},
		objective: Cube{
			q: -3,
			r: 3,
			s: 0,
		},
	})
	g.state.dangerToken = Cube{0, 0, 0}
	g.state.turn = 0

	fmt.Println(g.state)
}

func (g *Game) Play() {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')

		g.output.Write([]byte(message))
	}
}
