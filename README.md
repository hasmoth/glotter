 [![Build Status](https://travis-ci.org/Arafatk/glot.svg?branch=master)](https://travis-ci.org/Arafatk/glot) [![GoDoc](https://godoc.org/github.com/arafat/glot?status.svg)](https://godoc.org/github.com/Arafatk/glot) [![Join the chat at https://gitter.im/tensorflowrb/Lobby](https://badges.gitter.im/tensorflowrb/Lobby.svg)](https://gitter.im/glot-dev/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)  ![License](https://img.shields.io/badge/License-MIT-blue.svg)  

# General
This fork maily attempts to solve the problem that several 2D and 3D gnuplot plotting styles accept multi-dimensional data for various reasons (errorbars, circle diameter etc.). The same is true for 1D plotting styles.
Secondly, gnuplot natievly allows for extensive configuration of plotted data. This is an attempt to address this as well.

The general method of this library is to construct a temporary file from the data that is to be plottet and create a gnuplot command string that is to be executed by gnuplot.

# Glotter
`glotter` is a plotting library for Golang built on top of [gnuplot](http://www.gnuplot.info/). `glot` currently supports styles like lines, points, bars, steps, histogram, circle, and many others. We are continuously making efforts to add more features.  

## Documentation
Documentation is available at [godoc](https://godoc.org/github.com/Arafatk/glot).

## Requirements
 - gnu plot
    - build gnu plot from [source](https://sourceforge.net/projects/gnuplot/files/gnuplot/)
    - linux users
       -  ```sudo apt-get update```
       -  ```sudo apt-get install gnuplot-x11```
    - mac users
       -  install homebrew
       -  ```brew cask install xquartz``` (for x-11)
       -  ```brew install gnuplot --with-x11```

## Installation
```go get github.com/hasmoth/glotter```

## Usage and Examples
In addition to the basic usage, a method was added that allows for configuration of the plotted data.
```
	new_style := NewPlotObjectStyle(
		SetPointType(6),
		SetPointSize(1.5),
		SetLineColor("rgb", "grey"),
		SetDashType(4),
		SetLineWidth(2),
		SetLineType(5),
	)
```
This style object allows for optional parameters and can be passed to a ```PointGroup```.
```plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11}, *new_style)```

Furthermore, ever plotting style (e.g. points, lines etc.) has a defined maximum number of allowed data columns.

## Examples
![](https://raw.githubusercontent.com/Arafatk/plot/master/Screenshot%20-%20Saturday%2014%20October%202017%20-%2004-51-13%20%20IST.png)

## Contributing
We really encourage developers coming in, finding a bug or requesting a new feature. Want to tell us about the feature you just implemented, just raise a pull request and we'll be happy to go through it. Please read the CONTRIBUTING and CODE_OF_CONDUCT file.

