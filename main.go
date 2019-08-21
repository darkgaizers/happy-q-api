package main

import (
	"net/http"
	"os"

	"happy-q-api/services"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"happy-q-api/instruments"
	"happy-q-api/interfaces"
	repo "happy-q-api/repositories/mongodb"
	"happy-q-api/transports"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "happyq",
		Subsystem: "queue_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "happyq",
		Subsystem: "queue_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	/* 	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	   		Namespace: "happy-q-api",
	   		Subsystem: "queue_service",
	   		Name:      "count_result",
	   		Help:      "The result of each count method.",
	   	}, []string{}) // no fields here

	   	var svc services.StringServiceInterface
	   	svc = services.StringService{}
	   	svc = logs.LoggingMiddleware{logger, svc}
	   	svc = instruments.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	   	uppercaseHandler := httptransport.NewServer(
	   		transports.MakeUppercaseEndpoint(svc),
	   		transports.DecodeUppercaseRequest,
	   		transports.EncodeResponse,
	   	)

	   	countHandler := httptransport.NewServer(
	   		transports.MakeCountEndpoint(svc),
	   		transports.DecodeCountRequest,
	   		transports.EncodeResponse,
	   	)
	   	/////////////////////////////////
	   	http.Handle("/uppercase", uppercaseHandler)
	   	http.Handle("/count", countHandler)
	   	http.Handle("/metrics", promhttp.Handler()) */

	var qs interfaces.QueueServiceInterface
	var ps interfaces.PersonServiceInterface
	var pr interfaces.PersonRepository
	pr = repo.NewMongoDBPersonRepository("")
	qs = services.QueueService{}
	ps = services.NewPersonService(pr)
	/* qs = logs.LoggingQueueMiddleware(logger, qs) */
	qs = instruments.InstrumentingQueueMiddleware{requestCount, requestLatency, qs}
	popHandler := httptransport.NewServer(
		transports.MakeQueuePopEndpoint(qs, ps),
		transports.DecodeQueuePopRequest,
		transports.EncodeResponse,
	)

	pushHandler := httptransport.NewServer(
		transports.MakeQueuePushEndpoint(qs),
		transports.DecodeQueuePushRequest,
		transports.EncodeResponse,
	)
	viewHandler := httptransport.NewServer(
		transports.MakeQueueViewEndpoint(qs, ps),
		transports.DecodeQueueViewRequest,
		transports.EncodeResponse,
	)
	http.Handle("/queue/push", pushHandler)
	http.Handle("/queue/pop", popHandler)
	http.Handle("/queue/view", viewHandler)
	http.Handle("/queue/metrics", promhttp.Handler())

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe("localhost:8080", nil))
}
