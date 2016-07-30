package pic

import (
	"fmt"
	"image"
	// "image/jpeg"
	_ "image/png"
	"io"

	"github.com/nfnt/resize"
)

type ImageMeta struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Format string `json:"format"`
}

type ResizeOption struct {
	Height        uint
	Width         uint
	Quality       uint
	Interpolation string
}

const (
	BILINEAR = "bilinear"
	BICUBIC  = "bicubic"
)

var INTERPOLATIONS []string = []string{BILINEAR, BICUBIC}

func decode(data io.Reader) (image.Image, string, error) {
	img, format, err := image.Decode(data)
	if err != nil {
		return nil, "", fmt.Errorf("err: %v", err)
	}
	return img, format, nil
}

func Meta(data io.Reader) (*ImageMeta, error) {
	img, format, err := decode(data)
	if err != nil {
		return nil, err
	}
	meta := &ImageMeta{img.Bounds().Max.X, img.Bounds().Max.Y, format}
	return meta, nil
}

func ImgResize(data io.Reader, opt *ResizeOption) (image.Image, error) {
	image, _, err := decode(data)
	if err != nil {
		return nil, err
	}

	imgh := image.Bounds().Max.X
	imgw := image.Bounds().Max.Y
	if opt.Width == 0 {
		opt.Width = uint(int(opt.Height) * imgw / imgh)
	}
	if opt.Height == 0 {
		opt.Height = uint(int(opt.Width) * imgh / imgw)
	}

	var interp resize.InterpolationFunction
	switch opt.Interpolation {
	case BILINEAR:
		interp = resize.Bilinear
	case BICUBIC:
		interp = resize.Bicubic
	default:
		return nil, fmt.Errorf("Unkonw interpolation method: %v", opt.Interpolation)
	}

	resized := resize.Resize(opt.Width, opt.Height, image, interp)
	if resized == nil {
		return nil, fmt.Errorf("Resize err, interpolation: %v", opt.Interpolation)
	}

	return resized, nil
}

func ImgCrop() {
}
