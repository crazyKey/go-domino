package player

import (
	"fmt"
	"github.com/crazyKey/go-domino/tile"
)

type Player struct {
	Name  string
	Tiles []*tile.Tile
}

// GetBiggerDouble returns the value of the bigger double tile
func (p *Player) GetBiggerDouble() int {
	m := 0

	for _, t := range p.Tiles {
		if t.Head == t.Tail {
			m = max(m, t.Head)
		}
	}

	return m
}

// GetMatchingTiles returns the list of tiles matching with two ends
func (p *Player) GetMatchingTiles(n1 int, n2 int) []*tile.Tile {
	var matching []*tile.Tile

	for _, t := range p.Tiles {
		if t.Head == n1 || t.Head == n2 || t.Tail == n1 || t.Tail == n2 {
			matching = append(matching, t)
		}
	}

	return matching
}

// TotalDots returns the count of the dots of the player tiles
func (p *Player) TotalDots() int {
	tot := 0

	for _, t := range p.Tiles {
		tot += t.Head + t.Tail
	}

	return tot
}

// DrawTilesList draw list of tiles available
func (p *Player) DrawTilesList(lineLeftEnd int, lineRightEnd int) *tile.Tile {
	fmt.Print("Tiles: ")

	for i, t := range p.Tiles {
		fmt.Printf("%v: %s ", i, t.Draw())
	}

	for {
		var tileToPLace int
		fmt.Print("Place tile: ")
		fmt.Scan(&tileToPLace)

		if tileToPLace >= len(p.Tiles) {
			continue
		}

		t := p.Tiles[tileToPLace]

		// No left or right end to match
		if lineLeftEnd == -1 && lineRightEnd == -1 {
			p.Tiles = append(p.Tiles[:tileToPLace], p.Tiles[tileToPLace+1:]...)
			return t
		}

		// Must match left or right end
		if t.Head == lineRightEnd || t.Head == lineLeftEnd || t.Tail == lineRightEnd || t.Tail == lineLeftEnd {
			p.Tiles = append(p.Tiles[:tileToPLace], p.Tiles[tileToPLace+1:]...)
			return t
		}
	}
}

// Return max between two integers
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
