package gube

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	testData := []struct {
		name    string
		cube    string
		want    *GubeImpl
		wantErr bool
	}{
		{
			name: "Simple 1D LUT",
			cube: `TITLE "Test"
LUT_1D_SIZE 3
DOMAIN_MIN 0 0 0
DOMAIN_MAX 2 4 6
0 0 0
0.5 1 1.5
1 1 1`,
			want: &GubeImpl{
				name:        "Test",
				tableType:   LUT_1D,
				tableSize:   3,
				tableData1D: &[]RGB{{0, 0, 0}, {0.5, 1, 1.5}, {1, 1, 1}},
				domainMax:   RGB{2, 4, 6},
			},
		},
		{
			name: "Simple 3D LUT",
			cube: `TITLE "3D LUT"
LUT_3D_SIZE 2
0 0 0
1 0 0
0 .75 0
1 .75 0
0 .25 1
1 .25 1
0 1 1
1 1 1`,
			want: &GubeImpl{
				name:        "3D LUT",
				tableType:   LUT_3D,
				tableSize:   2,
				tableData3D: &[][][]RGB{{{{0, 0, 0}, {0, 0.25, 1}}, {{0, 0.75, 0}, {0, 1, 1}}}, {{{1, 0, 0}, {1, 0.25, 1}}, {{1, 0.75, 0}, {1, 1, 1}}}},
				domainMax:   RGB{1, 1, 1},
			},
		},
		{
			name:    "Invalid data",
			cube:    "DEFEKT Techniker ist informiert",
			wantErr: true,
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			got, gotErr := parseFromReader(strings.NewReader(test.cube))
			if (gotErr != nil) != test.wantErr {
				t.Errorf("Test: %q :  Got error %v, wanted err=%v", test.name, gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.want, got, cmp.AllowUnexported(GubeImpl{})); diff != "" {
				t.Errorf("parseFromReader() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
