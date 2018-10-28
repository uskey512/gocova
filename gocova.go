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
		cli.Int64Flag{
			Name:  "saturation, s",
			Usage: "color saturation offset [-100...100]",
			Value: 0,
		},
	}
)

func clamp01(v float64) float64 {
	if 1.0 < v {
		return 1.0
	}
	if v < 0.0 {
		return 0.0
	}
	return v
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

func generateImage(srcImage image.Image, dstPath string, rotation int, colorSaturation float64) {
	dstFile, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	srcBounds := srcImage.Bounds()
	filteredImage := image.NewRGBA(srcBounds)

	for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
		for x := srcBounds.Min.X; x < srcBounds.Max.X; x++ {
			orgColor := srcImage.At(x, y)

			colorfulColor, _ := colorful.MakeColor(orgColor)
			h, s, v := colorfulColor.Hsl()

			h = math.Mod(h+float64(rotation), 360.0)
			s = clamp01(s + colorSaturation)
			resultHsv := colorful.Hsl(h, s, v)

			r, g, b := resultHsv.RGB255()
			_, _, _, a := orgColor.RGBA()

			filteredImage.Set(x, y, color.RGBA{r, g, b, uint8(a)})
		}
	}
	png.Encode(dstFile, image.Image(filteredImage))
}

func process(c *cli.Context) {
	image := loadImage(c.String("input"))
	dstPath := c.String("output")
	pattern := c.Int("pattern")
	colorSaturation := c.Int("saturation")
	hslDegree := 360 / (pattern + 1)

	for i := 1; i <= pattern; i++ {
		dstFilePath := fmt.Sprintf("%s_%03d.png", dstPath, i)
		generateImage(image, dstFilePath, i*hslDegree, float64(colorSaturation)/100.0)
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "gocova"
	app.Usage = "Go color variation, generate images of various colors"
	app.Version = "1.0.0"

	app.Action = process
	app.Flags = flags
	app.Run(os.Args)
}
