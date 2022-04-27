package gube

import (
	"io"
	"math"
)

// Ensure interface compliance.
var _ Gube = (*GubeImpl)(nil)

type GubeImpl struct {
	name        string
	tableType   int64
	tableSize   int64
	tableData1D *[]RGB
	tableData3D *[][][]RGB
	domainMin   RGB
	domainMax   RGB
}

// NewFromReader returns a Gube instance created from the reader data.
func NewFromReader(r io.Reader) (*GubeImpl, error) {
	return parseFromReader(r)
}

func (gi *GubeImpl) LookUp(r float64, g float64, b float64) (RGB, error) {
	switch gi.tableType {
	case LUT_1D:
		return gi.lookUp1D(r, g, b)
	case LUT_3D:
		return gi.lookUp3D(r, g, b)
	default:
		return RGB{}, ErrInvalidLutType
	}
}

func (gi *GubeImpl) Name() string {
	return gi.name
}

func (gi *GubeImpl) TableType() int64 {
	return gi.tableType
}

func (gi *GubeImpl) TableSize() int64 {
	return gi.tableSize
}

func (gi *GubeImpl) TableData1D() *[]RGB {
	return gi.tableData1D
}

func (gi *GubeImpl) TableData3D() *[][][]RGB {
	return gi.tableData3D
}

func (gi *GubeImpl) Domain() (RGB, RGB) {
	return gi.domainMin, gi.domainMax
}

func (gi *GubeImpl) lookUp3D(r, g, b float64) (RGB, error) {
	var res RGB

	if !gi.withinDomain(r, g, b) {
		return res, ErrOutsideOfDomain
	}

	return gi.trilinear(r*float64(gi.tableSize-1), g*float64(gi.tableSize-1), b*float64(gi.tableSize-1)), nil
}

func (gi *GubeImpl) lookUp1D(r, g, b float64) (RGB, error) {
	var res RGB

	if !gi.withinDomain(r, g, b) {
		return res, ErrOutsideOfDomain
	}

	res[0] = gi.lookUp1DSingleValue(r, 0)
	res[1] = gi.lookUp1DSingleValue(g, 1)
	res[2] = gi.lookUp1DSingleValue(b, 2)

	return res, nil
}

// For 1D LUTs we perform a linear interpolation if necessary.
func (gi *GubeImpl) lookUp1DSingleValue(v float64, index int) float64 {
	vInt, t := math.Modf(v)
	v1Index := int(vInt * float64(gi.tableSize-1))
	if t == 0 {
		return (*gi.tableData1D)[v1Index][index]
	} else {
		v2Index := v1Index + 1
		v0 := (*gi.tableData1D)[v1Index][index]
		v1 := (*gi.tableData1D)[v2Index][index]
		return lerp(v0, v1, t)
	}
}

func (gi *GubeImpl) withinDomain(r, g, b float64) bool {
	return r >= gi.domainMin[0] && g >= gi.domainMin[1] && b >= gi.domainMin[2] &&
		r <= gi.domainMax[0] && g <= gi.domainMax[1] && b <= gi.domainMax[2]
}
