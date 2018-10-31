package main

import "testing"

func TestGrayscaleDetect(t *testing.T) {
	image := readImage("testdata/normal._png")
	if grayscaleDetect(image) != false {
		t.Fail()
	}

	image = readImage("testdata/grayscale._png")
	if grayscaleDetect(image) != true {
		t.Fail()
	}

	image = readImage("testdata/bitonal._png")
	if grayscaleDetect(image) != true {
		t.Fail()
	}
}

func TestBitonalDetect(t *testing.T) {
	image := readImage("testdata/normal._png")
	if bitonalDetect(image) != false {
		t.Fail()
	}

	image = readImage("testdata/grayscale._png")
	if bitonalDetect(image) != false {
		t.Fail()
	}

	image = readImage("testdata/bitonal._png")
	if bitonalDetect(image) != true {
		t.Fail()
	}
}
