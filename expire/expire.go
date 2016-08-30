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

type NullExpireor struct{}

func (_ NullExpireor) ExpireLinestring([]element.Node) {}
func (_ NullExpireor) ExpirePolygon([]element.Node)    {}
func (_ NullExpireor) ExpirePoint(lon, lat float64)    {}
