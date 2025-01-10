package bunzap

import (
	"context"
	"time"

	"github.com/dreamph/dbre/adapters/bun/utils"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field names
const (
	OperationFieldName     = "operation"
	OperationTimeFieldName = "operation_time_ms"
)

// QueryHook defines the
// structure of our adapters hook
// it implements the bun.QueryHook
// interface
type QueryHook struct {
	bun.QueryHook

	logger       *zap.Logger
	slowDuration time.Duration
}

// QueryHookOptions defines the
// available options for a new
// adapters hook.
type QueryHookOptions struct {
	Logger       *zap.Logger
	SlowDuration time.Duration
}

// NewQueryHook returns a new adapters hook for use with
// uptrace/bun.
func NewQueryHook(options QueryHookOptions) QueryHook {
	return QueryHook{
		logger:       options.Logger,
		slowDuration: options.SlowDuration,
	}
}

func (qh QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (qh QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	queryDuration := time.Since(event.StartTime)
	fields := []zapcore.Field{
		zap.String(OperationFieldName, event.Operation()),
		zap.Int64(OperationTimeFieldName, queryDuration.Milliseconds()),
	}

	err := utils.DbError(event.Err)
	if err != nil {
		fields = append(fields, zap.Error(event.Err))
		qh.logger.Error(event.Query, fields...)
		return
	}

	// Queries over a slow time duration
	// will be logged as debug
	if queryDuration >= qh.slowDuration {
		qh.logger.Debug(event.Query, fields...)
	}
}
