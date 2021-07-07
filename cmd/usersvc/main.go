package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/quyenphamkhac/skoppi/pkg/usersvc/datasources"
	"github.com/quyenphamkhac/skoppi/pkg/usersvc/endpoints"
	"github.com/quyenphamkhac/skoppi/pkg/usersvc/interactor"
	"github.com/quyenphamkhac/skoppi/pkg/usersvc/transports"
)

func main() {
	fs := flag.NewFlagSet("usersvc", flag.ExitOnError)
	var (
		httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var (
		datasource  = datasources.NewMockDatasource()
		svc         = interactor.NewInteractor(datasource, logger)
		endpoints   = endpoints.MakeEndpoints(svc)
		httpHandler = transports.NewHTTPTransport(endpoints)
	)

	httpServer := &http.Server{
		Addr:    *httpAddr,
		Handler: httpHandler,
	}
	httpServer.ListenAndServe()
}
