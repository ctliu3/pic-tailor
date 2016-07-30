package util

import (
	"encoding/json"
)

const (
	Unavailable uint8 = iota
	Invalid
	Unknown
	BadRequest
	Unsupported
	InternalError
	NotFound
)

type Error struct {
	Message string
	Code    uint8
}

var (
	ErrNotFound            = NewError("Not found", NotFound)
	ErrUnsupportedParamter = NewError("Unsupported parameter", Unsupported)
	ErrUnsupportedFormat   = NewError("Unsupported media type", Unsupported)
	ErrDecode              = NewError("Decode image error", Invalid)
	ErrOutputFormat        = NewError("Unsupported output image format", Unsupported)
	ErrEmptyBody           = NewError("Empty body", BadRequest)
	ErrUnsupportedOps      = NewError("Unsupported image operation", Unsupported)
	ErrUnknown             = NewError("Unknown error", Unknown)
	ErrInernalError        = NewError("Internal error", InternalError)
	ErrInvalidParametr     = NewError("Invalid parameter", Invalid)
)

func (e Error) JSON() []byte {
	buf, _ := json.Marshal(e)
	return buf
}

func NewError(msg string, code uint8) Error {
	return Error{msg, code}
}
