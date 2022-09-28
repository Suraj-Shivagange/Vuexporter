package vusmartmapsexporter

import (
	"errors"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtls"
)

// Config defines configuration for Elastic APM exporter.
type Config struct {
	config.ExporterSettings    `mapstructure:",squash"`
	configtls.TLSClientSetting `mapstructure:"tls,omitempty"`

	// APMServerURLs holds the APM Server URL.
	//
	// This is required.
	APMServerURL string `mapstructure:"apm_server_url"`

	// APIKey holds an optional API Key for authorization.
	//
	// https://www.elastic.co/guide/en/apm/server/7.7/api-key-settings.html
	APIKey string `mapstructure:"api_key"`

	// SecretToken holds the optional secret token for authorization.
	//
	// https://www.elastic.co/guide/en/apm/server/7.7/secret-token.html
	SecretToken string `mapstructure:"secret_token"`
}

// Validate validates the configuration.
func (cfg Config) Validate() error {
	if cfg.APMServerURL == "" {
		return errors.New("APMServerURL must be specified")
	}
	return nil
}
