package prettyslog

import (
	"bytes"
	"context"
	stdLog "log"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrettyHandler_Handle(t *testing.T) {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	h := opts.NewPrettyHandler(os.Stdout)

	record := slog.Record{
		Time:    time.Now(),
		Level:   slog.LevelInfo,
		Message: "Test message",
	}
	record.AddAttrs(slog.String("key1", "value1"), slog.Int64("key2", 42))

	err := h.Handle(context.Background(), record)

	assert.NoError(t, err, "Handle should return no error")

	record = slog.Record{
		Time:    time.Now(),
		Level:   slog.LevelWarn,
		Message: "Test message",
	}

	err = h.Handle(context.Background(), record)

	assert.NoError(t, err, "Handle should return no error")

	record = slog.Record{
		Time:    time.Now(),
		Level:   slog.LevelDebug,
		Message: "Test message",
	}

	err = h.Handle(context.Background(), record)

	assert.NoError(t, err, "Handle should return no error")

	record = slog.Record{
		Time:    time.Now(),
		Level:   slog.LevelError,
		Message: "Test message",
	}

	err = h.Handle(context.Background(), record)

	assert.NoError(t, err, "Handle should return no error")
}

func TestPrettyHandler_WithAttrs(t *testing.T) {
	// Create a PrettyHandler instance
	handler := &PrettyHandler{
		// Initialize a stub Logger
		logger: stdLog.New(bytes.NewBuffer(nil), "", 0),
		// Create a stub Handler
		Handler: slog.NewJSONHandler(nil, nil),
	}

	// Create a sample slice of attributes
	attrs := []slog.Attr{
		slog.String("key1", "value1"),
		slog.Int64("key2", 42),
	}

	// Call the WithAttrs method
	newHandler := handler.WithAttrs(attrs)

	// Check if the returned handler is a PrettyHandler
	prettyHandler, ok := newHandler.(*PrettyHandler)
	assert.True(t, ok, "Returned handler should be a PrettyHandler")

	// Check if the attrs field of the new handler matches the provided attributes
	assert.Equal(t, attrs, prettyHandler.attrs, "Attrs should match")
}

func TestPrettyHandler_WithGroup(t *testing.T) {
	// Create a PrettyHandler instance
	handler := &PrettyHandler{
		// Initialize a stub Logger
		logger: stdLog.New(bytes.NewBuffer(nil), "", 0),
		// Create a stub Handler
		Handler: slog.NewJSONHandler(nil, nil),
	}

	// Call the WithGroup method
	newHandler := handler.WithGroup("test-group")

	// Check if the returned handler is a PrettyHandler
	_, ok := newHandler.(*PrettyHandler)
	assert.True(t, ok, "Returned handler should be a PrettyHandler")
}
