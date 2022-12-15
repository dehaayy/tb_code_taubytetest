package lib

import (
	"bitbucket.org/taubyte/go-sdk/event"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//export ping
func ping(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	// Create a new plot
	p, err := plot.New()
	if err != nil {
		return 1
	}

	// Create a new line plotter
	l, err := plotter.NewLine(plotter.XYs{{0, 0}, {1, 1}, {2, 4}, {3, 9}})
	if err != nil {
		return 1
	}

	// Add the line plotter to the plot
	p.Add(l)

	// Write the plot to the HTTP response
	err = p.Write(10*vg.Centimeter, 10*vg.Centimeter, h)
	if err != nil {
		return 1
	}

	// Write the "heyy" text to the HTTP response
	h.Write([]byte("heyy"))

	return 0
}


