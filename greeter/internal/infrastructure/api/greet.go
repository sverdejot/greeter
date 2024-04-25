package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sverdejot/greeter/greeter/internal/application"
)

func NewGreeterHandler(greeter *application.Greeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		msg, err := greeter.Greet(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := struct {
			Message string `json:"message"`
		}{
			Message: msg,
		}

		json.NewEncoder(w).Encode(body)
	}
}
