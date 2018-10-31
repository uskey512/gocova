package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNormalImage(t *testing.T) {
	src := readImage("testdata/normal._png")
	pattern := 3
	filter := filter{
		hue:        float64(360.0 / (pattern + 1)),
		saturation: 0,
		lightness:  0,
	}

	hueInterval := filter.hue

	filter.isBitonal = bitonalDetect(src)
	filter.isGrayscale = grayscaleDetect(src)

	var resultImage imageFile
	for i := 1; i <= pattern; i++ {
		filter.hue = hueInterval * float64(i)
		writeImage(filterImage(&src, filter), "/tmp/testdata.png")

		testImage := readImage("/tmp/testdata.png")
		resultImage = readImage(fmt.Sprintf("testdata/result_normal/filtered_%d._png", i))

		if !reflect.DeepEqual(testImage.image, resultImage.image) {
			t.Fail()
		}
	}

	os.Remove("/tmp/testdata.png")
}

func TestGrayscaleImage(t *testing.T) {
	src := readImage("testdata/grayscale._png")
	pattern := 3
	filter := filter{
		hue:        float64(360.0 / (pattern + 1)),
		saturation: 0,
		lightness:  0,
	}

	hueInterval := filter.hue

	filter.isBitonal = bitonalDetect(src)
	filter.isGrayscale = grayscaleDetect(src)

	var resultImage imageFile
	for i := 1; i <= pattern; i++ {
		filter.hue = hueInterval * float64(i)
		writeImage(filterImage(&src, filter), "/tmp/testdata.png")

		testImage := readImage("/tmp/testdata.png")
		resultImage = readImage(fmt.Sprintf("testdata/result_grayscale/filtered_%d._png", i))

		if !reflect.DeepEqual(testImage.image, resultImage.image) {
			t.Fail()
		}
	}

	os.Remove("/tmp/testdata.png")
}

func TestBitonalImage(t *testing.T) {
	src := readImage("testdata/bitonal._png")
	pattern := 3
	filter := filter{
		hue:        float64(360.0 / (pattern + 1)),
		saturation: 0,
		lightness:  0,
	}

	hueInterval := filter.hue

	filter.isBitonal = bitonalDetect(src)
	filter.isGrayscale = grayscaleDetect(src)

	var resultImage imageFile
	for i := 1; i <= pattern; i++ {
		filter.hue = hueInterval * float64(i)
		writeImage(filterImage(&src, filter), "/tmp/testdata.png")

		testImage := readImage("/tmp/testdata.png")
		resultImage = readImage(fmt.Sprintf("testdata/result_bitonal/filtered_%d._png", i))

		if !reflect.DeepEqual(testImage.image, resultImage.image) {
			t.Fail()
		}
	}

	os.Remove("/tmp/testdata.png")
}
