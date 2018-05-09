package helper

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

//BuildMini comprimir imagenes
func BuildMini(source, destino string) error {
	img, err := imaging.Open(source)
	if err != nil {
		return err
	}
	rst := imaging.Fill(img, 250, 250, imaging.Center, imaging.Lanczos)
	fle, err := os.Create(destino)
	if err != nil {
		return err
	}
	jpeg.Encode(fle, rst, nil)
	return nil
}

//BuildJPG convertir en jpg
func BuildJPG(source, destino string) error {
	img, err := imaging.Open(source)
	if err != nil {
		return err
	}
	X := img.Bounds().Size().X
	Y := img.Bounds().Size().Y
	var rst *image.NRGBA
	if X > Y {
		rst = imaging.Resize(img, 700, 0, imaging.Lanczos)
	} else {
		rst = imaging.Resize(img, 0, 700, imaging.Lanczos)
	}
	fle, err := os.Create(destino)
	if err != nil {
		return err
	}
	var e interface{} = "hola"
	fmt.Println(fmt.Sprintf("%v", e))

	defer fle.Close()
	err = jpeg.Encode(fle, rst, nil)
	return err
}
