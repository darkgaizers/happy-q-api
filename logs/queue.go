package logs

import (
	"happy-q-api/interfaces"
	"happy-q-api/models"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingQueueMiddleware struct {
	Logger log.Logger
	Next   interfaces.QueueServiceInterface
}

func (mw LoggingQueueMiddleware) Push(s *models.Service, p *models.Person) (output *models.QueueResult, err error) {
	input := "SID=" + s.ID + ";UID=" + p.ID
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", input,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Push(s, p)
	return
}

func (mw LoggingQueueMiddleware) Pop(s *models.Service) (u *models.Person) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "count",
			"input", "SID="+s.ID,
			"output", "UID="+u.ID,
			"took", time.Since(begin),
		)
	}(time.Now())

	u = mw.Next.Pop(s)
	return
}
func (mw LoggingQueueMiddleware) View(s *models.Service) (v *models.QueueView, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "count",
			"input", "SID="+s.ID,
			"output", "total_queue="+strconv.Itoa(v.CurrentQueue),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	v, err = mw.Next.View(s)
	return
}
