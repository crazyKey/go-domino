package game

import (
	"errors"
	"math/rand"

	"github.com/crazyKey/go-domino/board"
	"github.com/crazyKey/go-domino/player"
	"github.com/crazyKey/go-domino/tile"
)

type Game struct {
	Players []*player.Player
	Tiles   []*tile.Tile
	Board   *board.Board
}

// Add players to the game, max 4 players
func (g *Game) AddPlayer(name string) error {
	if len(g.Players) == 4 {
		return errors.New("max four players")
	}

	p := player.Player{
		Name: name,
	}

	g.Players = append(g.Players, &p)

	return nil
}

// Pick 7 starting tiles
func (g *Game) PickStartingTiles() []*tile.Tile {
	picked := g.Tiles[:7]
	g.Tiles = g.Tiles[7:]

	return picked
}

// Generate the starting 28 tiles
func (g *Game) GenerateTiles() {
	g.Board = &board.Board{}

	for i := 0; i <= 6; i++ {
		for j := 0; j <= i; j++ {
			t := tile.Tile{
				Head: i,
				Tail: j,
			}

			g.Tiles = append(g.Tiles, &t)
		}
	}

	rand.Shuffle(len(g.Tiles), func(i, j int) {
		g.Tiles[i], g.Tiles[j] = g.Tiles[j], g.Tiles[i]
	})
}

// Find the starting player position
func (g *Game) FindStartingPlayer() int {
	startingPlayer := 0
	var biggerDouble int

	for i, p := range g.Players {
		playerMaxDouble := p.GetBiggerDouble()
		if playerMaxDouble > biggerDouble {
			biggerDouble = playerMaxDouble
			startingPlayer = i
		}
	}

	return startingPlayer
}

// PlaceStartingTile places the starting tile to board
func (g *Game) PlaceStartingTile(t *tile.Tile) {
	g.Board.StartingTile = t
	g.Board.RightTile = t
	g.Board.LeftTile = t
}

// PlaceTile places tile to board
func (g *Game) PlaceTile(t *tile.Tile, left bool) {
	if left {
		// [new] [left]
		t.Next = g.Board.LeftTile
		g.Board.LeftTile.Previous = t
		g.Board.LeftTile = t
	} else {
		// [right] [new]
		t.Previous = g.Board.RightTile
		g.Board.RightTile.Next = t
		g.Board.RightTile = t
	}
}

// PickTile picks a tile from the free tiles
func (g *Game) PickTile() *tile.Tile {
	var picked *tile.Tile

	if len(g.Tiles) > 0 {
		picked = g.Tiles[0]
		g.Tiles = g.Tiles[1:]
	}

	return picked
}
