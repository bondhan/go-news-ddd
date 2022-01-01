package client

import (
	"encoding/json"
	"fmt"
	"github.com/guregu/null"
	"github.com/pkg/errors"
	"net/http"
)

type ErrorCode int

type Message struct {
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (i *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

type ErrorMessageDict struct {
	Status  int       `json:"-"`
	Type    string    `json:"type"`
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

type ErrorMessage struct {
	ErrorMessageDict
	Trace        *string `json:"-"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

func (e ErrorMessage) Error() string {
	return fmt.Sprintf("%+v: %+v", e.Message, null.StringFromPtr(e.Trace).ValueOrZero())
}

// ComposeErrMsg ..
func ComposeErrMsg(emd ErrorMessageDict, err error) ErrorMessage {

	var errMsg *string
	var errTrace *string

	if err != nil {
		s := err.Error()
		errMsg = &s

		ss := fmt.Sprintf("%+v", err)
		errTrace = &ss
	}

	em := ErrorMessage{
		emd,
		errTrace,
		errMsg,
	}

	return em
}

type errorClient struct {
	errorCodes map[ErrorCode]ErrorMessageDict
}

type ResponseClient interface {
	Error(w http.ResponseWriter, r *http.Request, err error)
	JSON(w http.ResponseWriter, r *http.Request, code int, data interface{})
}

func NewResp(errCodes map[ErrorCode]ErrorMessageDict) ResponseClient {
	return &errorClient{
		errorCodes: errCodes,
	}
}

// Respond is response write to ResponseWriter
func (e *errorClient) respond(w http.ResponseWriter, httpStatusCode int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			e.errorL(w, http.StatusInternalServerError, errors.New("invalid json"))
			return
		}
		body = s
	case string:
		body = []byte(s)
	default:
		if body, err = json.Marshal(src); err != nil {
			e.errorL(w, http.StatusInternalServerError, fmt.Errorf("failed to parse json: %s", err.Error()))
			return
		}
	}
	w.WriteHeader(httpStatusCode)
	w.Write(body)
}

func (e *errorClient) errorL(w http.ResponseWriter, httpStatusCode ErrorCode, err error) {
	var em ErrorMessage
	em.ErrorMessageDict = e.errorCodes[httpStatusCode]
	em.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	e.respond(w, em.ErrorMessageDict.Status, em)
}

// Error always receive err type ErrorMessage, other than that might produce error
func (e *errorClient) Error(w http.ResponseWriter, r *http.Request, err error) {
	em, success := err.(ErrorMessage)
	if !success {
		e.errorL(w, 500, errors.WithStack(errors.New("invalid error message format")))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	e.respond(w, em.ErrorMessageDict.Status, em)
}

// JSON is wrapped Respond when success response
func (e *errorClient) JSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	e.respond(w, code, data)
}
