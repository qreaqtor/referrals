package config

type Config struct {
	Addr Adress `yaml:"addr"`

	Api API `yaml:"api"`
}

type API struct {
	Version int `yaml:"version"`
}
