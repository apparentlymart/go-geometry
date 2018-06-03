package svgpath

import (
	"fmt"
	"testing"

	"github.com/apparentlymart/go-geometry/geom"
)

func TestCommandToAbsolute(t *testing.T) {
	start := geom.Point{10, 20}
	prev := geom.Point{30, 40}

	tests := []struct {
		Recv      Command
		WantCmd   Command
		WantPoint geom.Point
	}{
		{
			Move(geom.Point{2, 7}),
			Move(geom.Point{2, 7}),
			geom.Point{2, 7},
		},
		{
			MoveRel(geom.Point{2, 7}),
			Move(geom.Point{12, 27}),
			geom.Point{12, 27},
		},
		{
			Line(geom.Point{2, 7}),
			Line(geom.Point{2, 7}),
			geom.Point{2, 7},
		},
		{
			LineRel(geom.Point{2, 7}),
			Line(geom.Point{12, 27}),
			geom.Point{12, 27},
		},
		{
			Close,
			Close,
			start,
		},
		{
			CloseRel,
			Close,
			start,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%#v", test.Recv), func(t *testing.T) {
			gotCmd, gotPoint := test.Recv.ToAbsolute(start, prev)

			if !gotCmd.Equal(test.WantCmd) {
				t.Errorf("wrong new command\ngot:  %#v\nwant: %#v", gotCmd, test.WantCmd)
			}
			if gotPoint != test.WantPoint {
				t.Errorf("wrong new point\ngot:  %#v\nwant: %#v", gotPoint, test.WantPoint)
			}
		})
	}
}

func TestInstruction(t *testing.T) {

	if got, want := MoveTo.Absolute(), true; got != want {
		t.Errorf("wrong MoveTo.Absolute() -> %t; want %t", got, want)
	}
	if got, want := MoveToRel.Absolute(), false; got != want {
		t.Errorf("wrong MoveToRel.Absolute() -> %t; want %t", got, want)
	}
	if got, want := MoveTo.Relative(), false; got != want {
		t.Errorf("wrong MoveTo.Relative() -> %t; want %t", got, want)
	}
	if got, want := MoveToRel.Relative(), true; got != want {
		t.Errorf("wrong MoveToRel.Relative() -> %t; want %t", got, want)
	}
	if got, want := MoveTo.ToAbsolute(), MoveTo; got != want {
		t.Errorf("wrong MoveTo.ToAbsolute() -> %s; want %s", got, want)
	}
	if got, want := MoveTo.ToRelative(), MoveToRel; got != want {
		t.Errorf("wrong MoveTo.ToRelative() -> %s; want %s", got, want)
	}
	if got, want := MoveToRel.ToAbsolute(), MoveTo; got != want {
		t.Errorf("wrong MoveToRel.ToAbsolute() -> %s; want %s", got, want)
	}
	if got, want := MoveToRel.ToRelative(), MoveToRel; got != want {
		t.Errorf("wrong MoveToRel.ToRelative() -> %s; want %s", got, want)
	}

}
