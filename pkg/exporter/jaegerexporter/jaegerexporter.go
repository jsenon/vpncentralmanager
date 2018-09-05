package jaegerexporter

import (
	"os"
	"runtime"

	"github.com/rs/zerolog/log"

	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

// NewExporterCollector register a new Opencensus to Jaeger exporter
func NewExporterCollector() {
	// Register the Jaeger exporter to be able to retrieve
	// the collected spans.
	addressjaeger := os.Getenv("JAEGER_URL")
	exporter, err := jaeger.NewExporter(jaeger.Options{
		Endpoint:    addressjaeger,
		ServiceName: "vpncentralmanager",
	},
	)
	if err != nil {
		log.Error().Msgf("Error %s", err.Error())
		runtime.Goexit()
	}
	trace.RegisterExporter(exporter)
}
