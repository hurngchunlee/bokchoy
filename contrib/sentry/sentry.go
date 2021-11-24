package sentry

import (
	"context"

	"github.com/getsentry/sentry-go"

	"github.com/hurngchunlee/bokchoy"
)

type SentryTracer struct {
}

func (s SentryTracer) String() string {
	return "sentry"
}

func (s *SentryTracer) Log(ctx context.Context, msg string, err error) {
	sentry.WithScope(func(scope *sentry.Scope) {
		task := bokchoy.GetContextTask(ctx)

		scope.SetContext("msg", msg)
		scope.SetContext("task", task)

		sentry.CaptureException(err)
	})
}
