package glot

import (
	"fmt"
	"slices"
)

// SetTitle sets the title for the plot
//
// Usage
//
//	dimensions := 3
//	persist := false
//	debug := false
//	plot, _ := glot.NewPlot(dimensions, persist, debug)
//	plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	plot.SetTitle("Test Results")
func (plot *Plot) SetTitle(title string) error {
	return plot.Cmd("set title \"%s\" ", title)
	// return plot.Cmd(fmt.Sprintf("set title \"%s\" ", title))
}

// SetXLabel changes the label for the x-axis
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetXLabel("X-Axis")
func (plot *Plot) SetXLabel(label string) error {
	return plot.Cmd("set xlabel '%s'", label)
}

// SetYLabel changes the label for the y-axis
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetYLabel("Y-Axis")
func (plot *Plot) SetYLabel(label string) error {
	return plot.Cmd("set ylabel '%s'", label)
}

// SetZLabel changes the label for the z-axis
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetZLabel("Z-Axis")
func (plot *Plot) SetZLabel(label string) error {
	return plot.Cmd("set zlabel '%s'", label)
}

// SetLabels Functions helps to set labels for x, y, z axis  simultaneously
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetLabels("X-axis","Y-Axis","Z-Axis")
func (plot *Plot) SetLabels(labels ...string) error {
	ndims := len(labels)
	if ndims > 3 || ndims <= 0 {
		return &gnuplotError{fmt.Sprintf("invalid number of dims '%v'", ndims)}
	}
	var err error

	for i, label := range labels {
		switch i {
		case 0:
			ierr := plot.SetXLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		case 1:
			ierr := plot.SetYLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		case 2:
			ierr := plot.SetZLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		}
	}
	return nil
}

// SetMXtics sets the minor tick marks of the x-axis
//
// Usage
//
//	SetMXtics(5)
func (plot *Plot) SetMXtics(n int) error {
	return plot.Cmd("set mxtics %d", n)
}

// SetMYtics sets the minor tick marks of the y-axis
//
// Usage
//
//	SetMYtics(5)
func (plot *Plot) SetMYtics(n int) error {
	return plot.Cmd("set mytics %d", n)
}

// SetGrid sets a grid to a plot
//
// Usage
//
//	SetGrid("") : default style
//	SetGrid(fmt.Sprintf("lt %d lc rgb \"black\"", 1))
func (plot *Plot) SetGrid(format string) error {
	if len(format) == 0 {
		format = fmt.Sprintf("lt %d lc rgb \"grey\"", 1)
	}
	return plot.Cmd("set grid %s", format)
}

// SetXrange changes the range for the x-axis
//
// Usage
//
//	dimensions := 3
//	persist := false
//	debug := false
//	plot, _ := glot.NewPlot(dimensions, persist, debug)
//	plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	plot.SetTitle("Test Results")
//	plot.SetXrange(-2,2)
func (plot *Plot) SetXrange(start int, end int) error {
	return plot.Cmd("set xrange [%d:%d]", start, end)
}

// SetLogscale changes the scale of an axis to log
//
// Usage
//
//	dimensions := 3
//	persist := false
//	debug := false
//	plot, _ := glot.NewPlot(dimensions, persist, debug)
//	plot.SetYrange(-2, 18)
//	plot.AddPointGroup("rates", "circle", [][]float64{{2, 4, 8, 16, 32}, {4, 7, 4, 10, 3}})
//	plot.SetLogscale("x", 2)
func (plot *Plot) SetLogscale(axis string, base int) error {
	return plot.Cmd("set logscale %s %d", axis, base)
}

// SetYrange changes the range for the y-axis
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetYrange(-2,2)
func (plot *Plot) SetYrange(start int, end int) error {
	return plot.Cmd("set yrange [%d:%d]", start, end)
}

// SetZrange changes the range for the z-axis
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetZrange(-2,2)
func (plot *Plot) SetZrange(start int, end int) error {
	return plot.Cmd("set zrange [%d:%d]", start, end)
}

// SavePlot function is used to save the plot at this point.
// The plot is dynamic and additional pointgroups can be added and removed and different versions
// of the same plot can be saved.
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetZrange(-2,2)
//	 plot.SavePlot("1.jpeg")
func (plot *Plot) SavePlot(filename string) (err error) {
	if plot.nplots == 0 {
		return &gnuplotError{"This plot has 0 curves and therefore its a redundant plot and it can't be printed."}
	}
	plot.CheckedCmd("set terminal %s", plot.format)
	plot.CheckedCmd("set output '%s'", filename)
	plot.CheckedCmd("replot  ")
	return nil
}

// SetFormat function is used to save the plot at this point.
// The plot is dynamic and additional pointgroups can be added and removed and different versions
// of the same plot can be saved.
//
// Usage
//
//	 dimensions := 3
//	 persist := false
//	 debug := false
//	 plot, _ := glot.NewPlot(dimensions, persist, debug)
//	 plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
//	 plot.SetTitle("Test Results")
//		plot.SetFormat("pdf")
//	 plot.SavePlot("1.pdf")
//
// NOTE: png is default format for saving files.
func (plot *Plot) SetFormat(newformat string) error {
	allowed := []string{
		"png", "pdf"}
	if slices.Contains(allowed, newformat) {
		plot.format = newformat
		return nil
	}
	fmt.Printf("** Format '%v' not in allowed list %v\n", newformat, allowed)
	fmt.Printf("** default to 'png'\n")
	err := &gnuplotError{fmt.Sprintf("invalid format '%s'", newformat)}
	return err
}
