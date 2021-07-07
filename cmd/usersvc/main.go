package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

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
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])

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

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
