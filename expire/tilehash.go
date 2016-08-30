package expire

import "sort"

type TileHash map[int]struct{}

func (th TileHash) AddTileFraction(x, y float64, z int) {
	th[toID(int(x), int(y), z)] = struct{}{}
}
func (th TileHash) AddTile(x, y int, z int) {
	th[toID(x, y, z)] = struct{}{}
}

func FromTiles(tiles []Tile) TileHash {
	th := TileHash{}
	for _, t := range tiles {
		th.AddTile(t.X, t.Y, t.Z)
	}
	return th
}

func (th TileHash) ToTiles() []Tile {
	tiles := []Tile{}
	for id, _ := range th {
		tiles = append(tiles, fromID(id))
	}
	sort.Sort(ByID(tiles))
	return tiles
}

func fromID(id int) Tile {
	z := id % 32
	dim := 2 * (1 << uint(z))
	xy := ((id - z) / 32)
	x := xy % dim
	y := ((xy - x) / dim) % dim
	return Tile{x, y, z}
}

func (t Tile) toID() int {
	return toID(t.X, t.Y, t.Z)
}

func toID(x, y int, z int) int {
	dim := 2 * (1 << uint(z))
	return ((dim*y + x) * 32) + z
}
