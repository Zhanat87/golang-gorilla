package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadEnvFile(t *testing.T) {
	assert.Equal(t, "8080", os.Getenv("PORT"))
}
