package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/solo-securing/apafana-simple-app/pkg/plugin"
)

func main() {
	app := plugin.NewApp()

	opts := backend.ServeOpts{
		CheckHealthHandler:  app,
		CallResourceHandler: app,
	}

	if err := backend.Serve(opts); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
