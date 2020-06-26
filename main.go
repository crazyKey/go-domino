package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crazyKey/go-domino/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("+++ The Domino Game +++")

	g := game.Game{}

	// Add players
	for {
		var newPlayer string
		fmt.Print("Add player? (y/n): ")
		fmt.Scan(&newPlayer)

		if newPlayer == "n" {
			if len(g.Players) >= 2 {
				break
			}
			fmt.Println("You need at least two players")
		}

		if newPlayer == "y" {
			var playerName string
			fmt.Print("Player name: ")
			fmt.Scan(&playerName)

			err := g.AddPlayer(playerName)

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// Give tiles to players
	g.GenerateTiles()
	for _, p := range g.Players {
		picked := g.PickStartingTiles()
		p.Tiles = picked
	}

	// Find starting player
	startingPlayer := g.FindStartingPlayer()
	var winner bool

	// Game
	for i := startingPlayer; ; i++ {
		p := g.Players[i]

		fmt.Println("Playing: ", p.Name)

		g.Board.Draw()
		fmt.Println("")

		if g.Board.StartingTile == nil {
			tileToPlace := p.DrawTilesList(-1, -1)
			g.PlaceStartingTile(tileToPlace)
		} else {
			lineLeftEnd, lineRightEnd := g.Board.GetLineEnds()
			matchingTiles := p.GetMatchingTiles(lineLeftEnd, lineRightEnd)

			// Game ends
			if len(matchingTiles) == 0 && len(g.Tiles) == 0 {
				break
			}

			// Pick tile
			if len(matchingTiles) == 0 {
				pickedTile := g.PickTile()
				p.Tiles = append(p.Tiles, pickedTile)
				fmt.Println("You picked up the tile ", pickedTile.Draw())

				if i == len(g.Players)-1 {
					i = 0
				}
				continue
			}

			// Place tile
			for {
				tileToPlace := p.DrawTilesList(lineLeftEnd, lineRightEnd)

				if tileToPlace.Head == lineRightEnd {
					g.PlaceTile(tileToPlace, false)
					break
				} else if tileToPlace.Tail == lineLeftEnd {
					g.PlaceTile(tileToPlace, true)
					break
				}

				tileToPlace.SwitchHeadTail()

				if tileToPlace.Head == lineRightEnd {
					g.PlaceTile(tileToPlace, false)
					break
				} else if tileToPlace.Tail == lineLeftEnd {
					g.PlaceTile(tileToPlace, true)
					break
				}

				panic("this should not happen")
			}

			if len(p.Tiles) == 0 {
				fmt.Printf("Well done %s, you won the game", p.Name)
				winner = true
				break
			}
		}

		if i == len(g.Players)-1 {
			i = 0
		}
	}

	// Find the winner by total dots
	if !winner {
		fmt.Println("good news, everyone is a winner!")
	}
}
