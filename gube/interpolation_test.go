package gube

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLerp(t *testing.T) {
	testData := []struct {
		name string
		v0   float64
		v1   float64
		t    float64
		want float64
	}{
		{
			name: "First value",
			v0:   0,
			v1:   1,
			t:    0,
			want: 0,
		},
		{
			name: "Last value",
			v0:   0,
			v1:   7.5,
			t:    1,
			want: 7.5,
		},
		{
			name: "75% through",
			v0:   0,
			v1:   20,
			t:    .75,
			want: 15,
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			got := lerp(test.v0, test.v1, test.t)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("lerp() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTetrahedron(t *testing.T) {
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
			name: "Top value",
			r:    1.0,
			g:    1.0,
			b:    1.0,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
				tableSize: 2,
				tableData3D: &[][][]RGB{
					{[]RGB{{0, 0, 0}, {.1, .1, .1}}, []RGB{{0.2, 0.2, 0.2}, {.4, .4, .4}}},
					{[]RGB{{.6, .6, .6}, {.8, .8, .8}}, []RGB{{.5, .5, .5}, {1, 1, 1}}},
				},
			},
			want: RGB{1, 1, 1},
		},
		{
			name: "Bottom value",
			r:    0.0,
			g:    0.0,
			b:    0.0,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
				tableSize: 2,
				tableData3D: &[][][]RGB{
					{[]RGB{{0, 0, 0}, {.1, .1, .1}}, []RGB{{0.2, 0.2, 0.2}, {.4, .4, .4}}},
					{[]RGB{{.6, .6, .6}, {.8, .8, .8}}, []RGB{{.5, .5, .5}, {1, 1, 1}}},
				},
			},
			want: RGB{0, 0, 0},
		},
		{
			name: "Halfway through",
			r:    0.2,
			g:    0.5,
			b:    0.8,
			lut: &GubeImpl{
				domainMax: RGB{1.0, 1.0, 1.0},
				tableSize: 2,
				tableData3D: &[][][]RGB{
					{[]RGB{{0.1, 0.2, 0.4}, {.2, .4, .6}}, []RGB{{0.2, 0.3, 0.4}, {.4, .5, .6}}},
					{[]RGB{{.6, .5, .4}, {.8, .7, .6}}, []RGB{{.5, .6, .7}, {1, 1, 1}}},
				},
			},
			want: RGB{0.4, 0.51, 0.64},
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			got, gotErr := test.lut.lookUp3D(test.r, test.g, test.b)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("Test: %q :  Got error %v, wanted err=%v", test.name, gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("tetrahedron() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
