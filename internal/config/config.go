package config

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DuplicatePipelineSafeguardPeriod time.Duration `env:"DUPLICATE_PIPELINE_SAFEGUARD_PERIOD" envDefault:"1h"`
	GRPCPort                         int64         `env:"GRPC_PORT" envDefault:"10000"`
	GRPCClientUseTLS                 bool          `env:"GRPC_CLIENT_USE_TLS" envDefault:"true"`
	GRPCClientCertFN                 string        `env:"GRPC_CLIENT_CERT_FN" envDefault:""`
}

func New() (cfg *Config, err error) {
	cfg = new(Config)
	return cfg, env.Parse(cfg)
}
