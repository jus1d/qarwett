package mock

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Call New function to create a logger
	logger := New()

	// Assert that the logger is not nil
	assert.NotNil(t, logger, "logger should not be nil")

	// Assert that the handler of the logger is a DiscardHandler
	handler, ok := logger.Handler().(*DiscardHandler)
	assert.True(t, ok, "handler should be a DiscardHandler")
	assert.NotNil(t, handler, "handler should not be nil")
}

func TestDiscardHandler_Handle(t *testing.T) {
	// Create a DiscardHandler instance
	handler := NewDiscardHandler()

	// Call the Handle function and assert that it returns nil error
	err := handler.Handle(context.Background(), slog.Record{})
	assert.NoError(t, err, "Handle should return nil error")
}

// Add tests for other methods of DiscardHandler as needed
func TestDiscardHandler_WithAttrs(t *testing.T) {
	// Create a DiscardHandler instance
	handler := NewDiscardHandler()

	// Call the WithAttrs function and assert that it returns the handler itself
	newHandler := handler.WithAttrs(nil)
	assert.Equal(t, handler, newHandler, "WithAttrs should return the handler itself")
}

func TestDiscardHandler_WithGroup(t *testing.T) {
	// Create a DiscardHandler instance
	handler := NewDiscardHandler()

	// Call the WithGroup function and assert that it returns the handler itself
	newHandler := handler.WithGroup("")
	assert.Equal(t, handler, newHandler, "WithGroup should return the handler itself")
}

func TestDiscardHandler_Enabled(t *testing.T) {
	// Create a DiscardHandler instance
	handler := NewDiscardHandler()

	// Call the Enabled function and assert that it always returns false
	enabled := handler.Enabled(context.Background(), slog.LevelInfo)
	assert.False(t, enabled, "Enabled should always return false")
}
