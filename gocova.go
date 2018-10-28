package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/urfave/cli"
)

var (
	flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "input image path",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output image path base",
			Value: "./result",
		},
		cli.IntFlag{
			Name:  "pattern, p",
			Usage: "number of images to generate",
			Value: 10,
		},
		cli.Float64Flag{
			Name:  "saturation, s",
			Usage: "saturation offset [-100.0...100.0]",
			Value: 0,
		},
		cli.Float64Flag{
			Name:  "lightness, l",
			Usage: "lightness offset [-100.0...100.0]",
			Value: 0,
		},
	}
)

type HslOffset struct {
	h, s, l float64
}

func clamp(v, max, min float64) float64 {
	if max < v {
		return max
	}
	if v < min {
		return min
	}
	return v
}

func clamp01(v float64) float64 {
	return clamp(v, 1.0, 0.0)
}

func getCarry(v int) int {
	c := 0
	for ; v != 0; c++ {
		v /= 10
	}
	return c
}

func getDstPathBase(pathbase string, pattern int) string {
	return fmt.Sprintf("%s_%%0%dd.png", pathbase, getCarry(pattern))
}

func loadImage(inputImage string) image.Image {
	reader, err := os.Open(inputImage)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	srcImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return srcImage
}

func generateImage(srcImage image.Image, dstPath string, offset HslOffset) {
	dstFile, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	srcBounds := srcImage.Bounds()
	filteredImage := image.NewRGBA(srcBounds)

	for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
		for x := srcBounds.Min.X; x < srcBounds.Max.X; x++ {
			orgColor := srcImage.At(x, y).(color.NRGBA)

			colorfulColor, _ := colorful.MakeColor(orgColor)
			h, s, l := colorfulColor.Hsl()

			h = math.Mod(h+float64(offset.h), 360.0)
			s = clamp01(s + offset.s)
			l = clamp01(l + offset.l)

			resultHsv := colorful.Hsl(h, s, l)

			r, g, b := resultHsv.RGB255()

			filteredImage.Set(x, y, color.NRGBA{r, g, b, orgColor.A})
		}
	}
	png.Encode(dstFile, image.Image(filteredImage))
}

func process(c *cli.Context) {
	image := loadImage(c.String("input"))
	dstPath := c.String("output")

	pattern := c.Int("pattern")
	hInterval := float64(360 / (pattern + 1))
	saturation := clamp(c.Float64("saturation"), 100.0, -100.0) / 100.0
	lightness := clamp(c.Float64("lightness"), 100.0, -100.0) / 100.0

	offset := HslOffset{hInterval, saturation, lightness}

	dstFilePathBase := getDstPathBase(dstPath, pattern)
	for i := 1; i <= pattern; i++ {
		dstFilePath := fmt.Sprintf(dstFilePathBase, i)
		offset.h = hInterval * float64(i)
		generateImage(image, dstFilePath, offset)
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "gocova"
	app.Usage = "Go color variation, generate images of various colors"
	app.Version = "1.1.0"

	app.Action = process
	app.Flags = flags
	app.Run(os.Args)
}
