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
