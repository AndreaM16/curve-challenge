package psql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andream16/curve-challenge/testdata"
)

func TestPSQL_CreateTable(t *testing.T) {

	q := `CREATE TABLE IF NOT EXISTS someTable ( some_field text PRIMARY KEY )`

	cfg := testdata.MockConfiguration

	svc := mockNewSession(t, cfg)

	err := svc.CreateTable(q)

	assert.NoError(t, err)

}
