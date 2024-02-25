package response

import (
	"bytes"
	"net/http"

	"github.com/ak1m1tsu/tech-tinker/internal/lib/decoder"
)

func JSON(w http.ResponseWriter, code int, data M) {
	buf := new(bytes.Buffer)
	err := decoder.EncodeJSON(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(buf.Bytes())
}
