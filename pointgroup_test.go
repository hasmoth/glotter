package glot

import "testing"

func TestResetPointGroupStyle(t *testing.T) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := NewPlot(dimensions, persist, debug)
	style := PlotObjectStyle{
		PointType: &PlotObjectType{"pt", 7},
		PointSize: &PlotObjectSize{"ps", 0.7},
	}

	new_style := NewPlotObjectStyle(
		SetPointType(6),
		SetPointSize(1.5),
		SetLineColor("rgb", "grey"),
	)
	plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11}, *new_style, style)
	err := plot.ResetPointGroupStyle("Sam", "lines")
	if err == nil {
		t.Error("The specified pointgroup to be reset does not exist")
	}
}
