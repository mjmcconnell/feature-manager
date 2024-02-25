package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	tracer  = otel.Tracer("rolldice")
	meter   = otel.Meter("rolldice")
	rollCnt metric.Int64Counter
)

func init() {
	var err error
	rollCnt, err = meter.Int64Counter("test_counter",
		metric.WithUnit("{count}"))
	if err != nil {
		panic(err)
	}
}

func NewHttpMetrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracer.Start(context.Background(), "roll")
		defer span.End()

		rollValueAttr := attribute.Int("roll.value", 1)
		span.SetAttributes(rollValueAttr)
		rollCnt.Add(ctx, 1, metric.WithAttributes(rollValueAttr))

		resp := next(c)

		return resp
	}
}
