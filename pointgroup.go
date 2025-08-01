package glot

import (
	"fmt"
)

// Requirement for each plot style concerning dimensionality and number of columns.
var plotting_styles = map[string]map[int]int{
	"lines":          {2: 2, 3: 3},
	"points":         {2: 2, 3: 3},
	"linepoints":     {2: 2, 3: 3},
	"dots":           {2: 2, 3: 3},
	"circles":        {2: 6},
	"boxerrorbars":   {2: 5},
	"boxxyerrorbars": {2: 6},
	"boxes":          {2: 3, 3: 5},
	"candlesticks":   {2: 7},
	"filledcurves":   {2: 3},
	"financebars":    {2: 5},
	"histograms":     {2: 3},
	"hsteps":         {2: 3},
	"histeps":        {2: 2},
	"image":          {2: 3, 3: 4},
	"impulses":       {2: 3},
	"labels":         {2: 3, 3: 4},
	"linespoints":    {2: 2, 3: 3},
	"rgbimage":       {2: 3, 3: 4},
	"fsteps":         {2: 2},
	"steps":          {2: 2},
	"vectors":        {2: 5, 3: 7},
	"xerrorbars":     {2: 4},
	"xyerrorbars":    {2: 6},
	"yerrorbars":     {2: 4},
	"xyerrorlines":   {2: 6},
	"xerrorlines":    {2: 4},
	"yerrorlines":    {2: 4},
	// "pm3d":           true,
	// "table":       true,
}

// A PointGroup refers to a set of points that need to plotted.
// It could either be a set of points or a function of co-ordinates.
// For Example z = Function(x,y)(3 Dimensional) or  y = Function(x) (2-Dimensional)
type PointGroup struct {
	name             string            // Name of the curve
	dimensions       int               // dimensions of the curve
	style            string            // current plotting style
	data             any               // Data inside the curve in any integer/float format
	castedData       any               // The data inside the curve typecasted to float64
	set              bool              // TODO: unused
	plotObjectStyles []PlotObjectStyle // style of the plotted data
}

// AddPointGroup function adds a group of points to a plot.
//
// Usage
//
//	dimensions := 2
//	persist := false
//	debug := false
//	plot, _ := glot.NewPlot(dimensions, persist, debug)
//	plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11})
//	plot.AddPointGroup("Sample2", "points", []int32{1, 2, 4, 11})
//	plot.SavePlot("1.png")
func (plot *Plot) AddPointGroup(name string, style string, data any, spec ...PlotObjectStyle) (err error) {
	_, exists := plot.PointGroup[name]
	if exists {
		return &gnuplotError{fmt.Sprintf("A PointGroup with the name %s  already exists, please use another name of the curve or remove this curve before using another one with the same name.", name)}
	}

	curve := &PointGroup{name: name, dimensions: plot.dimensions, data: data, set: true, plotObjectStyles: spec}
	var allowed []string
	curve.style = defaultStyle
	discovered := 0
	for s := range plotting_styles {
		allowed = append(allowed, s)
	}
	var max_cols int
	if d, ok := plotting_styles[style]; ok {
		var ok bool
		if max_cols, ok = d[curve.dimensions]; ok {
			curve.style = style
			err = nil
			discovered = 1
		}
	}

	if discovered == 0 {
		fmt.Printf("** style '%v' not in allowed list %v\n", style, allowed)
		fmt.Printf("** default to 'points'\n")
		return &gnuplotError{fmt.Sprintf("invalid style '%s'", style)}
	}
	switch d := data.(type) {
	case [][]float64:
		if max_cols >= len(d) {
			return &gnuplotError{"The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot."}
		}
		curve.castedData = d
		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]float32:
		if max_cols >= len(d) {
			return &gnuplotError{"The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot."}
		}
		originalSlice := d
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int:
		if max_cols >= len(d) {
			return &gnuplotError{"The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot."}
		}
		originalSlice := d
		if len(originalSlice) != 2 {
			return &gnuplotError{"this is not a 2d matrix"}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int8:
		if max_cols >= len(d) {
			return &gnuplotError{"The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot."}
		}
		originalSlice := d
		if len(originalSlice) != 2 {
			return &gnuplotError{"this is not a 2d matrix"}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int16:
		if max_cols >= len(d) {
			return &gnuplotError{"The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot."}
		}
		originalSlice := d
		if len(originalSlice) != 2 {
			return &gnuplotError{"this is not a 2d matrix"}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int32:
		if max_cols >= len(d) {
			return &gnuplotError{"The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot."}
		}
		originalSlice := d
		if len(originalSlice) != 2 {
			return &gnuplotError{"this is not a 2d matrix"}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int64:
		if max_cols >= len(d) {
			return &gnuplotError{"The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot."}
		}
		originalSlice := d
		if len(originalSlice) != 2 {
			return &gnuplotError{"this is not a 2d matrix"}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case []float64:
		curve.castedData = d
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []float32:
		originalSlice := d
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int:
		originalSlice := d
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int8:
		originalSlice := d
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int16:
		originalSlice := d
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int32:
		originalSlice := d
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int64:
		originalSlice := d
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := range originalSlice {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	default:
		return &gnuplotError{"invalid number of dims "}

	}
	return err
}

// RemovePointGroup helps to remove a particular point group from the plot.
// This way you can remove a pointgroup if it's un-necessary.
//
// Usage
//
//	dimensions := 3
//	persist := false
//	debug := false
//	plot, _ := glot.NewPlot(dimensions, persist, debug)
//	plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11})
//	plot.AddPointGroup("Sample2", "points", []int32{1, 2, 4, 11})
//	plot.RemovePointGroup("Sample1")
func (plot *Plot) RemovePointGroup(name string) {
	delete(plot.PointGroup, name)
	plot.cleanplot()
	for _, pointGroup := range plot.PointGroup {
		plot.plotX(pointGroup)
	}
}

// ResetPointGroupStyle helps to reset the style of a particular point group in a plot.
// Using both AddPointGroup and RemovePointGroup you can add or remove point groups.
// And dynamically change the plots.
//
// Usage
//
//	dimensions := 2
//	persist := false
//	debug := false
//	plot, _ := glot.NewPlot(dimensions, persist, debug)
//	plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11})
//	plot.ResetPointGroupStyle("Sample1", "points")
func (plot *Plot) ResetPointGroupStyle(name string, style string) (err error) {
	pointGroup, exists := plot.PointGroup[name]
	if !exists {
		return &gnuplotError{fmt.Sprintf("A curve with name %s does not exist.", name)}
	}
	plot.RemovePointGroup(name)
	pointGroup.style = style
	plot.plotX(pointGroup)
	return err
}
