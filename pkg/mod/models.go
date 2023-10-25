package models

type (
	Config struct {
		Metrics   *MetricsConfig
		Loglevel  int    `mapstructure:"loglevel,omitempty"`
		Bootstrap bool   `mapstructure:"bootstrap,omitempty"`
		Name      string `mapstructure:"name,omitempty"`
		Discovery *DiscoveryConfig
		Server    *ServerConfig
		Join      string   `mapstructure:"join,omitempty"`
		Group     string   `mapstructure:"group,omitempty"`
		Peers     []string `mapstructure:"peers,omitempty"`
		Rack      string   `mapstructure:"rack,omitempty"`
	}

	MetricsConfig struct {
		Bind string `mapstructure:"bind,omitempty"`
		Port int    `mapstructure:"port,omitempty"`
	}

	ServerConfig struct {
		Bind string `mapstructure:"bind,omitempty"`
		Port uint16    `mapstructure:"port,omitempty"`
	}

	DiscoveryConfig struct {
		Bind       string `mapstructure:"bind,omitempty"`
		Port       int    `mapstructure:"port,omitempty"`
		AdvAddress string `mapstructure:"adv-address,omitempty"`
		AdvPort    int    `mapstructure:"adv-port,omitempty"`
	}

	NodeMeta struct {
		Group string
	}
)
