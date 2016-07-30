package util

import (
	"fmt"
	"net/http"
)

func HTTPResponse(w http.ResponseWriter, httpCode int, msg interface{}) error {
	w.WriteHeader(httpCode)

	switch m := msg.(type) {
	case string:
		w.Write([]byte(m))
	case []byte:
		w.Write([]byte(m))
	case Error:
		w.Header().Set("Content-Type", "application/json")
		w.Write(m.JSON())
	default:
		return fmt.Errorf("illegal message format %v", m)
	}

	return nil
}
