package utils

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
	"io"
	"net/http"
)

func ParseAndValidateRequest(body io.Reader, target interface{}) error {
	if err := json.NewDecoder(body).Decode(&target); err != nil {
		return err
	}
	if err := validator.Validate(target); err != nil {
		return fmt.Errorf("request body validation error: %s", err.Error())
	}

	return nil
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	if statusCode >= 400 {
		if err := json.NewEncoder(w).Encode(struct {
			Error interface{} `json:"error"'`
		}{
			Error: data,
		}); err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
			return
		}
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
}

func InvalidQueryParamError(message string) string {
	return fmt.Sprintf("Invalid query param: %s", message)
}
func SuccessfulDeleteMessage(table string, id string) string {
	return fmt.Sprintf("Successfully deleted from table `%s` with id: `%s`", table, id)
}
func ConvertToKey(index int, key string) string { return fmt.Sprintf("$%s_%v", key, index) }
