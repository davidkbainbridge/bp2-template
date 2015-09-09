package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	gkmetrics "github.com/go-kit/kit/metrics"
	gkprometheus "github.com/go-kit/kit/metrics/prometheus"
	gkhttptransport "github.com/go-kit/kit/transport/http"

	"golang.org/x/net/context"
)

func Register() {
	ctx := context.Background()

	fieldKeys := []string{"method", "error"}
	requestCount := gkprometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "blueplanet",
		Subsystem: "bp2_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := gkmetrics.NewTimeHistogram(time.Microsecond, gkprometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "blueplanet",
		Subsystem: "bp2_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duraction of the requests in microseconds.",
	}, fieldKeys))

	var svc StringService
	svc = stringService{}
	svc = instrumentingMiddleware{requestCount, requestLatency, svc}

	uppercaseHandler := gkhttptransport.Server{
		Context:            ctx,
		Endpoint:           makeUppercaseEndpoint(svc),
		DecodeRequestFunc:  decodeUppercaseRequest,
		EncodeResponseFunc: encodeResponse,
	}

	countHandler := gkhttptransport.Server{
		Context:            ctx,
		Endpoint:           makeCountEndpoint(svc),
		DecodeRequestFunc:  decodeCountRequest,
		EncodeResponseFunc: encodeResponse,
	}

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", prometheus.Handler())
}

func decodeUppercaseRequest(r *http.Request) (interface{}, error) {
	var req uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCountRequest(r *http.Request) (interface{}, error) {
	var req countRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}
