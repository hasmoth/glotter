package glot

import "fmt"

// with <style> { {linestyle | ls <line_style>}
// 	|  {{linetype   | lt <line_type>}
// 		{linewidth  | lw <line_width>}
// 		{linecolor  | lc <colorspec>}
// 		{pointtype  | pt <point_type>}
// 		{pointsize  | ps <point_size>}
// 		{arrowstyle | as <arrowstyle_index>}
// 		{fill | fs <fillstyle>} {fillcolor | fc <colorspec>}
// 		{nohidden3d} {nocontours} {nosurface}
// 		{palette}}
// }

type PlotObjectType struct {
	Name  string
	Value int
}

type PlotObjectSize struct {
	Name  string
	Value float64
}

type PlotObjectColor struct {
	Name      string
	ColorSpec string
	Value     string
}

type PlotObjectStyle struct {
	PointType *PlotObjectType
	PointSize *PlotObjectSize
	LineColor *PlotObjectColor
	LineType  *PlotObjectType
	LineWidth *PlotObjectSize
	DashType  *PlotObjectType
}

type PlotObjectOptions func(*PlotObjectStyle)

func prependWhitespace(buf *string) {
	if buf != nil && len(*buf) != 0 {
		*buf += " "
	}
}

func (s PlotObjectStyle) String() string {
	var object_style string
	if s.PointType != nil {
		prependWhitespace(&object_style)
		object_style += fmt.Sprintf("%s %d", s.PointType.Name, s.PointType.Value)
	}
	if s.PointSize != nil {
		prependWhitespace(&object_style)
		object_style += fmt.Sprintf("%s %f", s.PointSize.Name, s.PointSize.Value)
	}
	if s.LineColor != nil {
		prependWhitespace(&object_style)
		object_style += fmt.Sprintf("%s %s \"%s\"", s.LineColor.Name, s.LineColor.ColorSpec, s.LineColor.Value)
	}
	if s.LineType != nil {
		prependWhitespace(&object_style)
		object_style += fmt.Sprintf("%s %d", s.LineType.Name, s.LineType.Value)
	}
	if s.LineWidth != nil {
		prependWhitespace(&object_style)
		object_style += fmt.Sprintf("%s %f", s.LineWidth.Name, s.LineWidth.Value)
	}
	if s.DashType != nil {
		prependWhitespace(&object_style)
		object_style += fmt.Sprintf("%s %d", s.DashType.Name, s.DashType.Value)
	}
	return object_style
}

func SetPointType(pt int) PlotObjectOptions {
	return func(s *PlotObjectStyle) {
		s.PointType = &PlotObjectType{}
		s.PointType.Name = "pt"
		s.PointType.Value = pt
	}
}

func SetPointSize(ps float64) PlotObjectOptions {
	return func(s *PlotObjectStyle) {
		s.PointSize = &PlotObjectSize{}
		s.PointSize.Name = "ps"
		s.PointSize.Value = ps
	}
}

func SetLineWidth(lw float64) PlotObjectOptions {
	return func(s *PlotObjectStyle) {
		s.LineWidth = &PlotObjectSize{}
		s.LineWidth.Name = "lw"
		s.LineWidth.Value = lw
	}
}

func SetLineType(lt int) PlotObjectOptions {
	return func(s *PlotObjectStyle) {
		s.LineType = &PlotObjectType{}
		s.LineType.Name = "lt"
		s.LineType.Value = lt
	}
}

func SetLineColor(cs string, c string) PlotObjectOptions {
	return func(s *PlotObjectStyle) {
		s.LineColor = &PlotObjectColor{}
		s.LineColor.Name = "lc"
		s.LineColor.ColorSpec = cs
		s.LineColor.Value = c
	}
}

func SetDashType(dt int) PlotObjectOptions {
	return func(s *PlotObjectStyle) {
		s.DashType = &PlotObjectType{}
		s.DashType.Name = "dt"
		s.DashType.Value = dt
	}
}

// Contructor for a plot object style with optional parameters
//
// Usage
//
//	plot_style := NewPlotObjectStyle(
//		SetPointType(6),
//		SetPointSize(1.5),
//		SetLineColor("rgb", "grey"),
//	)
func NewPlotObjectStyle(options ...PlotObjectOptions) *PlotObjectStyle {
	pos := &PlotObjectStyle{}

	for _, option := range options {
		option(pos)
	}
	return pos
}
