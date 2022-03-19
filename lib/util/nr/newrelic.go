/**
** TODO: #1 Implementation new project newrelic
** nr is package to handle newrelic
**/
package nr

import (
	"os"
	"time"

	newrelic "github.com/newrelic/go-agent"
)

var client *Client

type Client struct {
	app newrelic.Application
}

// New will initial distributed tracing
func New(appName, licenseKey string, tags map[string]string) (err error) {
	var app newrelic.Application

	// Create new config
	cfg := newrelic.NewConfig(appName, licenseKey)

	// Enable logger preview
	cfg.Logger = newrelic.NewDebugLogger(os.Stdout)

	// Enable Distributed Tracer
	cfg.DistributedTracer.Enabled = true

	// Enable label
	if tags != nil {
		cfg.Labels = tags
	}

	// Initial app newrelic
	if app, err = newrelic.NewApplication(cfg); err != nil {
		return err
	}

	// Wait for the application to connect.
	if err = app.WaitForConnection(5 * time.Second); nil != err {
		return err
	}

	client = &Client{
		app: app,
	}

	return nil
}

// GetApp of newrelic
func GetApp() newrelic.Application {
	return client.app
}
