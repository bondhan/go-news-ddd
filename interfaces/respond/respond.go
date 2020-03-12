package respond

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// ErrorResponse is Error response template
type ErrorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
	Code   int    `json:"code"`
}

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("reason: %s", e.Reason)

}

// Respond is response write to ResponseWriter
func Respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, errors.New("Invalid JSON"))
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *ErrorResponse, ErrorResponse:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, fmt.Errorf("Failed to parse Json: %s", err.Error()))
			return
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int, err error) {
	e := &ErrorResponse{
		Status: "Unsuccessful",
		Reason: err.Error(),
		Code:   code,
	}
	logrus.Debugf("%v\n", e)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, e)
}

// JSON is wrapped Respond when success response
func JSON(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, src)
}

// JSON is wrapped Respond when success response
func Success(w http.ResponseWriter, code int, msg string) {
	status := Response{
		Status: "Successful",
		Code:   code,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := json.Marshal(status)
	if err != nil {
		Error(w, http.StatusInternalServerError, fmt.Errorf("Failed to parse Json: %s", err.Error()))
		return
	}
	w.WriteHeader(code)
	w.Write(body)
}
