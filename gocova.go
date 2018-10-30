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
		cli.BoolFlag{
			Name:  "grayscale, g",
			Usage: "input image is grayscale\n\tsaturation of fixed value : [50.0]",
		},
	}
)

type hslOption struct {
	h, s, l float64
	g       bool
}

func getDstPathBase(pathbase string, ext string, pattern int) string {
	return fmt.Sprintf("%s_%%0%dd.%s", pathbase, len(strconv.Itoa(pattern)), ext)
}

func loadImage(inputImage string) (image.Image, string) {
	reader, err := os.Open(inputImage)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	srcImage, format, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return srcImage, format
}

func generateImage(srcImage image.Image, format, dstPath string, offset hslOption) {
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
			h, s, l := colorfulColor.Hsl()

			h = math.Mod(h+float64(offset.h), 360.0)
			s = Clamp01(s + offset.s)
			l = Clamp01(l + offset.l)

			if offset.g {
				s = 0.5
			}

			resultHsv := colorful.Hsl(h, s, l)

			r, g, b := resultHsv.RGB255()

			if format == "png" {
				filteredImage.Set(x, y, color.NRGBA{r, g, b, orgColor.(color.NRGBA).A})
			} else {
				filteredImage.Set(x, y, color.RGBA{r, g, b, 255})
			}
		}
	}

	switch format {
	case "png":
		png.Encode(dstFile, image.Image(filteredImage))
		break
	case "jpeg":
		jpeg.Encode(dstFile, image.Image(filteredImage), nil)
		break
	case "gif":
		gif.Encode(dstFile, image.Image(filteredImage), nil)
		break
	}
}

func process(c *cli.Context) {
	image, format := loadImage(c.Args().Get(0))
	dstPath := c.String("output")

	pattern := c.Int("pattern")
	hInterval := float64(360 / (pattern + 1))
	saturation := Clamp(c.Float64("saturation"), -100.0, 100.0) / 100.0
	lightness := Clamp(c.Float64("lightness"), -100.0, 100.0) / 100.0
	grayscale := c.Bool("grayscale")

	offset := hslOption{
		h: hInterval,
		s: saturation,
		l: lightness,
		g: grayscale,
	}

	dstFilePathBase := getDstPathBase(dstPath, format, pattern)
	for i := 1; i <= pattern; i++ {
		dstFilePath := fmt.Sprintf(dstFilePathBase, i)
		offset.h = hInterval * float64(i)
		generateImage(image, format, dstFilePath, offset)
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "gocova"
	app.Usage = "Go color variation, generate images of various colors"
	app.Version = "1.3.0"

	app.Action = process
	app.Flags = flags
	app.Run(os.Args)
}
