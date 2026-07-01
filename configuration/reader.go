// Provides methods dealing with config files.
package configuration

import (
	"go-base-end/resource"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigReader struct{}

// ReadConfig equals to github.com/ilyakaznacheev/cleanenv.ReadConfig
func (cr ConfigReader) ReadConfig(path string, config any) error {
	return cleanenv.ReadConfig(path, config)
}

func init() {
	resource.Register[ConfigReader](func(opts ...resource.Option) any {
		return ConfigReader{}
	})
}
