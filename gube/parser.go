package gube

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func parseFromReader(r io.Reader) (*GubeImpl, error) {
	var tableData1D []RGB
	var tableData3D [][][]RGB
	var i, j, k int64

	g := &GubeImpl{
		domainMin: RGB{},
		domainMax: RGB{1.0, 1.0, 1.0},
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "#") {
			continue
		}
		if strings.Contains(s, "TITLE") {
			title := strings.Split(s, "\"")
			g.name = strings.Trim(title[1], "\"")
			continue
		}
		if strings.Contains(s, "DOMAIN_MIN") {
			domainMin := strings.Split(s, " ")
			rl, err := strconv.ParseFloat(domainMin[1], 32)
			if err != nil {
				return nil, err
			}
			gl, err := strconv.ParseFloat(domainMin[2], 32)
			if err != nil {
				return nil, err
			}
			bl, err := strconv.ParseFloat(domainMin[3], 32)
			if err != nil {
				return nil, err
			}
			g.domainMin = RGB{rl, gl, bl}
			continue
		}
		if strings.Contains(s, "DOMAIN_MAX") {
			domainMax := strings.Split(s, " ")
			rh, err := strconv.ParseFloat(domainMax[1], 32)
			if err != nil {
				return nil, err
			}
			gh, err := strconv.ParseFloat(domainMax[2], 32)
			if err != nil {
				return nil, err
			}
			bh, err := strconv.ParseFloat(domainMax[3], 32)
			if err != nil {
				return nil, err
			}
			g.domainMax = RGB{rh, gh, bh}
			continue
		}
		if strings.Contains(s, "LUT_1D_SIZE") {
			g.tableType = LUT_1D
			lut1dSize := strings.Split(s, " ")
			tableSize, err := strconv.ParseInt(lut1dSize[1], 10, 32)
			if err != nil {
				return nil, err
			}
			g.tableSize = tableSize
			continue
		}
		if strings.Contains(s, "LUT_3D_SIZE") {
			g.tableType = LUT_3D
			lut1dSize := strings.Split(s, " ")
			tableSize, err := strconv.ParseInt(lut1dSize[1], 10, 32)
			if err != nil {
				return nil, err
			}
			g.tableSize = tableSize
			tableData3D = make([][][]RGB, tableSize)
			for i := 0; i < int(tableSize); i++ {
				tableData3D[i] = make([][]RGB, tableSize)
				for j := 0; j < int(tableSize); j++ {
					tableData3D[i][j] = make([]RGB, tableSize)
				}
			}
			continue
		}

		// Try to parse rgb values. At this point we must know the table type and size.
		rgb, err := parseRGB(s)
		if err != nil {
			return nil, err
		}
		switch g.tableType {
		case LUT_1D:
			tableData1D = append(tableData1D, rgb)
		case LUT_3D:
			tableData3D[i][j][k] = rgb
			i++
			if i > g.tableSize-1 {
				i = 0
				j++
				if j > g.tableSize-1 {
					j = 0
					k++
				}
			}
		default:
			return nil, ErrInvalidLutData
		}
	}

	switch g.tableType {
	case LUT_1D:
		g.tableData1D = &tableData1D
	case LUT_3D:
		g.tableData3D = &tableData3D
	}

	return g, nil
}

func parseRGB(s string) (RGB, error) {
	var res RGB

	values := strings.Split(s, " ")
	r, err := strconv.ParseFloat(values[0], 32)
	if err != nil {
		return res, err
	}
	g, err := strconv.ParseFloat(values[1], 32)
	if err != nil {
		return res, err
	}
	b, err := strconv.ParseFloat(values[2], 32)
	if err != nil {
		return res, err
	}

	res[0] = r
	res[1] = g
	res[2] = b

	return res, nil
}
