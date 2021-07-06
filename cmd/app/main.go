package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/quyenphamkhac/skoppi/services/usersvc/datasources"
	"github.com/quyenphamkhac/skoppi/services/usersvc/endpoints"
	"github.com/quyenphamkhac/skoppi/services/usersvc/interactor"
	"github.com/quyenphamkhac/skoppi/services/usersvc/transports"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "order",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	ds := datasources.NewMockDatasource()
	svc := interactor.NewInteractor(ds, logger)
	enpoints := endpoints.MakeEndpoints(svc)
	handler := transports.NewHTTPTransport(enpoints)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	httpServer.ListenAndServe()
}
