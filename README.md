# gocova  
gocova is various color image generator.  
It generates images of multiple patterns from one image.  

<img width="500" alt="title" src="https://user-images.githubusercontent.com/4005383/47617620-65def080-db0c-11e8-905c-aa2ea607272c.png">

## Installation  
`$ go get -u github.com/uskey512/gocova`  

## Usage  
```
gocova　[options] <input>　　

options:
   --output value, -o value      output image path base (default: "./result")
   --pattern value, -p value     number of images to generate (default: 10)
   --saturation value, -s value  saturation offset [-100.0 ... 100.0] (default: 0)
   --lightness value, -l value   lightness offset [-100.0 ... 100.0] (default: 0)
```
Supported formats : png, jpeg, gif  

In gocova, converting from RGB to HSL internally.  
( see https://en.wikipedia.org/wiki/HSL_and_HSV )

### -l, --lightness
```
$ gocova -l 30 source.png
```
<img width="500" alt="l30" src="https://user-images.githubusercontent.com/4005383/47617653-ae96a980-db0c-11e8-99eb-b1c24180904b.png">

### -s, --saturation
```
$ gocova -s -40 source.png
```
<img width="500" alt="s-40" src="https://user-images.githubusercontent.com/4005383/47617662-c3733d00-db0c-11e8-984f-a28bead1efce.png">

## Thanks
http://pictogram2.com/?lang=en

## Author
Yusuke Uehara (@uskey512)

## License
MIT
