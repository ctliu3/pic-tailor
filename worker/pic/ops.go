package pic

import (
	"fmt"
	"image"
	// "image/jpeg"
	_ "image/png"
	"io"
	// "github.com/nfnt/resize"
)

type ImageMeta struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Format string `json:"format"`
}

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

func ImgResize() {
}

func ImgCrop() {
}
