package main

import (
	"github.com/bmatsuo/gotest"
	"testing"
)

func TestDiagnose(t *testing.T) {
	assert.Test = t
	assert.True(diagnose())
}
