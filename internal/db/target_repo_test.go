package db

import (
	"testing"
)

func TestTargetRepo(t *testing.T) {
	Connect()
	GetTargetDB()
}
