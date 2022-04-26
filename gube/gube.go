// Package gube implements functions to work with .cube LUT files.
package gube

import "errors"

const (
	// LUT_INVALID represents an invalid LUT type.
	LUT_INVALID = iota
	// LUT_1D is a LUT with 1D table data.
	LUT_1D
	// LUT_3D is a LUT with 3D table data.
	LUT_3D
)

// RGB represents a r,g,b value.
type RGB [3]float64

// Gube defines the methods used to work with .cube LUT files.
type Gube interface {
	// LookUp returns the transformed RGB value based on the LUT.
	LookUp(r float64, g float64, b float64) (RGB, error)
	// Name returns the name of this LUT.
	Name() string
	// TableType returns the table type (1D or 3D).
	TableType() int
	// TableSize returns the number of entries in this LUT.
	TableSize() int
	// TableData1D returns the table data for 1D LUTs.
	TableData1D() *[]RGB
	// TableData3D returns the table data for 3D LUTs.
	TableData3D() *[][][]RGB
	// Domain returns the min and max domain values for this LUT.
	Domain() (RGB, RGB)
}

var (
	// ErrOutsideOfDomain is the error returned when trying to look up values outside of the domain.
	ErrOutsideOfDomain = errors.New("value is outside of domain limits")
	// ErrInvalidLutType is the error returned when the LUT type is not a supported one.
	ErrInvalidLutType = errors.New("invalid LUT type")
)
