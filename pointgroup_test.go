package glot

import "testing"

func TestResetPointGroupStyle(t *testing.T) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := NewPlot(dimensions, persist, debug)

	new_style := NewPlotObjectStyle(
		SetPointType(6),
		SetPointSize(1.5),
		SetLineColor("rgb", "red"),
	)
	plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11}, *new_style)
	err := plot.ResetPointGroupStyle("Sam", "lines")
	if err == nil {
		t.Error("The specified pointgroup to be reset does not exist")
	}
}
