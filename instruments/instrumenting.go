package instruments

import (
	"fmt"
	"time"

	"happy-q-api/services"

	"github.com/go-kit/kit/metrics"
	/* 	"happy-q-api/logs"
	   	"happy-q-api/transports" */)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           services.StringServiceInterface
}

func (mw InstrumentingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Uppercase(s)
	return
}

func (mw InstrumentingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.CountResult.Observe(float64(n))
	}(time.Now())

	n = mw.Next.Count(s)
	return
}
