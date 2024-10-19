package test

import (
	"testing"
)

func TestFirstRun(t *testing.T) {
	db := getDb(t)
	migrate(t, db)
}
