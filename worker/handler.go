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
	"github.com/nfnt/resize"
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
	if _, ok := query["op"]; ok {
		// switch op {
		// case pic.RESIZE:
		// case pic.CROP:
		// }
	} else {
		HTTPResponse(w, 400, ErrUnsupportedOps)
		return
	}

	log.Printf("%v", r.URL.Query())
	start := time.Now()
	img, _, err := image.Decode(imageData)
	resized := resize.Resize(300, 300, img, resize.Bicubic)
	elapsed := time.Since(start)
	if err != nil {
		HTTPResponse(w, 400, ErrDecode)
		return
	}
	log.Printf("size: %v, elapsed: %v", resized.Bounds(), elapsed)

	buf := new(bytes.Buffer)
	if err = jpeg.Encode(buf, resized, nil); err != nil {
	}

	if _, err = w.Write(buf.Bytes()); err != nil {
	}
	w.WriteHeader(http.StatusOK)
}
