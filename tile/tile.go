package tile

import "fmt"

type Tile struct {
	Head     int
	Tail     int
	Next     *Tile
	Previous *Tile
}

// SwitchHeadTail switches head and tail of the tile, same as rotating the tile
func (t *Tile) SwitchHeadTail() {
	if t.Next == nil && t.Previous == nil {
		t.Head, t.Tail = t.Tail, t.Head
	}
}

// Draw tile
func (t *Tile) Draw() string {
	return fmt.Sprintf("[%v|%v]", t.Head, t.Tail)
}
