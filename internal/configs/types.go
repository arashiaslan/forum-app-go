package configs

type (
	Configs struct {
		Service Service `mapstructure:"service"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}
)
