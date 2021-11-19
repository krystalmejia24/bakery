package main

import (
	"log"
	"net/http"

	"github.com/cbsinteractive/bakery/config"
	"github.com/cbsinteractive/bakery/handlers"
	"github.com/cbsinteractive/pkg/tracing"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err = c.ValidateAuthHeader(); err != nil {
		log.Fatal(err)
	}

	handler := c.SetupMiddleware().Then(handlers.LoadHandler(c))
	hcHandler := c.SetupMiddleware().Then(&handlers.HealthcheckHandler{})

	c.Logger.Info().Str("port", c.Listen).Str("hostname", c.Hostname).Str("git_sha", handlers.GitSHA).Msg("Starting Bakery")
	http.Handle(handlers.HealthcheckPath, c.Client.Tracer.Handle(tracing.FixedNamer("bakery"), hcHandler))
	http.Handle("/", c.Client.Tracer.Handle(tracing.FixedNamer("bakery"), handler))
	if err := http.ListenAndServe(c.Listen, nil); err != nil {
		log.Fatal(err)
	}
}
