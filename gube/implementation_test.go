package gube

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWithinDomain(t *testing.T) {
	testData := []struct {
		name string
		r    float64
		g    float64
		b    float64
		lut  *GubeImpl
		want bool
	}{
		{
			name: "All values ok",
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
			},
			want: true,
		},
		{
			name: "R is outside of range (too high)",
			r:    8.0,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
			},
		},
		{
			name: "G is outside of range (too high)",
			g:    3.0,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
			},
		},
		{
			name: "B is outside of range (too high)",
			g:    100.0,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
			},
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			got := test.lut.withinDomain(test.r, test.g, test.b)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("withinDomain() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLookUp1D(t *testing.T) {
	testData := []struct {
		name    string
		r       float64
		g       float64
		b       float64
		lut     *GubeImpl
		want    RGB
		wantErr bool
	}{
		{
			name: "Fetch some values from the table",
			r:    1.0,
			g:    0.5,
			b:    0,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
				tableSize: 10,
				tableData1D: &[]RGB{
					{0, 0, 1},
					{0.1, 0.05, .9},
					{0.2, 0.1, .8},
					{0.3, 0.15, .7},
					{0.4, 0.2, .6},
					{0.5, 0.25, .5},
					{0.6, 0.3, .4},
					{0.7, 0.35, .3},
					{0.9, 0.4, .2},
					{1, 0.5, .1},
				},
			},
			want: RGB{1, 0.025, 1},
		},
		{
			name: "Interpolate",
			r:    0.5,
			g:    0.5,
			b:    0.5,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
				tableSize: 2,
				tableData1D: &[]RGB{
					{0, 0, 0},
					{1, 0.5, .1},
				},
			},
			want: RGB{0.5, 0.25, 0.05},
		},
		{
			name: "Outside of domain values",
			r:    0.5,
			g:    4.5,
			b:    0.5,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
				tableSize: 2,
				tableData1D: &[]RGB{
					{0, 0, 0},
					{1, 0.5, .1},
				},
			},
			wantErr: true,
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			got, gotErr := test.lut.lookUp1D(test.r, test.g, test.b)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("Test: %q :  Got error %v, wanted err=%v", test.name, gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("lookUp1D() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
