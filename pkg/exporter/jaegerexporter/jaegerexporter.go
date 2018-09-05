package jaegerexporter

import (
	"github.com/rs/zerolog/log"

	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

// NewExporterCollector register a new Opencensus to Jaeger exporter
func NewExporterCollector() {
	// Register the Jaeger exporter to be able to retrieve
	// the collected spans.
	exporter, err := jaeger.NewExporter(jaeger.Options{
		Endpoint:    "http://localhost:14268",
		ServiceName: "vpncentralmanager",
	},
	)
	if err != nil {
		log.Fatal().Msg("Error initialize jaeger exporter")
	}
	trace.RegisterExporter(exporter)
}
