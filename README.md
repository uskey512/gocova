# gocova  
gocova is various color image generator.  
It generates images of multiple patterns from one image.  

<img width="874" alt="" src="https://user-images.githubusercontent.com/4005383/47588820-5b084c80-d9a2-11e8-891e-49aed3ff3323.png">



## Install  
`$ go get -u github.com/uskey512/gocova`  

## Usage  
`$ gocova -i {input file path} -o {output file path} -p {number of images to generate}`  

```
Options:
   --input value, -i value       input image path
   --output value, -o value      output image path base (default: "./result")
   --pattern value, -p value     number of images to generate (default: 10)
   --saturation value, -s value  saturation offset [-100.0...100.0] (default: 0)
   --lightness value, -l value   lightness offset [-100.0...100.0] (default: 0)
```

(Input and output should be in png format)


## Author
Yusuke Uehara (@uskey512)

## License
MIT
