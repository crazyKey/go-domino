package board

import (
	"fmt"
	"github.com/crazyKey/go-domino/tile"
)

type Board struct {
	StartingTile *tile.Tile
	LeftTile     *tile.Tile
	RightTile    *tile.Tile
}

// GetLineEnds returns the ends of the line
func (b *Board) GetLineEnds() (int, int) {
	return b.LeftTile.Head, b.RightTile.Tail
}

// Draw list of tiles
func (b *Board) Draw() {
	fmt.Print("Line: ")

	if b.StartingTile == nil {
		return
	}

	t := b.LeftTile

	for t != nil {
		fmt.Print(t.Draw())
		t = t.Next
	}
}
