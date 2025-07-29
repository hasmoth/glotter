package glot

import (
	"fmt"
	"testing"
)

func TestPlotObjectStyle(t *testing.T) {
	new_style := NewPlotObjectStyle(
		SetPointType(6),
		SetPointSize(1.5),
		SetLineColor("rgb", "grey"),
		SetDashType(4),
		SetLineWidth(2),
		SetLineType(5),
	)

	if new_style.PointType.Name != "pt" || new_style.PointType.Value != 6 {
		t.Errorf("Wrong pointtype: %s %d", new_style.PointType.Name, new_style.PointType.Value)
	}
	if new_style.DashType.Name != "dt" || new_style.DashType.Value != 4 {
		t.Errorf("Wrong pointtype: %s %d", new_style.DashType.Name, new_style.DashType.Value)
	}
	if new_style.LineType.Name != "lt" || new_style.LineType.Value != 5 {
		t.Errorf("Wrong pointtype: %s %d", new_style.LineType.Name, new_style.LineType.Value)
	}
	if new_style.PointSize.Name != "ps" || new_style.PointSize.Value != 1.5 {
		t.Errorf("Wrong pointtype: %s %f", new_style.PointSize.Name, new_style.PointSize.Value)
	}
	if new_style.LineWidth.Name != "lw" || new_style.LineWidth.Value != 2 {
		t.Errorf("Wrong pointtype: %s %f", new_style.LineWidth.Name, new_style.LineWidth.Value)
	}
	if new_style.LineColor.Name != "lc" || new_style.LineColor.Value != "grey" {
		t.Errorf("Wrong pointtype: %s %s %s", new_style.LineColor.Name, new_style.LineColor.ColorSpec, new_style.LineColor.Value)
	}
	fmt.Println(new_style)
}
