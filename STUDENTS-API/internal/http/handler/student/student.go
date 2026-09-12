package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/sunny/students-api/internal/types"
	"github.com/sunny/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student;
		err := json.NewDecoder(r.Body).Decode(&student);
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Empty body")));
			return;
		}
		slog.Info("create a student")
		response.WriteJson(w, http.StatusCreated, map[string] string {"success": "OK"})
		//w.Write([]byte("Welcome to the Students API!"))
	}
}
