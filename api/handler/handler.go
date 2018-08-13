package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response embeds a generic response
type Response struct {
	Message string `json:"message"`
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}

// NewResponse returns a *Response initialized with a code and a message
func NewResponse(message string) *Response {

	out := new(Response).SetMessage(message)

	return out

}

// JsonMarshal marshals a Response
func (r Response) JsonMarshal() ([]byte, error) {

	b, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	return b, nil

}

// UnmarshalBody unmarshals a *http.Request body into an output. Output should be a ptr
func UnmarshalBody(r *http.Request, out interface{}) error {

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		return readErr
	}

	unmarshalErr := json.Unmarshal(body, out)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil

}