package database

import (
	"gotest.tools/assert"
	"testing"
)

// test db
func TestGetDB(t *testing.T) {
	db1 := GetDB()
	db2 := GetDB()
	assert.Equal(t, db1, db2)

}
