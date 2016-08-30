package expire

import "github.com/omniscale/imposm3/element"

type Expireor interface {
	// Expire everything covered by the linestring
	ExpireLinestring([]element.Node)
	// Expire everything covered and enclosed by the linear ring
	// TODO: Only supports the outer ring of polygon. No holes.
	ExpirePolygon([]element.Node)
	// Expire a single point
	ExpirePoint(lon, lat float64)
}
