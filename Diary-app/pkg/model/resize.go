package model

import (
	"image"

	"golang.org/x/image/draw"
)

func ResizeImage(src image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

//func main() {
//	file, err := os.Open("../../../../../Desktop/test.png")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer file.Close()
//
//	image, _, err := image.Decode(file)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	resizedImg := ResizeImage(image, 100, 100)
//
//	outFile, err := os.Create("../../../../../Desktop/test_resized.png")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer outFile.Close()
//	if err := png.Encode(outFile, resizedImg); err != nil {
//		log.Fatalln(err)
//	}
//}
//
