package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {

	someMsg := "someVal"

	out := NewResponse(someMsg)

	assert.Equal(t, someMsg, out.Message)

}

func TestResponse_SetMessage(t *testing.T) {

	someMsg := "someVal"

	out := new(Response)
	out = out.SetMessage(someMsg)

	assert.Equal(t, someMsg, out.Message)

}

func TestResponse_JsonMarshal(t *testing.T) {

	someMsg := "someVal"

	res := NewResponse(someMsg)

	b, marshalErr := res.JsonMarshal()

	assert.NoError(t, marshalErr)
	assert.NotEmpty(t, b)

	var out Response

	unmarshallErr := json.Unmarshal(b, &out)

	assert.NoError(t, unmarshallErr)
	assert.Equal(t, someMsg, out.Message)

}
