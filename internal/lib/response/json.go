package response

import (
	"bytes"
	"github.com/insan1a/tech-tinker/internal/lib/decoder"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, data interface{}) {
	var buf *bytes.Buffer
	err := decoder.EncodeJSON(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(buf.Bytes())
}
