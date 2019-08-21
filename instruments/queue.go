package instruments

import (
	"fmt"
	"happy-q-api/interfaces"
	"happy-q-api/models"
	"time"

	"github.com/go-kit/kit/metrics"
	/* 	"happy-q-api/logs"
	   	"happy-q-api/transports" */)

type InstrumentingQueueMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           interfaces.QueueServiceInterface
}

func (mw InstrumentingQueueMiddleware) Push(s *models.Service, p *models.Person) (output *models.QueueResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "push", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Push(s, p)
	return
}

func (mw InstrumentingQueueMiddleware) Pop(s *models.Service) (q *models.QueueMetadata) {
	defer func(begin time.Time) {
		lvs := []string{"method", "pop", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	q = mw.Next.Pop(s)
	return
}
func (mw InstrumentingQueueMiddleware) View(s *models.Service) (v *models.QueueView, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "view", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	v, err = mw.Next.View(s)
	return
}
