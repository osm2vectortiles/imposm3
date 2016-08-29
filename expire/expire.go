package expire

import (
	"fmt"
	"io"

	"github.com/omniscale/imposm3/element"
)

type Expireor interface {
	ExpireLinestring([]element.Node)
	ExpirePolygon([]element.Node)
	ExpirePoint(lon, lat float64)
}

func NewTileExpireor(maxZoom int) TileExpireor {
	return TileExpireor{
		tiles:   make(map[int]bool),
		maxZoom: maxZoom,
	}
}

type TileExpireor struct {
	// Space efficient tile store
	tiles map[int]bool
	// Max zoom level to evaluate
	maxZoom int
}

func (te *TileExpireor) ExpireLinestring(nodes []element.Node) {
	linestring := []Point{}
	for _, node := range nodes {
		linestring = append(linestring, Point{node.Long, node.Lat})
	}

	tiles, _ := CoverLinestring(linestring, te.maxZoom)
	for id, _ := range tiles {
		te.tiles[id] = true
	}
}

func (te *TileExpireor) ExpirePolygon(nodes []element.Node) {
	linearRing := []Point{}
	for _, node := range nodes {
		linearRing = append(linearRing, Point{node.Long, node.Lat})
	}

	tiles := CoverPolygon(linearRing, te.maxZoom)
	for id, _ := range tiles {
		te.tiles[id] = true
	}
}

func (te *TileExpireor) ExpirePoint(lon, lat float64) {
	tile := PointToTile(lon, lat, te.maxZoom)
	te.tiles[tile.toID()] = true
}

func (te *TileExpireor) WriteTiles(w io.Writer) {
	for id, _ := range te.tiles {
		tile := fromID(id)
		fmt.Fprintf(w, "%d/%d/%d\n", tile.X, tile.Y, tile.Z)
	}
}
