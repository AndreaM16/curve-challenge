package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/testdata"
)

func TestCreateMerchant(t *testing.T) {

	name := "someMerchant"
	location := "someLocation"

	cfg := testdata.MockConfiguration

	svc, svcErr := psql.New(cfg)

	assert.NoError(t, svcErr)

	m := new(model.Merchant)
	m.SetName(name).SetLocation(location)

	out, err := CreateMerchant(svc, m)

	assert.NoError(t, err)
	assert.Equal(t, name, out.Name)
	assert.Equal(t, location, out.Location)

}
