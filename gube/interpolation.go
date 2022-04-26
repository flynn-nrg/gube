package gube

import (
	"math"
)

func lerp(v0 float64, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

// Based on http://www.filmlight.ltd.uk/pdf/whitepapers/FL-TL-TN-0057-SoftwareLib.pdf
func (gi *GubeImpl) tetrahedron(r, g, b float64) RGB {
	var res RGB

	rIntF, rFrac := math.Modf(r)
	gIntF, gFrac := math.Modf(g)
	bIntF, bFrac := math.Modf(b)
	rNextF := gi.next(rIntF)
	gNextF := gi.next(gIntF)
	bNextF := gi.next(bIntF)
	rInt := int(rIntF)
	gInt := int(gIntF)
	bInt := int(bIntF)
	rNext := int(rNextF)
	gNext := int(gNextF)
	bNext := int(bNextF)

	v000 := (*gi.tableData3D)[rInt][gInt][bInt]
	v111 := (*gi.tableData3D)[rNext][gNext][bNext]

	if rFrac > gFrac {
		if gFrac > bFrac {
			v100 := (*gi.tableData3D)[rNext][gInt][bInt]
			v110 := (*gi.tableData3D)[rNext][gNext][bInt]
			res[0] = (1-rFrac)*v000[0] + (rFrac-gFrac)*v100[0] + (gFrac-bFrac)*v110[0] + bFrac*v111[0]
			res[1] = (1-rFrac)*v000[1] + (rFrac-gFrac)*v100[1] + (gFrac-bFrac)*v110[1] + bFrac*v111[1]
			res[2] = (1-rFrac)*v000[2] + (rFrac-gFrac)*v100[2] + (gFrac-bFrac)*v110[2] + bFrac*v111[2]
		} else if rFrac > bFrac {
			v100 := (*gi.tableData3D)[rNext][gInt][bInt]
			v101 := (*gi.tableData3D)[rNext][gInt][bNext]
			res[0] = (1-rFrac)*v000[0] + (rFrac-bFrac)*v100[0] + (bFrac-gFrac)*v101[0] + gFrac*v111[0]
			res[1] = (1-rFrac)*v000[1] + (rFrac-bFrac)*v100[1] + (bFrac-gFrac)*v101[1] + gFrac*v111[1]
			res[2] = (1-rFrac)*v000[2] + (rFrac-bFrac)*v100[2] + (bFrac-gFrac)*v101[2] + gFrac*v111[2]
		} else {
			v001 := (*gi.tableData3D)[rInt][gInt][bNext]
			v101 := (*gi.tableData3D)[rNext][gInt][bNext]
			res[0] = (1-bFrac)*v000[0] + (bFrac-rFrac)*v001[0] + (rFrac-gFrac)*v101[0] + gFrac*v111[0]
			res[1] = (1-bFrac)*v000[1] + (bFrac-rFrac)*v001[1] + (rFrac-gFrac)*v101[1] + gFrac*v111[1]
			res[1] = (1-bFrac)*v000[2] + (bFrac-rFrac)*v001[2] + (rFrac-gFrac)*v101[2] + gFrac*v111[2]
		}
	} else {
		if bFrac > gFrac {
			v001 := (*gi.tableData3D)[rInt][gInt][bNext]
			v011 := (*gi.tableData3D)[rInt][gNext][bNext]
			res[0] = (1-bFrac)*v000[0] + (bFrac-gFrac)*v001[0] + (gFrac-rFrac)*v011[0] + rFrac*v111[0]
			res[1] = (1-bFrac)*v000[1] + (bFrac-gFrac)*v001[1] + (gFrac-rFrac)*v011[1] + rFrac*v111[1]
			res[2] = (1-bFrac)*v000[2] + (bFrac-gFrac)*v001[2] + (gFrac-rFrac)*v011[2] + rFrac*v111[2]
		} else if bFrac > rFrac {
			v010 := (*gi.tableData3D)[rInt][gNext][bInt]
			v011 := (*gi.tableData3D)[rInt][gNext][bNext]
			res[0] = (1-gFrac)*v000[0] + (gFrac-bFrac)*v010[0] + (bFrac-rFrac)*v011[0] + rFrac*v111[0]
			res[1] = (1-gFrac)*v000[1] + (gFrac-bFrac)*v010[1] + (bFrac-rFrac)*v011[1] + rFrac*v111[1]
			res[2] = (1-gFrac)*v000[2] + (gFrac-bFrac)*v010[2] + (bFrac-rFrac)*v011[2] + rFrac*v111[2]
		} else {
			v010 := (*gi.tableData3D)[rInt][gNext][bInt]
			v110 := (*gi.tableData3D)[rNext][gNext][bInt]
			res[0] = (1-gFrac)*v000[0] + (gFrac-rFrac)*v010[0] + (rFrac-bFrac)*v110[0] + bFrac*v111[0]
			res[1] = (1-gFrac)*v000[0] + (gFrac-rFrac)*v010[0] + (rFrac-bFrac)*v110[0] + bFrac*v111[0]
			res[2] = (1-gFrac)*v000[0] + (gFrac-rFrac)*v010[0] + (rFrac-bFrac)*v110[0] + bFrac*v111[0]
		}
	}

	return res
}

func (gi *GubeImpl) next(x float64) int {
	return min(int(x+1), gi.tableSize-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
