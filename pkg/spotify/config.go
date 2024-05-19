package spotify

type Config struct {
	Credentials `yaml:",inline"`

	RedirectURI string `yaml:"redirect_uri"`
}
