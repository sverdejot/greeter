package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/sverdejot/greeter/users/internal/application"
	"github.com/sverdejot/greeter/users/internal/application/uc"
)

func NewCreateUserHandler(creator *uc.UseCaseCreateUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(fmt.Errorf("cannot parse id: %w", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		req := &struct{
			Name string `json:"name"`
			Mail string `json:"mail"`
			Age int `json:"age"`
			Status int `json:"status"`
		}{}
		err = json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			log.Println(fmt.Errorf("cannot parse req body: %w", err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err = creator.Run(r.Context(), id, req.Name, req.Mail, req.Age, req.Status)

		if errors.Is(err, application.ValidationError{}) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
