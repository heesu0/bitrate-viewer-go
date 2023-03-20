package main

import (
	"github.com/heesu0/bitrate-viewer-go/pkg/ffprobe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func getPoints(probeData ffprobe.ProbeData) (plotter.XYs, plotter.XYs, error) {
	bitrates := make(plotter.XYs, len(probeData.Frames))
	iFrames := make(plotter.XYs, 0)

	for i, frame := range probeData.Frames {
		bitrates[i].X = float64(i)
		bitrates[i].Y = float64(frame.PktSize)

		if frame.PictType == "I" {
			iFramePoint := plotter.XY{
				X: float64(i),
				Y: float64(frame.PktSize),
			}
			iFrames = append(iFrames, iFramePoint)
		}
	}

	return bitrates, iFrames, nil
}

func plotGraph(bitrates, iFrames plotter.XYs, outputFile string) error {
	p := plot.New()
	p.Title.Text = "Bitrate and I-Frames"
	p.X.Label.Text = "Frame"
	p.Y.Label.Text = "Size (bytes)"

	err := plotutil.AddLines(p, "Bitrate", bitrates)
	if err != nil {
		return err
	}

	scatter, err := plotter.NewScatter(iFrames)
	if err != nil {
		return err
	}

	scatter.GlyphStyle.Color = plotutil.Color(1)
	scatter.Shape = plotutil.Shape(1)

	p.Add(scatter)
	p.Legend.Add("I-Frames", scatter)

	err = p.Save(10*vg.Inch, 5*vg.Inch, outputFile)
	if err != nil {
		return err
	}

	return nil
}
