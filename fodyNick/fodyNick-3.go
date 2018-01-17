package fodyNick

import "math"

// you have a box that is valid and has a non zero area
// therefore it has an aspect ratio
// aspect ratio is defined by boxWidth/boxHeight
// given a box and a new aspect ratio
// convert the aspect ratio of the existing box
// to the new aspect ratio
// only change height or the width.
// and keep the old content centered

func changeRatio(b *box, ratio float32) {
	height := b.maxY - b.minX
	width := b.maxY - b.minY
	originalRatio := float32(width)/float32(height)
	if originalRatio == ratio {
		return
	}
	// ar = w / h
	if originalRatio < ratio {
		// find width.
		// w = ar * h
		newWidthIncrease := ratio*height - width / 2
		b.minX -= newWidthIncrease
		b.maxX += newWidthIncrease
	} else {
		// h = w/ar
		newHeightIncrease := width/ratio - height / 2
		b.minY -= newHeightIncrease
		b.maxY += newHeightIncrease
	}
}

