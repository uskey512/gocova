# gocova  
gocova is various color image generator.  
It generates images of multiple patterns from one image.  

<img width="874" alt="" src="https://user-images.githubusercontent.com/4005383/47588820-5b084c80-d9a2-11e8-891e-49aed3ff3323.png">

## Installation  
`$ go get -u github.com/uskey512/gocova`  

## Usage  
```
gocova　[options] <input>　　

Options:
   --output value, -o value      output image path base (default: "./result")
   --pattern value, -p value     number of images to generate (default: 10)
   --saturation value, -s value  saturation offset [-100.0 ... 100.0] (default: 0)
   --lightness value, -l value   lightness offset [-100.0 ... 100.0] (default: 0)
```
Input file should be in png format.

In gocova, converting from RGB to HSL internally.  
( see https://en.wikipedia.org/wiki/HSL_and_HSV )

#### -l, --lightness
```
$ gocova -l 30 source.png
```

#### -s, --saturation
```
$ gocova -s -40 source.png
```

## Author
Yusuke Uehara (@uskey512)

## License
MIT
