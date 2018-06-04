package svgpath

import (
	"fmt"
	"testing"

	"github.com/apparentlymart/go-geometry/geom"

	"github.com/go-test/deep"
)

func TestPathSubpaths(t *testing.T) {
	tests := []struct {
		Path Path
		Want []Path
	}{
		{
			nil,
			nil,
		},
		{
			Path{},
			nil,
		},
		{
			Path{
				Move(geom.Point{10, 10}),
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
				},
			},
		},
		{
			Path{
				Move(geom.Point{10, 10}),
				Line(geom.Point{100, 200}),
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
					Line(geom.Point{100, 200}),
				},
			},
		},
		{
			Path{
				Move(geom.Point{10, 10}),
				Line(geom.Point{100, 200}),
				Close,
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
					Line(geom.Point{100, 200}),
					Close,
				},
			},
		},
		{
			Path{
				Move(geom.Point{10, 10}),
				Line(geom.Point{100, 200}),
				Move(geom.Point{20, 30}),
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
					Line(geom.Point{100, 200}),
				},
				Path{
					Move(geom.Point{20, 30}),
				},
			},
		},
		{
			Path{
				Move(geom.Point{10, 10}),
				Line(geom.Point{100, 200}),
				MoveRel(geom.Point{20, 30}),
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
					Line(geom.Point{100, 200}),
				},
				Path{
					MoveRel(geom.Point{20, 30}),
				},
			},
		},
		{
			Path{
				Move(geom.Point{10, 10}),
				Line(geom.Point{100, 200}),
				Move(geom.Point{20, 30}),
				Line(geom.Point{101, 201}),
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
					Line(geom.Point{100, 200}),
				},
				Path{
					Move(geom.Point{20, 30}),
					Line(geom.Point{101, 201}),
				},
			},
		},
		{
			Path{
				Move(geom.Point{10, 10}),
				Line(geom.Point{100, 200}),
				Close,
				Line(geom.Point{101, 201}),
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
					Line(geom.Point{100, 200}),
					Close,
				},
				Path{
					Line(geom.Point{101, 201}),
				},
			},
		},
		{
			Path{
				Move(geom.Point{10, 10}),
				Line(geom.Point{100, 200}),
				Close,
				Move(geom.Point{20, 30}),
				Line(geom.Point{101, 201}),
			},
			[]Path{
				Path{
					Move(geom.Point{10, 10}),
					Line(geom.Point{100, 200}),
					Close,
				},
				Path{
					Move(geom.Point{20, 30}),
					Line(geom.Point{101, 201}),
				},
			},
		},
		{
			Path{ // Degenerate case
				Close,
				Close,
				Close,
			},
			[]Path{
				Path{
					Close,
				},
				Path{
					Close,
				},
				Path{
					Close,
				},
			},
		},
		{
			Path{ // Degenerate case
				Close,
				CloseRel,
				Close,
			},
			[]Path{
				Path{
					Close,
				},
				Path{
					CloseRel,
				},
				Path{
					Close,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%#v", test.Path), func(t *testing.T) {
			got := test.Path.Subpaths()

			for _, problem := range deep.Equal(got, test.Want) {
				t.Error(problem)
			}
		})
	}
}
