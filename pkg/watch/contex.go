// Package chronos provides performance monitoring routines
package watch

import (
	"context"

	"github.com/google/uuid"
)

// TraceID is context trace id.
type TraceID uuid.UUID

// TraceIDKey is the type of value to use for the key. The key is
// type specific and only values of the same type will match.
type TraceIDKey int

// Declare a key with the value of zero of type userKey.
const traceIDKey TraceIDKey = 0

// WatcherType is context watcher.
type WatcherType *Watch

// WatcherIDKey is the type of value to use for the key. The key is
// type specific and only values of the same type will match.
type WatcherIDKey int

// Declare a key with the value of zero of type userKey.
const watcherIDKey WatcherIDKey = 1

// NewTraceContext creates new context with trace id.
func NewTraceContext(ctx context.Context) context.Context {
	// Create a traceID for this request.
	traceID := TraceID(uuid.New())

	// Store the traceID value inside the context with a value of
	// zero for the key type.
	ctxTrace := context.WithValue(ctx, traceIDKey, traceID)

	return ctxTrace
}

// GetTraceID extracts trace id form the context.
func GetTraceID(ctx context.Context) uuid.UUID {
	// Retrieve that traceID value from the Context value bag.
	if t, ok := ctx.Value(traceIDKey).(TraceID); ok {
		return uuid.UUID(t)
	}

	return uuid.Nil
}

// NewTraceWatchContext creates new context with trace id and watch.
func NewTraceWatchContext(ctx context.Context, watch *Watch) context.Context {
	ctxTrace := NewTraceContext(ctx)

	// Create a traceID for this request.
	watcherID := WatcherType(watch)

	// Store the traceID value inside the context with a value of
	// zero for the key type.
	ctxWatcher := context.WithValue(ctxTrace, watcherIDKey, watcherID)

	return ctxWatcher
}

// GetWatch extracts trace id form the context.
func GetWatch(ctx context.Context) *Watch {
	// Retrieve watcher value from the Context value bag.
	if t, ok := ctx.Value(watcherIDKey).(WatcherType); ok {
		return t
	}

	return nil
}
