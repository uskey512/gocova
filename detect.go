package main

import "github.com/lucasb-eyer/go-colorful"

const bitonalWhiteThreshold = 0.99
const bitonalBlackThreshold = 0.01
const bitonalThreshold = 0.9999

func bitonalDetect(src imageFile) bool {
	pixelCount := src.bounds.Max.Y * src.bounds.Max.X
	bitoneCount := 0

	for y := src.bounds.Min.Y; y < src.bounds.Max.Y; y++ {
		for x := src.bounds.Min.X; x < src.bounds.Max.X; x++ {
			orgColor := src.image.At(x, y)

			colorfulColor, _ := colorful.MakeColor(orgColor)
			_, _, l := colorfulColor.Hsl()

			if l < bitonalBlackThreshold || bitonalWhiteThreshold < l {
				bitoneCount++
			}
		}
	}

	return bitonalThreshold < (float64(bitoneCount) / float64(pixelCount))
}

func grayscaleDetect(src imageFile) bool {
	pixelCount := src.bounds.Max.Y * src.bounds.Max.X
	bitoneCount := 0

	for y := src.bounds.Min.Y; y < src.bounds.Max.Y; y++ {
		for x := src.bounds.Min.X; x < src.bounds.Max.X; x++ {
			orgColor := src.image.At(x, y)

			colorfulColor, _ := colorful.MakeColor(orgColor)
			_, s, _ := colorfulColor.Hsl()

			if s == 0.0 {
				bitoneCount++
			}
		}
	}

	return bitoneCount == pixelCount
}
