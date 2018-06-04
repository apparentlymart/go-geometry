package svgpath

import (
	"testing"

	"github.com/apparentlymart/go-geometry/geom"
	"github.com/go-test/deep"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Src      string
		WantPath Path
		WantErr  string
	}{
		{
			``,
			nil,
			``,
		},
		{
			`Z`,
			Path{
				Close,
			},
			``,
		},
		{
			`z`,
			Path{
				CloseRel,
			},
			``,
		},
		{
			`M 10,20`,
			Path{
				Move(geom.Point{10, 20}),
			},
			``,
		},
		{
			`M 10 20`,
			Path{
				Move(geom.Point{10, 20}),
			},
			``,
		},
		{
			`M -10 -20`,
			Path{
				Move(geom.Point{-10, -20}),
			},
			``,
		},
		{
			`M 10-20`,
			Path{
				Move(geom.Point{10, -20}),
			},
			``,
		},
		{
			`M 10.2.1`,
			Path{
				Move(geom.Point{10.2, 0.1}),
			},
			``,
		},
		{
			`M 10,20 30,40`,
			Path{
				Move(geom.Point{10, 20}),
				Line(geom.Point{30, 40}),
			},
			``,
		},
		{
			`M 10 20`,
			Path{
				Move(geom.Point{10, 20}),
			},
			``,
		},
		{
			`M 10 20 30 40`,
			Path{
				Move(geom.Point{10, 20}),
				Line(geom.Point{30, 40}),
			},
			``,
		},
		{
			`m 10 20 30 40`,
			Path{
				MoveRel(geom.Point{10, 20}),
				LineRel(geom.Point{30, 40}),
			},
			``,
		},
		{
			`L 10 20`,
			Path{
				Line(geom.Point{10, 20}),
			},
			``,
		},
		{
			`L 10 20 30 40`,
			Path{
				Line(geom.Point{10, 20}),
				Line(geom.Point{30, 40}),
			},
			``,
		},
		{
			`H 10`,
			Path{
				HorizLine(10),
			},
			``,
		},
		{
			`H 10 20`,
			Path{
				HorizLine(10),
				HorizLine(20),
			},
			``,
		},
		{
			`h 10 20`,
			Path{
				HorizLineRel(10),
				HorizLineRel(20),
			},
			``,
		},
		{
			`H 10,20`,
			Path{
				HorizLine(10),
				HorizLine(20),
			},
			``,
		},
		{
			`V 10`,
			Path{
				VertLine(10),
			},
			``,
		},
		{
			`V 10 20`,
			Path{
				VertLine(10),
				VertLine(20),
			},
			``,
		},
		{
			`v 10 20`,
			Path{
				VertLineRel(10),
				VertLineRel(20),
			},
			``,
		},
		{
			`V 10,20`,
			Path{
				VertLine(10),
				VertLine(20),
			},
			``,
		},
		{
			`C 10,20 30,40 50,60`,
			Path{
				Curve(
					geom.Point{10, 20},
					geom.Point{30, 40},
					geom.Point{50, 60},
				),
			},
			``,
		},
		{
			`C 10,20 30,40 50,60 1,2 3,4 5,6`,
			Path{
				Curve(
					geom.Point{10, 20},
					geom.Point{30, 40},
					geom.Point{50, 60},
				),
				Curve(
					geom.Point{1, 2},
					geom.Point{3, 4},
					geom.Point{5, 6},
				),
			},
			``,
		},
		{
			`c 10,20 30,40 50,60 1,2 3,4 5,6`,
			Path{
				CurveRel(
					geom.Point{10, 20},
					geom.Point{30, 40},
					geom.Point{50, 60},
				),
				CurveRel(
					geom.Point{1, 2},
					geom.Point{3, 4},
					geom.Point{5, 6},
				),
			},
			``,
		},
		{
			`C 10,20 30,40`,
			Path{},
			`expecting number but got end of string`, // not enough args
		},
		{
			`S 10,20 30,40`,
			Path{
				SmoothCurve(
					geom.Point{10, 20},
					geom.Point{30, 40},
				),
			},
			``,
		},
		{
			`S 10,20 30,40 50,60 70,80`,
			Path{
				SmoothCurve(
					geom.Point{10, 20},
					geom.Point{30, 40},
				),
				SmoothCurve(
					geom.Point{50, 60},
					geom.Point{70, 80},
				),
			},
			``,
		},
		{
			`s 10,20 30,40 50,60 70,80`,
			Path{
				SmoothCurveRel(
					geom.Point{10, 20},
					geom.Point{30, 40},
				),
				SmoothCurveRel(
					geom.Point{50, 60},
					geom.Point{70, 80},
				),
			},
			``,
		},
		{
			`Q 10,20 30,40`,
			Path{
				QuadCurve(
					geom.Point{10, 20},
					geom.Point{30, 40},
				),
			},
			``,
		},
		{
			`Q 10,20 30,40 50,60 70,80`,
			Path{
				QuadCurve(
					geom.Point{10, 20},
					geom.Point{30, 40},
				),
				QuadCurve(
					geom.Point{50, 60},
					geom.Point{70, 80},
				),
			},
			``,
		},
		{
			`q 10,20 30,40 50,60 70,80`,
			Path{
				QuadCurveRel(
					geom.Point{10, 20},
					geom.Point{30, 40},
				),
				QuadCurveRel(
					geom.Point{50, 60},
					geom.Point{70, 80},
				),
			},
			``,
		},
		{
			`T 10,20`,
			Path{
				SmoothQuadCurve(geom.Point{10, 20}),
			},
			``,
		},
		{
			`T 10,20 30,40 50,60`,
			Path{
				SmoothQuadCurve(geom.Point{10, 20}),
				SmoothQuadCurve(geom.Point{30, 40}),
				SmoothQuadCurve(geom.Point{50, 60}),
			},
			``,
		},
		{
			`t 10,20 30,40 50,60`,
			Path{
				SmoothQuadCurveRel(geom.Point{10, 20}),
				SmoothQuadCurveRel(geom.Point{30, 40}),
				SmoothQuadCurveRel(geom.Point{50, 60}),
			},
			``,
		},
		{
			`A 10,20 45 1 0 30,40`,
			Path{
				Arc(
					geom.Point{10, 20}, 45,
					true, false,
					geom.Point{30, 40},
				),
			},
			``,
		},
		{
			`A 10,20 45 1 0 30,40 50,60 30 0 1 70,80`,
			Path{
				Arc(
					geom.Point{10, 20}, 45,
					true, false,
					geom.Point{30, 40},
				),
				Arc(
					geom.Point{50, 60}, 30,
					false, true,
					geom.Point{70, 80},
				),
			},
			``,
		},
		{
			`a 10,20 45 1 0 30,40 50,60 30 0 1 70,80`,
			Path{
				ArcRel(
					geom.Point{10, 20}, 45,
					true, false,
					geom.Point{30, 40},
				),
				ArcRel(
					geom.Point{50, 60}, 30,
					false, true,
					geom.Point{70, 80},
				),
			},
			``,
		},

		// Now some more-interesting compound examples
		{
			`M600,350 l 50,-25 
             a25,25 -30 0,1 50,-25 l 50,-25 
             a25,50 -30 0,1 50,-25 l 50,-25 
             a25,75 -30 0,1 50,-25 l 50,-25 
             a25,100 -30 0,1 50,-25 l 50,-25`,
			Path{
				Move(geom.Point{600, 350}),
				LineRel(geom.Point{50, -25}),
				ArcRel(
					geom.Point{25, 25}, -30,
					false, true,
					geom.Point{50, -25},
				),
				LineRel(geom.Point{50, -25}),
				ArcRel(
					geom.Point{25, 50}, -30,
					false, true,
					geom.Point{50, -25},
				),
				LineRel(geom.Point{50, -25}),
				ArcRel(
					geom.Point{25, 75}, -30,
					false, true,
					geom.Point{50, -25},
				),
				LineRel(geom.Point{50, -25}),
				ArcRel(
					geom.Point{25, 100}, -30,
					false, true,
					geom.Point{50, -25},
				),
				LineRel(geom.Point{50, -25}),
			},
			``,
		},
		{
			`M100,200 C100,100 250,100 250,200 S400,300 400,200`,
			Path{
				Move(geom.Point{100, 200}),
				Curve(
					geom.Point{100, 100},
					geom.Point{250, 100},
					geom.Point{250, 200},
				),
				SmoothCurve(
					geom.Point{400, 300},
					geom.Point{400, 200},
				),
			},
			``,
		},

		// Invalid paths
		{
			`M`,
			Path{},
			`expecting number but got end of string`,
		},
		{
			`M 10`,
			Path{},
			`expecting number but got end of string`,
		},
		{
			`X`,
			nil,
			`expecting instruction but got invalid token`,
		},
		{
			`M 10, 10 Z 100 M 20, 20`,
			Path{
				Move(geom.Point{10, 10}),
				Close,
			},
			`expecting instruction but got number`,
		},
	}

	for _, test := range tests {
		t.Run(test.Src, func(t *testing.T) {
			gotPath, gotErr := Parse(test.Src)

			for _, problem := range deep.Equal(gotPath, test.WantPath) {
				t.Errorf(problem)
			}

			if test.WantErr != "" {
				if gotErr == nil {
					t.Errorf("succeeded; want error: %s", test.WantErr)
				} else {
					gotErrStr := gotErr.Error()
					if gotErrStr != test.WantErr {
						t.Errorf("wrong error\ngot:  %s\nwant: %s", gotErrStr, test.WantErr)
					}
				}
			} else {
				if gotErr != nil {
					t.Errorf("unexpected error: %s", gotErr.Error())
				}
			}
		})
	}
}
