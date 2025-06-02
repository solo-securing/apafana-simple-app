package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/solo-securing/apafana-simple-app/pkg/plugin"
)

func main() {
	app := plugin.NewApp()

	err := backend.Serve(backend.ServeOpts{
		CallResourceHandler: app,
	})
	if err != nil {
		log.DefaultLogger.Error("Failed to start backend server", "error", err)
		os.Exit(1)
	}
}
