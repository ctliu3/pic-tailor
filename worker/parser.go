package main

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/ctliu3/pic-tailor/worker/pic"
)

func ParseResizeOption(values *url.Values) (*pic.ResizeOption, error) {
	opt := &pic.ResizeOption{}

	strw, strh := values.Get("w"), values.Get("h")
	if strw == "" && strh == "" {
		return nil, fmt.Errorf("w or h parameter not found")
	}
	if h, err := strconv.ParseUint(strh, 10, 32); err != nil {
		return nil, fmt.Errorf("Parse height not found, h: %v", strh)
	} else {
		opt.Height = uint(h)
	}
	if w, err := strconv.ParseUint(strh, 10, 32); err != nil {
		return nil, fmt.Errorf("Parse height not found, w: %v", strw)
	} else {
		opt.Width = uint(w)
	}
	if opt.Width == 0 && opt.Height == 0 {
		return nil, fmt.Errorf("Invalid resize height and width")
	}

	interp := values.Get("interp")
	if interp == "" {
		interp = pic.BILINEAR
	} else {
		found := false
		for _, method := range pic.INTERPOLATIONS {
			if method == interp {
				found = true
			}
		}
		if !found {
			return nil, fmt.Errorf("unkown interpolation method: %v", interp)
		}
	}
	opt.Interpolation = interp

	// TODO
	// Parse qualtiy

	return opt, nil
}
