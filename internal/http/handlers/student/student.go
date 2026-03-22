package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Vaibhavkum96/student-api-go/internal/types"
	"github.com/Vaibhavkum96/student-api-go/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		//Empty Body Error
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Error is not nil

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		slog.Info("Creating A Student!")
		w.Write([]byte("Welcome To Students Api!"))

		//Request Validation

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "Ok"})
	}
}
