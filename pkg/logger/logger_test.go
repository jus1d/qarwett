package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// Test for local environment
	log := Init("local")
	assert.NotNil(t, log, "logger should not be nil for local environment")

	// Test for development environment
	log = Init("dev")
	assert.NotNil(t, log, "logger should not be nil for development environment")

	// Test for production environment
	log = Init("prod")
	assert.NotNil(t, log, "logger should not be nil for production environment")

	// Test for unknown environment
	log = Init("unknown")
	assert.Nil(t, log, "logger should be nil for unknown environment")
}

func TestInitPretty(t *testing.T) {
	// Test InitPretty function
	log := InitPretty()
	assert.NotNil(t, log, "logger should not be nil")
}
