package gube

import (
	"math"
)

func lerp(v0 float64, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

func (gi *GubeImpl) trilinear(r, g, b float64) RGB {
	rIntF := math.Trunc(r)
	gIntF := math.Trunc(g)
	bIntF := math.Trunc(b)
	rNextF := gi.next(rIntF)
	gNextF := gi.next(gIntF)
	bNextF := gi.next(bIntF)
	rInt := int(rIntF)
	gInt := int(gIntF)
	bInt := int(bIntF)
	rNext := int(rNextF)
	gNext := int(gNextF)
	bNext := int(bNextF)

	c000 := (*gi.tableData3D)[rInt][gInt][bInt]
	c001 := (*gi.tableData3D)[rInt][gInt][bNext]
	c010 := (*gi.tableData3D)[rInt][gNext][bInt]
	c011 := (*gi.tableData3D)[rInt][gNext][bNext]
	c100 := (*gi.tableData3D)[rNext][gInt][bInt]
	c101 := (*gi.tableData3D)[rNext][gInt][bNext]
	c110 := (*gi.tableData3D)[rNext][gNext][bInt]
	c111 := (*gi.tableData3D)[rNext][gNext][bNext]

	newR := trilinearSingleValue(r, g, b, rIntF, rNextF, gIntF, gNextF, bIntF, bNextF,
		c000[0], c001[0], c010[0], c011[0], c100[0], c101[0], c110[0], c111[0])
	newG := trilinearSingleValue(r, g, b, rIntF, rNextF, gIntF, gNextF, bIntF, bNextF,
		c000[1], c001[1], c010[1], c011[1], c100[1], c101[1], c110[1], c111[1])
	newB := trilinearSingleValue(r, g, b, rIntF, rNextF, gIntF, gNextF, bIntF, bNextF,
		c000[2], c001[2], c010[2], c011[2], c100[2], c101[2], c110[2], c111[2])

	return RGB{newR, newG, newB}
}

// https://en.wikipedia.org/wiki/Trilinear_interpolation
func trilinearSingleValue(x, y, z, x0, x1, y0, y1, z0, z1, c000, c001, c010, c011, c100, c101, c110, c111 float64) float64 {
	var xd, yd, zd float64

	if x1 == x0 {
		xd = x
	} else {
		xd = (x - x0) / (x1 - x0)
	}

	if y1 == y0 {
		yd = y
	} else {
		yd = (y - y0) / (y1 - y0)
	}

	if z1 == z0 {
		zd = z
	} else {
		zd = (z - z0) / (z1 - z0)
	}

	c00 := c000*(1-xd) + c100*xd
	c01 := c001*(1-xd) + c101*xd
	c10 := c010*(1-xd) + c110*xd
	c11 := c011*(1-xd) + c111*xd

	c0 := c00*(1-yd) + c10*yd
	c1 := c01*(1-yd) + c11*yd
	c := c0*(1-zd) + c1*zd

	return c
}

func (gi *GubeImpl) next(x float64) float64 {
	return float64(min(int(x+1), int(gi.tableSize)-1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
