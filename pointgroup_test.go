package glot

import "testing"

func TestResetPointGroupStyle(t *testing.T) {
	dimensions := 2
	persist := true
	debug := true
	plot, _ := NewPlot(dimensions, persist, debug)

	plot.SetTitle("Test plot")
	plot.SetXLabel("X")
	plot.SetYLabel("Y")
	new_style := NewPlotObjectStyle(
		SetPointType(7),
		SetPointSize(1.5),
		SetLineColor("rgb", "red"),
	)
	plot.AddPointGroup("Sample1", "points", [][]int32{{1, 2, 3, 4}, {51, 8, 4, 11}}, *new_style)
	err := plot.ResetPointGroupStyle("Sam", "lines")
	if err == nil {
		t.Error("The specified pointgroup to be reset does not exist")
	}
}
