package vusmartmapsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.uber.org/zap"
)

const (
	// The value of "type" key in configuration.
	typeStr = "vusmartmaps"
	// The stability level of the exporter.
	stability = component.StabilityLevelDeprecated
)

var once sync.Once

// NewFactory creates a factory for Elastic exporter.
func NewFactory() component.ExporterFactory {
	return component.NewExporterFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesExporter(createTracesExporter, stability),
		component.WithMetricsExporter(createMetricsExporter, stability),
	)
}

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings: config.NewExporterSettings(config.NewComponentID(typeStr)),
	}
}

func logDeprecation(logger *zap.Logger) {
	once.Do(func() {
		logger.Warn("elastic exporter is deprecated and will be removed in future versions.")
	})
}

func createTracesExporter(
	ctx context.Context,
	params component.ExporterCreateSettings,
	cfg config.Exporter,
) (component.TracesExporter, error) {
	logDeprecation(params.Logger)
	return newElasticTracesExporter(params, cfg)
}

func createMetricsExporter(
	ctx context.Context,
	params component.ExporterCreateSettings,
	cfg config.Exporter,
) (component.MetricsExporter, error) {
	logDeprecation(params.Logger)
	return newElasticMetricsExporter(params, cfg)
}
