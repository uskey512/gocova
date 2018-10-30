package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/urfave/cli"
)

var (
	flags = []cli.Flag{
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
			Usage: "saturation offset [-100.0 ... 100.0]",
			Value: 0,
		},
		cli.Float64Flag{
			Name:  "lightness, l",
			Usage: "lightness offset [-100.0 ... 100.0]",
			Value: 0,
		},
	}
)

type filter struct {
	hue        float64
	saturation float64
	lightness  float64

	isGrayscale bool
	isBitonal   bool
}

type imageFile struct {
	image  image.Image
	format string
	bounds image.Rectangle
}

func getDstPathBase(pathbase string, ext string, pattern int) string {
	return fmt.Sprintf("%s_%%0%dd.%s", pathbase, len(strconv.Itoa(pattern)), ext)
}

func readImage(inputImage string) imageFile {
	reader, err := os.Open(inputImage)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	srcImage, format, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return imageFile{
		image:  srcImage,
		format: format,
		bounds: srcImage.Bounds(),
	}
}

func writeImage(dst imageFile, dstPath string) {
	dstFile, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	switch dst.format {
	case "png":
		png.Encode(dstFile, image.Image(dst.image))
		break
	case "jpeg":
		jpeg.Encode(dstFile, image.Image(dst.image), nil)
		break
	case "gif":
		gif.Encode(dstFile, image.Image(dst.image), nil)
		break
	}
}

func filterImage(src *imageFile, filter filter) imageFile {
	filteredImage := image.NewRGBA(src.bounds)

	for y := src.bounds.Min.Y; y < src.bounds.Max.Y; y++ {
		for x := src.bounds.Min.X; x < src.bounds.Max.X; x++ {
			orgColor := src.image.At(x, y)

			colorfulColor, _ := colorful.MakeColor(orgColor)
			h, s, l := colorfulColor.Hsl()

			h = math.Mod(h+float64(filter.hue), 360.0)
			s = Clamp01(s + filter.saturation)
			l = Clamp01(l + filter.lightness)

			if filter.isGrayscale {
				s = 0.5
			}
			if filter.isBitonal {
				l = 0.5
			}

			resultHsv := colorful.Hsl(h, s, l)

			r, g, b := resultHsv.RGB255()

			if src.format == "png" {
				filteredImage.Set(x, y, color.NRGBA{r, g, b, orgColor.(color.NRGBA).A})
			} else {
				filteredImage.Set(x, y, color.RGBA{r, g, b, 255})
			}
		}
	}

	return imageFile{
		image:  filteredImage,
		format: src.format,
		bounds: src.bounds,
	}
}

func parseFilter(c *cli.Context) (filter, int) {
	pattern := c.Int("pattern")
	h := float64(360.0 / (pattern + 1))
	s := Clamp(c.Float64("saturation"), -100.0, 100.0) / 100.0
	l := Clamp(c.Float64("lightness"), -100.0, 100.0) / 100.0

	return filter{
		hue:        h,
		saturation: s,
		lightness:  l,
	}, pattern
}

func process(c *cli.Context) {
	src := readImage(c.Args().Get(0))
	dstPath := c.String("output")
	filter, pattern := parseFilter(c)
	hueInterval := filter.hue

	filter.isBitonal = bitonalDetect(src)
	filter.isGrayscale = grayscaleDetect(src)

	dstFilePathBase := getDstPathBase(dstPath, src.format, pattern)
	for i := 1; i <= pattern; i++ {
		dstFilePath := fmt.Sprintf(dstFilePathBase, i)
		filter.hue = hueInterval * float64(i)

		writeImage(filterImage(&src, filter), dstFilePath)
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "gocova"
	app.Usage = "Go color variation, generate images of various colors"
	app.Version = "1.4.0"

	app.Action = process
	app.Flags = flags
	app.Run(os.Args)
}
