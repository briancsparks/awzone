package awzone

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type ProgressCounter struct {
	Total int64
	Count int64

	Bar *progressbar.ProgressBar
}

func (pc *ProgressCounter) Write(p []byte) (int, error) {
	n := int64(len(p))
	pc.Count += n

	//fmt.Printf("Read %d bytes for a total of %d/%d (%6.2f%%)\n", n, pc.Count, pc.Total, 100.0*float64(pc.Count)/float64(pc.Total))
	pc.Bar.Add64(n)

	return int(n), nil
}

func NewProgressCounter(size int64, x, y int, descr string) *ProgressCounter {
	description := fmt.Sprintf("[cyan][%d/%d][reset] %s", x, y, descr)

	return &ProgressCounter{
		Total: size,
		Bar: progressbar.NewOptions64(size,
			progressbar.OptionUseANSICodes(true),
			progressbar.OptionEnableColorCodes(true),

			//progressbar.OptionSetWidth(18),
			progressbar.OptionFullWidth(),

			progressbar.OptionSetRenderBlankState(true),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetElapsedTime(false),
			progressbar.OptionSetPredictTime(true),

			progressbar.OptionShowCount(),
			//progressbar.OptionShowIts(),

			progressbar.OptionSetDescription(description),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[green]=[reset]",
				SaucerHead:    "[green]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		),
	}
}
