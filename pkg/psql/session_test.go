package psql

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/testdata"
)

func TestPSQL_SetDB(t *testing.T) {

	in := new(sql.DB)

	db := new(PSQL)

	out := db.SetDB(in)

	assert.NotEmpty(t, out)

}

func TestNew(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svc, err := New(cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, svc)

}
