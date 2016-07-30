package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"reflect"
	"time"

	. "github.com/ctliu3/pic-tailor/util"
	"github.com/ctliu3/pic-tailor/worker/pic"
)

func apiMeta(w http.ResponseWriter, r *http.Request) {
	imageData := r.Body
	if imageData == nil {
		HTTPResponse(w, 400, ErrEmptyBody)
		return
	}

	log.Printf("%v\n", reflect.TypeOf(r.Body))
	start := time.Now()
	meta, err := pic.Meta(imageData)
	elapsed := time.Since(start)
	log.Printf("size: %v, elapsed: %v", meta, elapsed)

	if err != nil {
		HTTPResponse(w, 400, ErrDecode)
		return
	}

	ret, err := json.Marshal(meta)
	if err != nil {
		log.Printf("Encode json data err: %v", err)
		HTTPResponse(w, 500, ErrInernalError)
		return
	}
	HTTPResponse(w, 200, ret)
}

func apiImage(w http.ResponseWriter, r *http.Request) {
	imageData := r.Body
	if imageData == nil {
		HTTPResponse(w, 400, ErrEmptyBody)
		return
	}

	query := r.URL.Query()
	var processed image.Image

	switch op := query.Get("op"); op {
	case pic.RESIZE:
		opt, err := ParseResizeOption(&query)
		if err != nil {
			log.Printf("Parse query parameter err: %v", err)
			HTTPResponse(w, 400, ErrInvalidParametr)
			return
		}
		resized, err := pic.ImgResize(imageData, opt)
		if err != nil {
			log.Printf("Resize image err: %v", err)
			HTTPResponse(w, 500, ErrInernalError)
			return
		}
		processed = resized
	case pic.CROP:
		log.Printf("=== %v\n", op)
	default:
		log.Printf("Invalid operation: %v", op)
		HTTPResponse(w, 400, ErrUnsupportedOps)
		return
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, processed, nil); err != nil {
		log.Printf("Encode image err: %v", err)
		HTTPResponse(w, 500, ErrInernalError)
		return
	}

	HTTPResponse(w, 200, buf.Bytes())
}
