package expire

import (
	"bytes"
	"testing"
)

func TestWriteTiles(t *testing.T) {
	expected := "8627/5753/14\n"
	point := Point{1065162.58495039, 5965498.83778885}

	expireor := NewTileExpireor(14)
	expireor.ExpirePoint(point.lon, point.lat)

	buf := new(bytes.Buffer)
	expireor.WriteTiles(buf)

	if buf.String() != expected {
		t.Error("Unexpected tiles were written", buf.String())
	}
}
