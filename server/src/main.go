package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	firebase "firebase.google.com/go"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	event "event"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

func main() {
	var (
		// Sets your Google Cloud Platform project ID.
		//projectID = "elliptical-city-243420"

		addr = envString("PORT", defaultPort)
		// rsurl = envString("ROUTINGSERVICE_URL", defaultRoutingServiceURL)

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		// routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")

		ctx = context.Background()
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		logger.Log("error initializing firebase app", err)
	}

	var eventRepo, _ = event.NewEventRepository(app, logger)

	var es event.Service
	es = event.NewService(eventRepo, logger)

	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()
	mux.Handle("/event/", event.MakeHandler(es, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
