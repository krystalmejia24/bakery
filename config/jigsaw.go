package config

type Jigsaw struct {
	Alias map[string]string `envconfig:"JIGSAW_HOST_ALIAS"`
}
